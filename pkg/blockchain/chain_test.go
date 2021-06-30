package blockchain

import (
	"math/rand"
	"testing"
	"time"
)

var payloads = []string{"hello!", "world!", "test", "this is a transaction", "here!"}

var expectedChain Chain = Chain{
	ID:     255,
	Blocks: make([]Block, len(payloads)+1), // 1 block per payload + genesis block
}

func init() {
	expectedChain.Blocks[0] = GetGenesisBlock()
	prevBlockHash := expectedChain.Blocks[0].Hash

	for i, p := range payloads {
		i = i + 1
		b := Block{
			Number:        i,
			PrevBlockHash: prevBlockHash,
			Payload:       p,
		}

		prevBlockHash = b.Mine()

		expectedChain.Blocks[i] = b
	}
}

func TestValidateBlock(t *testing.T) {}

func TestValidateNewBlock(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	// randBlockNum := rand.Intn(len(expectedChain.Blocks))
	expectedRandBlock := expectedChain.Blocks[1]

	if err := expectedChain.ValidateBlock(expectedRandBlock); err != nil {
		t.Errorf("Expected no error.\n Got: %v", err.Error())
	}

	// badBlock := expectedRandBlock

	// badBlock.Hash = ""
	// if err := expectedChain.ValidateBlock(badBlock); err == nil {
	// 	t.Errorf("Expected error: block hash is empty: block number: %v", randBlockNum)
	// }
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
