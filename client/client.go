package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"mine-it/blockchain"
	"mine-it/wallet"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("start client")
	args := os.Args

	name := args[1]
	recipient := args[2]
	amount, err := strconv.ParseFloat(strings.TrimSpace(args[3]), 64)
	coin := args[4]

	if err != nil {
		log.Fatal("Connection error", err)
	}

	clientWallet := wallet.NewWallet(name)

	clientWallet.PrintWallet()

	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal("Connection error", err)
	}
	encoder := gob.NewEncoder(conn)

	transactions := &[]blockchain.Transaction{
		{Sender: name, Recipient: recipient, Amount: amount, Coin: coin},
	}

	fmt.Printf("Sending : %+v", transactions)
	encoder.Encode(transactions)
	conn.Close()
	fmt.Println("done")
}
