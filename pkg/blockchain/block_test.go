package blockchain

import "testing"

func TestCalcHash(t *testing.T) {
	testBlock := Block{
		Number:        GENESIS_BLOCK_NUMBER,
		PrevBlockHash: GENESIS_BLOCK_PREV_HASH,
		Payload:       GENESIS_BLOCK_PAYLOAD,
	}

	expectedTestBlockHash := GENESIS_BLOCK_HASH
	testBlockHash := testBlock.CalcHash()
	if testBlockHash != expectedTestBlockHash {
		t.Errorf("wrong block hash\n Expected: %v\n Got:      %v", expectedTestBlockHash, testBlockHash)
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
