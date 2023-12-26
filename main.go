package main

import (
	"fmt"
	"mine-it/blockchain"
	"mine-it/wallet"
)

func main() {

	bc := blockchain.NewBlockchain()

	samwallet := wallet.NewWallet("Sam")

	jjwallet := wallet.NewWallet("JJ")

	transactions := []blockchain.Transaction{
		{Sender: samwallet.Name, Recipient: jjwallet.Name, Amount: 1.0, Coin: "Bitcoin"},
		{Sender: jjwallet.Name, Recipient: samwallet.Name, Amount: 0.01, Coin: "Ethereum"},
	}

	bc.AddBlock(transactions)

	transactions2 := []blockchain.Transaction{
		{Sender: samwallet.Name, Recipient: jjwallet.Name, Amount: 0.01, Coin: "Bitcoin"},
		{Sender: jjwallet.Name, Recipient: samwallet.Name, Amount: 0.01, Coin: "Ethereum"},
	}

	bc.AddBlock(transactions2)

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		for _, tx := range block.Transactions {
			fmt.Printf("  Sender: %s\n", tx.Sender)
			fmt.Printf("  Recipient: %s\n", tx.Recipient)
			fmt.Printf("  Amount: %f %s\n", tx.Amount, tx.Coin)
		}
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

	samwallet.PrintWallet()
	jjwallet.PrintWallet()
}
