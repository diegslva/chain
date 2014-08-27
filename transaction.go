package chain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Input represents a Bitcoin transaction input.
type Input struct {
	TransactionHash string `json:"transaction_hash"`
	OutputHash      string `json:"output_hash"`
	OutputIndex     int64  `json:"output_index"`
	Value           int64
	Addresses       []string
	ScriptSignature string `json:"script_signature"`
}

// Output represents a Bitcoin transaction output.
type Output struct {
	TransactionHash    string `json:"transaction_hash"`
	OutputIndex        uint32 `json:"output_index"`
	Value              int64
	Addresses          []string
	Script             string
	ScriptHex          string `json:"script_hex"`
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
	return fmt.Sprintf("%s/transactions/%s", c.network, hash)
}

// GetTransaction returns details about a Bitcoin transaction, including
// inputs and outputs.
//
// Chain documentation can be found here
// https://chain.com/docs#bitcoin-transaction.
func (c *Chain) GetTransaction(hash string) (Transaction, error) {
	url, tx := c.transactionURL(hash), &Transaction{}
	return *tx, c.httpGetJSON(url, tx)
}

func (c *Chain) sendTransactionURL() string {
	return fmt.Sprintf("%s/transactions", c.network)
}

// SendTransaction accepts a signed transaction in hex format and sends it to
// the Bitcoin network. See http://blog.chain.com/post/86529167421/sending-bitcoin-transactions-with-node-js
// for information on creating and signing raw transactions.
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
