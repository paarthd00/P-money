package main

import (
	"fmt"
	"mine-it/blockchain"
	"mine-it/wallet"
)

func main() {
	bc := blockchain.NewBlockchain()

	// Create a new wallet
	w := wallet.NewWallet()

	// Add a block to the blockchain with the wallet's public key
	bc.AddWalletBlock(w, "Transaction data from wallet")
	bc.AddBlock("Send 1 P-money to Big J")
	// Print the blockchain
	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
