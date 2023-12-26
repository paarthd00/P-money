package main

import (
    "fmt"
    "net"
    "encoding/gob"
	"mine-it/blockchain"
)

type P struct {
    M, N int64
}

func handleConnection(conn net.Conn) {
    dec := gob.NewDecoder(conn)
    transactions := &blockchain.Transaction{}
    dec.Decode(transactions)
    fmt.Printf("Received : %+v", transactions);

	conn.Close()
}

func main() {
    fmt.Println("start");
	ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        // handle error
    }

	// bc := blockchain.NewBlockchain()

    for {
        conn, err := ln.Accept() // this blocks until connection or error
        if err != nil {
            // handle error
            continue
        }
        go handleConnection(conn) // a goroutine handles conn so that the loop can accept other connections
    }
}
