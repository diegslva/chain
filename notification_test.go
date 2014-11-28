package chain_test

import (
	"testing"

	"github.com/qedus/chain"
)

func TestNotifications(t *testing.T) {
	c := newChain(t, chain.TestNet3)

	resp, err := c.CreateNewTxNotification("https://localhost.com")
	if err != nil {
		t.Fatal(err)
	}

	responses, err := c.ListNotifications()
	if err != nil {
		t.Fatal(err)
	}

	if len(responses) != 1 {
		t.Fatal("expected one notification")
	}

	if resp.ID != responses[0].ID {
		t.Fatal("expected same ID")
	}

	resp, err = c.DeleteNotification(resp.ID)
	if err != nil {
		t.Fatal(err)
	}
	if resp.ID != responses[0].ID {
		t.Fatal("expected same ID")
	}

	responses, err = c.ListNotifications()
	if err != nil {
		t.Fatal(err)
	}

	if len(responses) != 0 {
		t.Fatal("expected no notifications")
	}
}
