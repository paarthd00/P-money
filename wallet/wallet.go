package wallet

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type Coin struct {
	Symbol  string
	Balance float64
}

type Wallet struct {
	PrivateKey  string
	PublicKey   string
	CoinAddress map[string]string
	Name        string
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
	hash := sha256.New()
	hash.Write(x509.MarshalPKCS1PublicKey(publicKey))
	address := fmt.Sprintf("%x", hash.Sum(nil))
	return address
}

func NewWallet(name string) *Wallet {
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

	coinAddress := generateAddress(publicKey)
	wallet := Wallet{
		PrivateKey:  string(privateKeyPEM),
		PublicKey:   string(publicKeyPEM),
		CoinAddress: map[string]string{"Bitcoin": coinAddress},
		Name:        name,
	}
	return &wallet
}

func (w *Wallet) PrintWallet() {
	fmt.Println("Wallet")
	fmt.Println("Coin address: " + fmt.Sprintf("%v", w.CoinAddress))
	fmt.Println("Name: " + w.Name)
}
