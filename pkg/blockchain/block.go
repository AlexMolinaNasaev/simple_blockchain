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
	b.Hash = fmt.Sprintf("%x", hash)
	return b.Hash
}
