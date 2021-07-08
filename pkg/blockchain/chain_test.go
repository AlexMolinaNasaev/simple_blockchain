package blockchain

import (
	"testing"
)

const (
	ExpectedNoErrorTemplate = "Expected no error. Got: %v"
	GotNoErrorTemplate      = "Got no error. Expected error: %s"
	ExpectedErrorTemplate   = "\n Expected error: %v\n Got:            %v"
)

var payloads = []string{"hello!", "world!", "test", "this is a transaction", "here!"}

var expectedChain *Chain

func init() {
	expectedChain = NewChain(255)
	prevBlockHash := expectedChain.Blocks[0].Hash

	for i, p := range payloads {
		i = i + 1
		b := Block{
			Number:        i,
			PrevBlockHash: prevBlockHash,
			Payload:       p,
		}

		prevBlockHash = b.Mine()

		expectedChain.Blocks = append(expectedChain.Blocks, b)
	}
}

func TestValidateBlock(t *testing.T) {
	testChain := expectedChain
	if err := testChain.ValidateBlock(0); err != nil {
		t.Errorf(ExpectedNoErrorTemplate, err.Error())
	}

	testChain.Blocks[0].Hash = "wrong hash"
	expectedErrMsg := "wrong block hash: genesis block"
	if err := testChain.ValidateBlock(0); err == nil {
		t.Errorf(GotNoErrorTemplate, expectedErrMsg)
	} else if err.Error() != expectedErrMsg {
		t.Errorf(ExpectedErrorTemplate, expectedErrMsg, err.Error())
	}
	testChain.Blocks[0].Hash = GENESIS_BLOCK_HASH

	expectedErrMsg = "negative block number: block number: -1"
	if err := testChain.ValidateBlock(-1); err == nil {
		t.Errorf(GotNoErrorTemplate, expectedErrMsg)
	} else if err.Error() != expectedErrMsg {
		t.Errorf(ExpectedErrorTemplate, expectedErrMsg, err.Error())
	}

	expectedErrMsg = "block number exceeds chain lenght: block number: 9999999999999. Chain lenght:  6"
	if err := testChain.ValidateBlock(9999999999999); err == nil {
		t.Errorf(GotNoErrorTemplate, expectedErrMsg)
	} else if err.Error() != expectedErrMsg {
		t.Errorf(ExpectedErrorTemplate, expectedErrMsg, err.Error())
	}

	testChain.Blocks[2].Hash = GENESIS_BLOCK_HASH
	expectedErrMsg = "wrong block hash: block number: 2"
	if err := testChain.ValidateBlock(2); err == nil {
		t.Errorf(GotNoErrorTemplate, expectedErrMsg)
	} else if err.Error() != expectedErrMsg {
		t.Errorf(ExpectedErrorTemplate, expectedErrMsg, err.Error())
	}
	testChain.Blocks[2].Mine()
}

func TestValidateNewBlock(t *testing.T) {
	testChain := NewChain(255)

	brokenBlock := expectedChain.Blocks[1]
	brokenBlock.Hash = ""

	expectedErrMsg := "block hash is empty: block number: 1"
	if err := testChain.ValidateNewBlock(brokenBlock); err == nil {
		t.Errorf(GotNoErrorTemplate, expectedErrMsg)
	} else if err.Error() != expectedErrMsg {
		t.Errorf(ExpectedErrorTemplate, expectedErrMsg, err.Error())
	}

	if err := testChain.ValidateNewBlock(expectedChain.Blocks[1]); err != nil {
		t.Errorf(ExpectedNoErrorTemplate, err.Error())
	}

	brokenBlock.Hash = GENESIS_BLOCK_HASH

	expectedErrMsg = "wrong block hash: block: 1"
	if err := testChain.ValidateNewBlock(brokenBlock); err == nil {
		t.Errorf(GotNoErrorTemplate, expectedErrMsg)
	} else if err.Error() != expectedErrMsg {
		t.Errorf(ExpectedErrorTemplate, expectedErrMsg, err.Error())
	}
}

func TestValidateChain(t *testing.T) {
	// copying expected chain to avoid data rewriting
	testChain := expectedChain

	if err := testChain.ValidateChain(); err != nil {
		t.Errorf(ExpectedNoErrorTemplate, err.Error())
	}

	testChain.Blocks[3].Number = 5
	expectedErrMsg := "block validation error: wrong block number: block number 5, block place: 3"
	if err := testChain.ValidateChain(); err == nil {
		t.Errorf(GotNoErrorTemplate, expectedErrMsg)
	} else if err.Error() != expectedErrMsg {
		t.Errorf(ExpectedErrorTemplate, expectedErrMsg, err.Error())
	}
}

func TestMineBlock(t *testing.T) {
	chainlen := len(expectedChain.Blocks)
	payload := "mine me!"
	controllBlock := Block{
		Number:        chainlen,
		PrevBlockHash: expectedChain.Blocks[chainlen-1].Hash,
		Payload:       payload,
	}

	controllBlock.Mine()

	expectedChain.MineBlock(payload)

	if controllBlock.Hash != expectedChain.Blocks[chainlen].Hash {
		t.Errorf("wrong mined block hash:\n block[%d][controll] - %s\n block[%d][mined] - %s",
			chainlen, controllBlock.Hash, chainlen, expectedChain.Blocks[chainlen].Hash)
	}
}
