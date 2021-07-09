package blockchain

import (
	"fmt"
)

const (
	UnknownError = -1
	// block errors
	WrongBlockNumberError = iota
	NegativeBlockNumberError
	ExcitingBlockNumberError
	WrongPrevBlockHashError
	WrongBlockHashError
	EmptyBlockHashError
	BlockValidationError
	ChainValidationError
	EmptyChainError
	MineBlockError
	PeerSyncError
	AddBlockError
)

var ErrorMessages = map[int]string{
	UnknownError:             "unkown error",
	WrongBlockNumberError:    "wrong block number",
	NegativeBlockNumberError: "negative block number",
	ExcitingBlockNumberError: "block number exceeds chain lenght",
	WrongPrevBlockHashError:  "wrong previous block hash",
	WrongBlockHashError:      "wrong block hash",
	EmptyBlockHashError:      "block hash is empty",
	BlockValidationError:     "block validation error",
	ChainValidationError:     "chain validation error",
	EmptyChainError:          "chain is empty",
	MineBlockError:           "cannot mine block",
	PeerSyncError:            "cannot synchronize to chain",
	AddBlockError:            "cannot add block to chain",
}

func NewBlockchainChainError(errorCode int, err error) error {
	if _, ok := ErrorMessages[errorCode]; !ok {
		return fmt.Errorf(ErrorMessages[UnknownError])
	}

	if err == nil {
		err = fmt.Errorf("")
	}

	return fmt.Errorf("%s: %s", ErrorMessages[errorCode], err.Error())
}
