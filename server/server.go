package main

import (
	"encoding/gob"
	"fmt"
	"mine-it/blockchain"
	"net"
)

func handleConnection(conn net.Conn, bc *blockchain.BlockChain) {
	dec := gob.NewDecoder(conn)

	var transactions []blockchain.Transaction
	dec.Decode(&transactions)

	fmt.Printf("Received : %+v", transactions)

	bc.AddBlock(transactions)

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

	conn.Close()
}

func main() {
	fmt.Println("start")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}

	bc := blockchain.NewBlockchain()

	for {
		conn, err := ln.Accept() // this blocks until connection or error
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(conn, bc) // a goroutine handles conn so that the loop can accept other connections
	}
}
