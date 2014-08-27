package chain_test

import (
	"testing"
)

func TestGetAddress(t *testing.T) {
	c := newTestChain(t)

	a, err := c.GetAddress("msk1uz21sUAXdmgqUiWvkRBLNfL1SXatyj")
	if err != nil {
		t.Fatal(err)
	}

	if a.Received < 34728440 {
		t.Fatal("incorrect received value")
	}
	if a.Sent < 0 {
		t.Fatal("incorrect received value")
	}
	if a.Received-a.Sent != a.Balance {
		t.Fatal("incorrect balance")
	}
}

func TestGetAddressMulti(t *testing.T) {
	hashes := []string{
		"msk1uz21sUAXdmgqUiWvkRBLNfL1SXatyj",
		"n4CyDypGn7jyfKamweA26gQyJGm2HwWbmE",
	}

	c := newTestChain(t)
	addrs, err := c.GetAddressMulti(hashes)
	if err != nil {
		t.Fatal(err)
	}

	if addrs[0].Received < 34728440 {
		t.Fatal("incorrect received value")
	}
	if addrs[0].Sent < 0 {
		t.Fatal("incorrect received value")
	}
	if addrs[0].Received-addrs[0].Sent != addrs[0].Balance {
		t.Fatal("incorrect balance")
	}

	if addrs[1].Received < 30289051865 {
		t.Fatal("incorrect received value")
	}
	if addrs[1].Sent < 30282061865 {
		t.Fatal("incorrect received value")
	}
	if addrs[1].Received-addrs[1].Sent != addrs[1].Balance {
		t.Fatal("incorrect balance")
	}
}

func TestGetAddressError(t *testing.T) {
	c := newTestChain(t)
	if _, err := c.GetAddress("fake address"); err == nil {
		t.Fatal("expected an error")
	} else {
		t.Log(err)
	}
}

func TestGetAddressTransactions(t *testing.T) {
	c := newTestChain(t)
	txns, err := c.GetAddressTransactions(
		"n4CyDypGn7jyfKamweA26gQyJGm2HwWbmE", 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(txns) != 1 {
		t.Fatal("transactions != 1", len(txns))
	}
}
