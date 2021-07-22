package blockchain

import (
	"fmt"
)

type Chain struct {
	ID     uint8
	Blocks []Block
}

// NewChain creates a new chain with pre-calculated genesis block
func NewChain(ID uint8) *Chain {
	return &Chain{
		ID:     ID,
		Blocks: []Block{GetGenesisBlock()},
	}
}

func (c *Chain) GetBlock(blockNum int) Block {
	return c.Blocks[blockNum]
}

func (c *Chain) ValidateChain() (int, error) {
	if len(c.Blocks) == 0 {
		return -1, nil
	}

	for i, b := range c.Blocks {
		if i != b.Number {
			return i, NewBlockchainChainError(BlockValidationError,
				NewBlockchainChainError(WrongBlockNumberError, fmt.Errorf("block number %d, block place: %d", b.Number, i)))
		}

		if err := c.ValidateBlock(i); err != nil {
			return i, NewBlockchainChainError(BlockValidationError, err)
		}
	}

	return -1, nil
}

func (c *Chain) ValidateBlock(blockNum int) error {
	if blockNum == GENESIS_BLOCK_NUMBER {
		if c.Blocks[blockNum].CalcHash() != GetGenesisBlock().Hash {
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

	prevBlock := c.GetBlock(blockNum - 1)

	validationBlock := Block{
		Number:        blockNum,
		PrevBlockHash: prevBlock.Hash,
		Payload:       c.GetBlock(blockNum).Payload,
	}

	validationBlock.Hash = validationBlock.CalcHash()

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

	prevBlock := c.GetBlock(chainLen - 1)

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

func (c *Chain) MineBlock(payload string) {
	chainLen := len(c.Blocks)
	b := Block{
		Number:        chainLen,
		PrevBlockHash: c.GetBlock(chainLen - 1).Hash,
		Payload:       payload,
	}
	b.Mine()

	c.Blocks = append(c.Blocks, b)
}

func (c *Chain) AddBlock(block Block) error {
	err := c.ValidateNewBlock(block)
	if err != nil {
		return NewBlockchainChainError(AddBlockError, err)
	}

	c.Blocks = append(c.Blocks, block)

	return nil
}

func (c *Chain) Print() {
	for _, b := range c.Blocks {
		fmt.Printf("Block[%d] hash: %s\n", b.Number, b.Hash)
	}
}
