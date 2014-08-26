package chain

import (
	"errors"
	"fmt"
	"strings"
)

const (
	MaxAddresses = 200

	MaxAddressTransactionsLimit     = 500
	DefaultAddressTransactionsLimit = 50
)

type Address struct {
	Hash                string
	Balance             int64
	Received            int64
	Sent                int64
	UnconfirmedReceived int64 `json:"unconfirmed_received"`
	UnconfirmedSent     int64 `json:"unconfirmed_sent"`
	UnconfirmedBalance  int64 `json:"unconfirmed_balance"`
}

func (c *Chain) addressURL(hashes []string) string {
	return fmt.Sprintf("%s/addresses/%s", c.network, strings.Join(hashes, ","))
}

func (c *Chain) GetAddressMulti(hashes []string) ([]Address, error) {
	if len(hashes) > MaxAddresses {
		return nil, fmt.Errorf("max addresses allowed is %d", MaxAddresses)
	}

	url, addresses := c.addressURL(hashes), make([]Address, len(hashes))
	return addresses, c.httpGetJSON(url, &addresses)
}

func (c *Chain) GetAddress(hash string) (Address, error) {
	url, address := c.addressURL([]string{hash}), &Address{}
	return *address, c.httpGetJSON(url, address)
}

func (c *Chain) addressTransactionsURL(hashes []string, limit int) string {
	return fmt.Sprintf("%s/addresses/%s/transactions?limit=%d",
		c.network, strings.Join(hashes, ","), limit)
}

func (c *Chain) GetAddressTransactionsMulti(
	hashes []string, limit int) ([]Transaction, error) {
	if len(hashes) > MaxAddresses {
		return nil, fmt.Errorf("max addresses allowed is %d", MaxAddresses)
	}
	switch {
	case limit < 0:
		return nil, errors.New("limit must be >= 0")
	case limit > MaxAddressTransactionsLimit:
		return nil, fmt.Errorf("limit must be < %d",
			MaxAddressTransactionsLimit)
	case limit == 0:
		limit = DefaultAddressTransactionsLimit
	}

	url := c.addressTransactionsURL(hashes, limit)
	transactions := make([]Transaction, limit)
	return transactions, c.httpGetJSON(url, &transactions)
}

func (c *Chain) GetAddressTransactions(
	hash string, limit int) ([]Transaction, error) {
	return c.GetAddressTransactionsMulti([]string{hash}, limit)
}

func (c *Chain) addressUnspentOutputsURL(hashes []string) string {
	return fmt.Sprintf("%s/addresses/%s/unspents",
		c.network, strings.Join(hashes, ","))
}

func (c *Chain) GetAddressUnspentOutputsMulti(
	hashes []string) ([]Output, error) {

	if len(hashes) > MaxAddresses {
		return nil, fmt.Errorf("max addresses allowed is %d", MaxAddresses)
	}

	url := c.addressUnspentOutputsURL(hashes)
	outputs := make([]Output, len(hashes))
	return outputs, c.httpGetJSON(url, &outputs)
}

func (c *Chain) GetAddressUnspentOutputs(hash string) ([]Output, error) {
	return c.GetAddressUnspentOutputsMulti([]string{hash})
}
