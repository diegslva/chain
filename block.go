package chain

import (
	"fmt"
)

// Block represents a Bitcoin block.
//
// Chain documentation can be found here
// https://chain.com/docs#object-bitcoin-block.
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

// GetBlockByHash returns a Bitcoin block with the specified height.
//
// Chain documentation can be found here
// https://chain.com/docs#bitcoin-block.
func (c *Chain) GetBlockByHash(hash string) (Block, error) {
	url := fmt.Sprintf("%s/%s/blocks/%s", baseURL, c.network, hash)
	block := Block{}
	return block, c.httpGetJSON(url, &block)
}

// GetBlockByHeight returns a Bitcoin block at the specified height.
//
// Chain documentation can be found here
// https://chain.com/docs#bitcoin-block.
func (c *Chain) GetBlockByHeight(height uint64) (Block, error) {
	url, block := fmt.Sprintf("%s/%s/blocks/%d",
		baseURL, c.network, height), Block{}
	return block, c.httpGetJSON(url, &block)
}

// GetLatestBlock returns the latest Bitcoin block.
//
// Chain documentation can be found here
// https://chain.com/docs#bitcoin-block.
func (c *Chain) GetLatestBlock() (Block, error) {
	url, block := fmt.Sprintf("%s/%s/blocks/latest",
		baseURL, c.network), Block{}
	return block, c.httpGetJSON(url, &block)
}
