package blockchain

import (
	"fmt"
)

type Chain struct {
	ID     uint8
	Blocks []Block
}

func NewChain(ID uint8) *Chain {
	return &Chain{
		ID:     ID,
		Blocks: []Block{GetGenesisBlock()},
	}
}

func (c *Chain) ValidateChain() error {
	if len(c.Blocks) == 0 {
		return nil
	}

	for i, b := range c.Blocks {
		if i != b.Number {
			return NewBlockchainChainError(WrongBlockNumberError,
				fmt.Errorf("block %d. chain length", b.Number))
		}

		if err := c.ValidateBlock(i); err != nil {
			return NewBlockchainChainError(WrongBlockNumberError,
				fmt.Errorf("block %d", b.Number))
		}
	}

	return nil
}

func (c *Chain) ValidateBlock(blockNum int) error {
	if blockNum == GENESIS_BLOCK_NUMBER {
		if c.Blocks[blockNum].Hash != GetGenesisBlock().Hash {
			return NewBlockchainChainError(WrongBlockHashError,
				fmt.Errorf("genesis block"))
		}
		return nil
	}

	if blockNum < 0 {
		return NewBlockchainChainError(NegativeBlockNumberError,
			fmt.Errorf("block number: %d", blockNum))
	}

	chainLen := len(c.Blocks)

	if blockNum > chainLen {
		return NewBlockchainChainError(ExcitingBlockNumberError,
			fmt.Errorf("block number: %d. Chain lenght:  %d", blockNum, chainLen))
	}

	prevBlock := c.Blocks[blockNum-1]

	validationBlock := Block{
		Number:        chainLen,
		PrevBlockHash: prevBlock.Hash,
		Payload:       c.Blocks[blockNum].Payload,
	}

	if c.Blocks[blockNum].Hash != validationBlock.CalcHash() {
		return NewBlockchainChainError(WrongBlockHashError,
			fmt.Errorf("block number: %d", blockNum))
	}

	return nil
}

func (c *Chain) ValidateNewBlock(block Block) error {
	if block.Hash == "" {
		return NewBlockchainChainError(EmptyBlockHashError,
			fmt.Errorf("block number: %d", block.Number))
	}

	chainLen := len(c.Blocks)

	if block.Number != chainLen {
		return NewBlockchainChainError(WrongBlockNumberError,
			fmt.Errorf("block number: %d, chain length: %d", block.Number, chainLen))
	}

	prevBlock := c.Blocks[chainLen-1]

	validationBlock := Block{
		Number:        chainLen,
		PrevBlockHash: prevBlock.Hash,
		Payload:       block.Payload,
	}

	if block.Hash != validationBlock.CalcHash() {
		return NewBlockchainChainError(WrongBlockHashError,
			fmt.Errorf("block: %d", block.Number))
	}

	return nil
}

func (c *Chain) MineBlock(block Block) error {
	err := c.ValidateNewBlock(block)
	if err != nil {
		return NewBlockchainChainError(MineBlockError, err)
	}

	c.Blocks = append(c.Blocks, block)

	return nil
}

func (c *Chain) Print() {
	for _, b := range c.Blocks {
		fmt.Printf("Block[%d] hash: %s\n", b.Number, b.Hash)
	}
}
