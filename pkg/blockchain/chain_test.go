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

// func TestValidateChain(t *testing.T) {
// 	// copying expected chain to avoid data rewriting
// 	testChain := expectedChain

// 	if err := testChain.ValidateChain(); err != nil {
// 		t.Errorf("chain validation error: %v\nExpected no error", err.Error())
// 	}

// 	rand.Seed(time.Now().UTC().UnixNano())
// 	randBlockNum := rand.Intn(len(expectedChain.Blocks))

// 	// breaking random block
// 	expectedRandBlock := testChain.Blocks[randBlockNum]
// 	testRandBlock := testChain.Blocks[randBlockNum]
// 	testRandBlock.Number = -1
// 	testRandBlock.Hash = "wrong hash"

// 	if err := testChain.ValidateChain(); err != nil {
// 		t.Errorf("wrong block number not catched\n random chosen block %v\n Expected: %v\n Got:      %v",
// 			randBlockNum, expectedRandBlock.Number, testRandBlock.Number)
// 	}

// 	testRandBlock.Number = randBlockNum

// 	if err := testChain.ValidateChain(); err != nil {
// 		t.Errorf("wrong block hash not catched\n random chosen block %v\n Expected: %v\n Got:      %v",
// 			randBlockNum, expectedRandBlock.Hash, testRandBlock.Hash)
// 	}
// }

// func TestMineBlock(t *testing.T) {
// 	testChain := NewChain(255)
// 	for i := range expectedChain.Blocks {
// 		testChain.MineBlock(expectedChain.Blocks[i])
// 	}

// 	expectedChainLen := len(expectedChain.Blocks)
// 	testChainLen := len(testChain.Blocks)

// 	if expectedChainLen != testChainLen {
// 		t.Fatalf("wrong chain lenght\n  Expected: %v\n Got:      %v", expectedChainLen, testChainLen)
// 	}

// 	// chosing random block pair to compare
// 	rand.Seed(time.Now().UTC().UnixNano())
// 	randBlockNum := rand.Intn(expectedChainLen)
// 	expectedRandBlock := expectedChain.Blocks[randBlockNum]
// 	testRandBlock := testChain.Blocks[randBlockNum]

// 	if expectedRandBlock.Number != testRandBlock.Number {
// 		t.Errorf("wrong block number\n random chosen block %v\n Expected: %v\n Got:      %v",
// 			randBlockNum, expectedRandBlock.Number, testRandBlock.Number)
// 	}

// 	if expectedRandBlock.PrevBlockHash != testRandBlock.PrevBlockHash {
// 		t.Errorf("wrong block PrevBlockHash\n random chosen block %v\n Expected: %v\n Got:      %v",
// 			randBlockNum, expectedRandBlock.PrevBlockHash, testRandBlock.PrevBlockHash)
// 	}

// 	if expectedRandBlock.Payload != testRandBlock.Payload {
// 		t.Errorf("wrong block payload\n random chosen block %v\n Expected: %v\n Got:      %v",
// 			randBlockNum, expectedRandBlock.Payload, testRandBlock.Payload)
// 	}

// 	if expectedRandBlock.Hash != testRandBlock.Hash {
// 		t.Errorf("wrong block hash\n random chosen block %v\n Expected: %v\n Got:      %v",
// 			randBlockNum, expectedRandBlock.Hash, testRandBlock.Hash)
// 	}
// }
