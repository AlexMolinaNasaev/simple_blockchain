package blockchain

import "testing"

func TestCalcHash(t *testing.T) {
	testBlock := Block{
		Number:        0,
		PrevBlockHash: "genesis",
		Payload:       "hello!",
	}

	expectedTestBlockHash := "74100db7065ebcde033b04141e13911e3fdd638cc7259c2ca7493a2c2e52d725"
	testBlockHash := testBlock.CalcHash()
	if testBlockHash != expectedTestBlockHash {
		t.Fatalf("wrong block hash\n Expected: %v\n Got:      %v", expectedTestBlockHash, testBlockHash)
	}
}

// As genesis block is hardcoded test checks genesis block pre-calculated const hash
func TestGetGenesisBlock(t *testing.T) {
	genesis := GetGenesisBlock()
	genesis.Mine()

	if GENESIS_BLOCK_HASH != genesis.Hash {
		t.Errorf("wrong block hash\n Expected: %v\n Got:      %v", GENESIS_BLOCK_HASH, genesis.Hash)
	}
}
