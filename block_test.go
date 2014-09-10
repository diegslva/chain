package chain_test

import (
	"testing"

	"github.com/qedus/chain"
)

func TestGetBlockByHash(t *testing.T) {
	c := newChain(t, chain.TestNet3)
	hash := "0000000086907a79fb7f040893a332200df4580fe6a83b0ffe47f3527a5f753f"
	block, err := c.GetBlockByHash(hash)
	if err != nil {
		t.Fatal(err)
	}

	if len(block.TransactionHashes) == 0 {
		t.Fatal("no transaction hashes")
	}
}

func TestGetBlockByHeight(t *testing.T) {
	c := newChain(t, chain.TestNet3)
	block, err := c.GetBlockByHeight(277469)
	if err != nil {
		t.Fatal(err)
	}

	if len(block.TransactionHashes) == 0 {
		t.Fatal("no transaction hashes")
	}
}

func TestGetLatestBlock(t *testing.T) {
	c := newChain(t, chain.TestNet3)
	block, err := c.GetLatestBlock()
	if err != nil {
		t.Fatal(err)
	}

	if len(block.TransactionHashes) == 0 {
		t.Fatal("no transaction hashes")
	}
}
