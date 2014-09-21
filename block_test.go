package chain_test

import (
	"testing"

	"github.com/qedus/chain"
)

func TestGetBlockByHash(t *testing.T) {
	c := newChain(t, chain.MainNet)
	hash := "000000000000000003dd5aa0232cc4e800295c348bc5ea3dc2f7db63c481d352"
	block, err := c.GetBlockByHash(hash)
	if err != nil {
		t.Fatal(err)
	}

	if block.Bits != "1824dbe9" {
		t.Fatalf("incorrect Bits")
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
