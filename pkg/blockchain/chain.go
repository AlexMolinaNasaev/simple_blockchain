package blockchain

import (
	"fmt"
)

type Chain struct {
	ID     uint8
	Blocks []*Block
}

func NewChain(ID uint8) *Chain {
	return &Chain{
		ID:     ID,
		Blocks: make([]*Block, 0),
	}
}

func (c *Chain) AddBlock(payload string) error {
	chainLen := len(c.Blocks)
	b := &Block{
		Number:  chainLen,
		Payload: payload,
	}

	if chainLen == 0 {
		b.PrevBlockHash = ""
	} else {
		b.PrevBlockHash = c.Blocks[chainLen-1].Hash
	}

	b.Mine()

	c.Blocks = append(c.Blocks, b)

	return nil
}

func (c *Chain) Validate() error {
	if len(c.Blocks) == 0 {
		return nil
	}

	for i, b := range c.Blocks {
		if i != b.Number {
			return fmt.Errorf("wrong block number: %d at height: %d", b.Number, i)
		}

		if b.CalcHash() != b.Hash {
			return fmt.Errorf("wrong hash at block: %d", b.Number)
		}
	}

	return nil
}
