package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Number        int
	PrevBlockHash string
	Payload       string
	Hash          string
}

func (b *Block) CalcHash() string {
	hashData := fmt.Sprintf("%d%s%s", b.Number, b.PrevBlockHash, b.Payload)
	hash := sha256.Sum256([]byte(hashData))
	return fmt.Sprintf("%x", hash)
}

func (b *Block) Mine() string {
	b.Hash = b.CalcHash()
	return b.Hash
}

func GetGenesisBlock() Block {
	return Block{
		Number:        GENESIS_BLOCK_NUMBER,
		PrevBlockHash: GENESIS_BLOCK_PREV_HASH,
		Payload:       GENESIS_BLOCK_PAYLOAD,
		Hash:          GENESIS_BLOCK_HASH,
	}
}
