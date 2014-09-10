package chain_test

import (
	"net/http"

	"os"
	"testing"

	"github.com/qedus/chain"
)

func newChain(t *testing.T, net chain.Network) *chain.Chain {
	apiKeyID := os.Getenv("CHAIN_API_KEY_ID")
	apiKeySecret := os.Getenv("CHAIN_API_KEY_SECRET")

	if apiKeyID == "" {
		t.Fatal("CHAIN_API_KEY_ID environment variable must be set")
	}
	if apiKeySecret == "" {
		t.Fatal("CHAIN_API_KEY_SECRET environment variable must be set")
	}

	return chain.New(http.DefaultClient, net, apiKeyID, apiKeySecret)
}
