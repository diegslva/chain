package chain

import (
	"fmt"
)

type Block struct {
	Hash              string
	PreviousBlockHash string `json:"previous_block_hash"`
	Height            int64
	Confirmations     int64
	MerkleRoot        string `json:"merkle_root"`
	Time              string
	Nonce             uint32
	Difficulty        float64
	TransactionHashes []string `json:"transaction_hashes"`
}

func (c *Chain) BlockByHash(hash string) (Block, error) {
	url, block := fmt.Sprintf("%s/blocks/%s", c.network, hash), Block{}
	return block, c.httpGetJSON(url, &block)
}

func (c *Chain) BlockByHeight(height uint64) (Block, error) {
	url, block := fmt.Sprintf("%s/blocks/%d", c.network, height), Block{}
	return block, c.httpGetJSON(url, &block)
}

func (c *Chain) LatestBlock() (Block, error) {
	url, block := fmt.Sprintf("%s/blocks/latest", c.network), Block{}
	return block, c.httpGetJSON(url, &block)
}
