package blockchain

import (
	"bytes"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"mine-it/wallet"
	"strconv"
	"time"
)

type Wallet struct {
	PrivateKeys    map[string]*rsa.PrivateKey
	PublicKeys     map[string]*rsa.PublicKey
	BitcoinAddress map[string]string
	Coins          map[string]wallet.Coin
}

type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
	Coin      string
}

type Block struct {
	Hash         []byte
	Transactions []Transaction
	PrevHash     []byte
	Timestamp    int64
}

type BlockChain struct {
	Blocks []*Block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	var txData []byte
	for _, tx := range b.Transactions {
		txData = append(txData, []byte(fmt.Sprintf("%s:%s:%f:%s", tx.Sender, tx.Recipient, tx.Amount, tx.Coin))...)
	}
	headers := bytes.Join([][]byte{b.PrevHash, txData, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(transactions []Transaction, prevHash []byte) *Block {
	timestamp := time.Now().Unix()
	block := &Block{Hash: []byte{}, Transactions: transactions, PrevHash: prevHash, Timestamp: timestamp}
	block.SetHash()
	return block
}

func NewBlockchain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *Block {
	transactions := []Transaction{
		{Sender: "Sam", Recipient: "JJ", Amount: 1.0, Coin: "BitCoin"},
	}
	return NewBlock(transactions, []byte{})
}

func (bc *BlockChain) AddBlock(transactions []Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *BlockChain) AddWalletBlock(wallet *wallet.Wallet, transactions []Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}
