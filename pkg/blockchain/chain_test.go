package blockchain

import (
	"math/rand"
	"testing"
	"time"
)

var expectedChain Chain = Chain{
	ID: 255,
	Blocks: []*Block{
		{
			Number:        0,
			PrevBlockHash: "",
			Payload:       "hello!",
			Hash:          "1654f5178184945cd5210f05def489fc087fc5c09be586261a0e65126d6a196b",
		},
		{
			Number:        1,
			PrevBlockHash: "1654f5178184945cd5210f05def489fc087fc5c09be586261a0e65126d6a196b",
			Payload:       "world!",
			Hash:          "3ff1be2be0fd129d4190c769e38cac3ca3f643850fa94107f22440b8e4828ea0",
		},
		{
			Number:        2,
			PrevBlockHash: "3ff1be2be0fd129d4190c769e38cac3ca3f643850fa94107f22440b8e4828ea0",
			Payload:       "test!",
			Hash:          "a81c83a869fc51a4771c18601c0070d5dc714004e9c3c5f6eb5fbc6a2990f4f4",
		},
		{
			Number:        3,
			PrevBlockHash: "a81c83a869fc51a4771c18601c0070d5dc714004e9c3c5f6eb5fbc6a2990f4f4",
			Payload:       "this is a transaction",
			Hash:          "d8c5ea4626230b7de9f665fd49c670bc4679510540afae055c3426812251e190",
		},
		{
			Number:        4,
			PrevBlockHash: "d8c5ea4626230b7de9f665fd49c670bc4679510540afae055c3426812251e190",
			Payload:       "here!",
			Hash:          "bf77dea3e6389c7f3acb31eafd7fe7df9acffdb58489f44edc9cddef70f5e874",
		},
	},
}

func TestAddBlock(t *testing.T) {
	testChain := NewChain(255)
	testChain.AddBlock("hello!")
	testChain.AddBlock("world!")
	testChain.AddBlock("test!")
	testChain.AddBlock("this is a transaction")
	testChain.AddBlock("here!")

	expectedChainLen := len(expectedChain.Blocks)
	testChainLen := len(testChain.Blocks)

	if expectedChainLen != testChainLen {
		t.Fatalf("wrong chain lenght\n  Expected: %v\n Got:      %v", expectedChainLen, testChainLen)
	}

	// chosing random block pair to compare
	rand.Seed(time.Now().UTC().UnixNano())
	randBlockNum := rand.Intn(expectedChainLen)
	expectedRandBlock := expectedChain.Blocks[randBlockNum]
	testRandBlock := testChain.Blocks[randBlockNum]

	if expectedRandBlock.Number != testRandBlock.Number {
		t.Errorf("wrong block number\n random chosen block %v\n Expected: %v\n Got:      %v",
			randBlockNum, expectedRandBlock.Number, testRandBlock.Number)
	}

	if expectedRandBlock.PrevBlockHash != testRandBlock.PrevBlockHash {
		t.Errorf("wrong block PrevBlockHash\n random chosen block %v\n Expected: %v\n Got:      %v",
			randBlockNum, expectedRandBlock.PrevBlockHash, testRandBlock.PrevBlockHash)
	}

	if expectedRandBlock.Payload != testRandBlock.Payload {
		t.Errorf("wrong block payload\n random chosen block %v\n Expected: %v\n Got:      %v",
			randBlockNum, expectedRandBlock.Payload, testRandBlock.Payload)
	}

	if expectedRandBlock.Hash != testRandBlock.Hash {
		t.Errorf("wrong block hash\n random chosen block %v\n Expected: %v\n Got:      %v",
			randBlockNum, expectedRandBlock.Hash, testRandBlock.Hash)
	}
}

func TestValidate(t *testing.T) {
	testChain := expectedChain

	err := testChain.Validate()
	if err != nil {
		t.Errorf("validation error: %v", err.Error())
	}

	rand.Seed(time.Now().UTC().UnixNano())
	randBlockNum := rand.Intn(len(expectedChain.Blocks))

	// breaking random block
	expectedRandBlock := *testChain.Blocks[randBlockNum]
	testRandBlock := testChain.Blocks[randBlockNum]
	testRandBlock.Number = -1
	testRandBlock.Hash = "wrong hash"

	err = testChain.Validate()
	if err == nil {
		t.Errorf("wrong block number not catched\n random chosen block %v\n Expected: %v\n Got:      %v",
			randBlockNum, expectedRandBlock.Number, testRandBlock.Number)
	}

	testRandBlock.Number = randBlockNum

	err = testChain.Validate()
	if err == nil {
		t.Errorf("wrong block hash not catched\n random chosen block %v\n Expected: %v\n Got:      %v",
			randBlockNum, expectedRandBlock.Hash, testRandBlock.Hash)
	}
}
