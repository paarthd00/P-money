package wallet

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type Wallet struct {
	PrivateKey     string
	PublicKey      string
	BitcoinAddress string
}

func newKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	bitSize := 4096
	key, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		panic(err)
	}

	return key, &key.PublicKey
}

func generateAddress(publicKey *rsa.PublicKey) string {
	// Hash the public key to simulate a Bitcoin address
	hash := sha256.New()
	hash.Write(x509.MarshalPKCS1PublicKey(publicKey))
	address := fmt.Sprintf("%x", hash.Sum(nil))
	return address
}

func NewWallet() *Wallet {
	privateKey, publicKey := newKeyPair()
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)

	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(publicKey),
		},
	)

	bitcoinAddress := generateAddress(publicKey)
	wallet := Wallet{string(privateKeyPEM), string(publicKeyPEM), bitcoinAddress}
	return &wallet
}
