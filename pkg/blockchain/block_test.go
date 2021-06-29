package blockchain

import "testing"

func TestCalcHash(t *testing.T) {
	testBlock := Block{
		Number:        0,
		PrevBlockHash: "",
		Payload:       "hello!",
	}

	expectedTestBlockHash := "1654f5178184945cd5210f05def489fc087fc5c09be586261a0e65126d6a196b"
	testBlockHash := testBlock.CalcHash()
	if testBlockHash != expectedTestBlockHash {
		t.Fatalf("wrong block hash\n Expected: %v\n Got:      %v", expectedTestBlockHash, testBlockHash)
	}
}
