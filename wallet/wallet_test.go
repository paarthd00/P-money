package wallet

import (
	"testing"
)

func TestNewWalletFields(t *testing.T) {
	testWallet := NewWallet("Gladys")
	name := testWallet.Name
	if name != "Gladys" {
		t.Fatalf("Expected: Gladys, Got: %s", name)
	}

	if testWallet.CoinAddress["Bitcoin"] == "" {
		t.Fatalf("Expected: Bitcoin address, Got: %s", testWallet.CoinAddress["Bitcoin"])
	}

	if testWallet.PrivateKey == "" {
		t.Fatalf("Expected: Private key, Got: %s", testWallet.PrivateKey)
	}

	if testWallet.PublicKey == "" {
		t.Fatalf("Expected: Public key, Got: %s", testWallet.PublicKey)
	}
}
