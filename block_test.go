package chain_test

import (
	"testing"
)

func TestBlockByHash(t *testing.T) {
	c := newTestChain()
	hash := "0000000086907a79fb7f040893a332200df4580fe6a83b0ffe47f3527a5f753f"
	block, err := c.BlockByHash(hash)
	if err != nil {
		t.Fatal(err)
	}

	if len(block.TransactionHashes) == 0 {
		t.Fatal("no transaction hashes")
	}
}

func TestBlockByHeight(t *testing.T) {
	c := newTestChain()
	block, err := c.BlockByHeight(277469)
	if err != nil {
		t.Fatal(err)
	}

	if len(block.TransactionHashes) == 0 {
		t.Fatal("no transaction hashes")
	}
}

func TestLatestBlock(t *testing.T) {
	c := newTestChain()
	block, err := c.LatestBlock()
	if err != nil {
		t.Fatal(err)
	}

	if len(block.TransactionHashes) == 0 {
		t.Fatal("no transaction hashes")
	}
}