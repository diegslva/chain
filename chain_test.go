package chain_test

import (
	"net/http"

	"github.com/intusco/chain"
)

func newTestChain() *chain.Chain {
	return chain.New(http.DefaultClient, chain.TestNet3,
		"{API-KEY-ID}", "{API-KEY-SECRET}")
}
