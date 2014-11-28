package chain

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// MaxAddresses is the maximum number of addresses that a Multi method
	// can take.
	MaxAddresses = 200

	// MaxAddressTransactionsLimit is the maximum number of transactions a
	// GetAddressTransactions call can return.
	MaxAddressTransactionsLimit = 500

	// DefaultAddressTransactionsLimit is the default number of transactions a
	// GetAddressTransactions call can return.
	DefaultAddressTransactionsLimit = 50
)

// Address represents a bitcoin address.
//
// Chain documentation can be found here
// https://chain.com/docs#object-bitcoin-address.
type Address struct {
	Address string
	Total   struct {
		Balance  int64
		Received int64
		Sent     int64
	}
	Confirmed struct {
		Balance  int64
		Received int64
		Sent     int64
	}
}

func (c *Chain) addressURL(hashes []string) string {
	return fmt.Sprintf("%s/%s/addresses/%s",
		baseURL, c.network, strings.Join(hashes, ","))
}

// GetAddressMulti allows you to get multiple addresses with one API call.
// It returns basic balance details.
//
// Chain documentation can be found here
// https://chain.com/docs#bitcoin-address
func (c *Chain) GetAddressMulti(hashes []string) ([]Address, error) {
	if len(hashes) > MaxAddresses {
		return nil, fmt.Errorf("max addresses allowed is %d", MaxAddresses)
	}

	url, addresses := c.addressURL(hashes), make([]Address, len(hashes))
	return addresses, c.httpGetJSON(url, &addresses)
}

// GetAddress allows you to get one address. It returns basic balance details.
//
// Chain documentation can be found here
// https://chain.com/docs#bitcoin-address.
func (c *Chain) GetAddress(hash string) (Address, error) {
	url, addresses := c.addressURL([]string{hash}), make([]Address, 1)
	return addresses[0], c.httpGetJSON(url, &addresses)
}

func (c *Chain) addressTransactionsURL(hashes []string, limit int) string {
	return fmt.Sprintf("%s/%s/addresses/%s/transactions?limit=%d",
		baseURL, c.network, strings.Join(hashes, ","), limit)
}

// GetAddressTransactionsMulti returns a set of transactions for one or more
// Bitcoin addresses.
//
// Chain documentation can be found here
// https://chain.com/docs#bitcoin-address-transactions.
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

// GetAddressTransactions returns a set of transactions for one Bitcoin
// address.
//
// Chain documentation can be found here
// https://chain.com/docs#bitcoin-address-transactions.
func (c *Chain) GetAddressTransactions(
	hash string, limit int) ([]Transaction, error) {
	return c.GetAddressTransactionsMulti([]string{hash}, limit)
}

func (c *Chain) addressUnspentOutputsURL(hashes []string) string {
	return fmt.Sprintf("%s/%s/addresses/%s/unspents",
		baseURL, c.network, strings.Join(hashes, ","))
}

// GetAddressUnspentOutputsMulti returns a collection of unspent outputs for
// several Bitcoin addresses. These outputs can be used as inputs for
// a new transaction.
//
// Chain documentation can be found here
// https://chain.com/docs#bitcoin-address-unspents.
func (c *Chain) GetAddressUnspentOutputsMulti(
	hashes []string) ([]Output, error) {

	if len(hashes) > MaxAddresses {
		return nil, fmt.Errorf("max addresses allowed is %d", MaxAddresses)
	}

	url := c.addressUnspentOutputsURL(hashes)
	outputs := make([]Output, len(hashes))
	return outputs, c.httpGetJSON(url, &outputs)
}

// GetAddressUnspentOutputs returns a collection of unspent outputs for a
// Bitcoin address. These outputs can be used as inputs for a new transaction.
//
// Chain documentation can be found here
// https://chain.com/docs#bitcoin-address-unspents.
func (c *Chain) GetAddressUnspentOutputs(hash string) ([]Output, error) {
	return c.GetAddressUnspentOutputsMulti([]string{hash})
}
