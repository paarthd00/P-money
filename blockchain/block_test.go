package blockchain

import (
	"testing"
)

func TestNewGenesisBlock(t *testing.T) {

	testInitBlock := NewGenesisBlock()

	if testInitBlock.Hash == nil {
		t.Fatalf("Expected: Genesis Block, Got: %s", testInitBlock.Hash)
	}

}
