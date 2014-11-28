package chain_test

import (
	"testing"

	"github.com/qedus/chain"
)

func TestGetAddress(t *testing.T) {
	c := newChain(t, chain.TestNet3)

	a, err := c.GetAddress("msk1uz21sUAXdmgqUiWvkRBLNfL1SXatyj")
	if err != nil {
		t.Fatal(err)
	}

	if a.Total.Received < 34728440 {
		t.Fatal("incorrect received value")
	}
	if a.Total.Sent < 0 {
		t.Fatal("incorrect received value")
	}
	if a.Total.Received-a.Total.Sent != a.Total.Balance {
		t.Fatal("incorrect balance")
	}
}

func TestGetAddressMulti(t *testing.T) {
	hashes := []string{
		"msk1uz21sUAXdmgqUiWvkRBLNfL1SXatyj",
		"n4CyDypGn7jyfKamweA26gQyJGm2HwWbmE",
	}

	c := newChain(t, chain.TestNet3)
	addrs, err := c.GetAddressMulti(hashes)
	if err != nil {
		t.Fatal(err)
	}

	if addrs[0].Total.Received < 34728440 {
		t.Fatal("incorrect received value")
	}
	if addrs[0].Total.Sent < 0 {
		t.Fatal("incorrect received value")
	}
	if addrs[0].Total.Received-addrs[0].Total.Sent != addrs[0].Total.Balance {
		t.Fatal("incorrect balance")
	}

	if addrs[1].Total.Received < 30289051865 {
		t.Fatal("incorrect received value")
	}
	if addrs[1].Total.Sent < 30282061865 {
		t.Fatal("incorrect received value")
	}
	if addrs[1].Total.Received-addrs[1].Total.Sent != addrs[1].Total.Balance {
		t.Fatal("incorrect balance")
	}
}

func TestGetAddressError(t *testing.T) {
	c := newChain(t, chain.TestNet3)
	if _, err := c.GetAddress("fake address"); err == nil {
		t.Fatal("expected an error")
	} else {
		t.Log(err)
	}
}

func TestGetAddressTransactions(t *testing.T) {
	c := newChain(t, chain.TestNet3)
	txns, err := c.GetAddressTransactions(
		"n4CyDypGn7jyfKamweA26gQyJGm2HwWbmE", 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(txns) != 1 {
		t.Fatal("transactions != 1", len(txns))
	}
}
