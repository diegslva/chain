package chain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// GetTransactionMultiWorkers determines how many worker go routines are used
// concurrently to get transactions from the Chain.com API endpoint.
const GetTransactionMultiWorkers = 5

// Input represents a Bitcoin transaction input.
type Input struct {
	TransactionHash string `json:"transaction_hash"`
	OutputHash      string `json:"output_hash"`
	OutputIndex     uint32 `json:"output_index"`
	Value           int64
	Addresses       []string
	ScriptSignature string `json:"script_signature"`
	Sequence        uint32

	// Only populated with coinbase transactions.
	Coinbase string
}

// Output represents a Bitcoin transaction output.
type Output struct {
	TransactionHash    string `json:"transaction_hash"`
	OutputIndex        uint32 `json:"output_index"`
	Value              int64
	Addresses          []string
	Script             string
	ScriptHex          string `json:"script_hex"`
	ScriptType         string `json:"script_type"`
	RequiredSignatures int64  `json:"required_signatures"`
	Spent              bool

	// Only populated with GetAddressUnspentOutputs.
	Confirmations int64
}

// Transaction representes a Bitcoin transaction.
//
// Chain documentation can be found here
// https://chain.com/docs#object-bitcoin-transaction.
type Transaction struct {
	Hash          string
	BlockHash     string `json:"block_hash"`
	BlockHeight   int64  `json:"block_height"`
	BlockTime     string `json:"block_time"`
	Inputs        []Input
	Outputs       []Output
	Amount        int64
	Fees          int64
	Confirmations int64
}

func (c *Chain) transactionURL(hash string) string {
	return fmt.Sprintf("%s/%s/transactions/%s",
		baseURL, c.network, hash)
}

// GetTransaction returns details about a Bitcoin transaction, including
// inputs and outputs.
//
// Chain documentation can be found here
// https://chain.com/docs#bitcoin-transaction.
func (c *Chain) GetTransaction(hash string) (Transaction, error) {
	url, tx := c.transactionURL(hash), Transaction{}
	return tx, c.httpGetJSON(url, &tx)
}

func (c *Chain) sendTransactionURL() string {
	return fmt.Sprintf("%s/%s/transactions",
		baseURL, c.network)
}

// GetTransactionMulti returns a Transaction slice for all the TransactionHashes
// within the block. Note that it currently calls the chain.com API endpoint
// multiple times. This function produces an error if any of the API endpoint
// calls fails.
func (c *Chain) GetTransactionMulti(hashes []string) ([]Transaction, error) {
	type request struct {
		index int
		hash  string
	}
	type response struct {
		index int
		tx    Transaction
		err   error
	}

	txns := make([]Transaction, len(hashes))
	errs := make(MultiError, len(hashes))
	requestChan := make(chan request, len(hashes))
	responseChan := make(chan response)

	for i := 0; i < GetTransactionMultiWorkers; i++ {
		go func() {
			for req := range requestChan {
				tx, err := c.GetTransaction(req.hash)
				responseChan <- response{req.index, tx, err}
			}
		}()
	}

	for i, hash := range hashes {
		requestChan <- request{i, hash}
	}
	close(requestChan)

	isErrors := false
	for i := 0; i < len(hashes); i++ {
		resp := <-responseChan

		txns[resp.index] = resp.tx
		if resp.err != nil {
			isErrors = true
		}
		errs[resp.index] = resp.err
	}
	close(responseChan)

	if isErrors {
		return txns, errs
	}
	return txns, nil
}

// SendTransaction accepts a signed transaction in hex format and sends it to
// the Bitcoin network. See http://blog.chain.com/post/86529167421/sending-bitcoin-transactions-with-node-js
// for information on creating and signing raw transactions. The transaction
// hash is returned on a successful send.
//
// Chain documentation can be found here
// https://chain.com/docs#bitcoin-transaction-send.
func (c *Chain) SendTransaction(hex string) (string, error) {
	url := c.sendTransactionURL()

	jsonRequest := struct {
		Hex string `json:"hex"`
	}{hex}

	requestBody, err := json.Marshal(jsonRequest)
	if err != nil {
		return "", err
	}
	response, err := c.httpPutJSON(url, bytes.NewReader(requestBody))
	if err != nil {
		return "", err
	}

	responseBody, err := ioutil.ReadAll(response)
	if err != nil {
		return "", err
	}
	jsonResponse := struct {
		TransactionHash string `json:"transaction_hash"`
	}{}

	if err := json.Unmarshal(responseBody, &jsonResponse); err != nil {
		return "", err
	}

	responseString := string(responseBody)
	if jsonResponse.TransactionHash != "" {
		return jsonResponse.TransactionHash, nil
	}
	return "", fmt.Errorf("unknown response %s", responseString)
}
