package wallet

import (
	"fmt"
)

type Wallet struct {
	PrivateKey     string
	PublicKey      string
	BitcoinAddress string
}

func newKeyPair() (string, string) {
	return "private", "public"
}

func generateAddress(publicKey string) string {
	return fmt.Sprintf("address for %s", publicKey)
}

func NewWallet() *Wallet {
	private, public := newKeyPair()
	bitcoinAddress := generateAddress(public)
	wallet := Wallet{private, public, bitcoinAddress}
	return &wallet
}
