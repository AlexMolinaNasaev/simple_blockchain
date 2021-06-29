package blockchain

import "fmt"

const (
	UnknownError          = -1
	WrongBlockNumberError = iota
	WrongBlockHashError
)

var ErrorMessages = map[int]string{
	UnknownError:          "unkown error",
	WrongBlockNumberError: "wrong block number",
	WrongBlockHashError:   "wrong block hash",
}

func NewChainValidationError(errorCode int, err error) error {
	if _, ok := ErrorMessages[errorCode]; !ok {
		return fmt.Errorf(ErrorMessages[UnknownError])
	}

	return fmt.Errorf("%s: %s", ErrorMessages[errorCode], err.Error())
}
