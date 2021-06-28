package blockchain

import (
	"fmt"
)

type Chain struct {
	ID     int8
	Blocks []*Block
}

func NewChain(ID int8) *Chain {
	return &Chain{
		ID:     ID,
		Blocks: make([]*Block, 0),
	}
}

func (c *Chain) AddBlock(b *Block) error {
	chainLen := len(c.Blocks)
	if chainLen != b.Number {
		return fmt.Errorf("block number: %d does not match chain lenght: %d", b.Number, chainLen)
	}

	b.CalcHash()

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
