package wallet

import (
	"crypto/ecdsa"
	"encoding/hex"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

type Wallet struct {
	Address    string
	PrivateKey string
}

// NewFromPrivatekey returns a new wallet from a given private key.
func NewFromPrivatekey(privateKey *ecdsa.PrivateKey) (*Wallet, error) {
	if privateKey == nil {
		return nil, errors.New("private key is nil")
	}

	publicKey := &privateKey.PublicKey

	// toString PrivateKey
	priveKeyBytes := crypto.FromECDSA(privateKey)
	privHex := make([]byte, len(priveKeyBytes)*2)
	hex.Encode(privHex, priveKeyBytes)
	privString := b2s(privHex)

	// toString PublicKey
	publicKeyBytes := crypto.Keccak256(crypto.FromECDSAPub(publicKey)[1:])[12:]
	if len(publicKeyBytes) > common.AddressLength {
		publicKeyBytes = publicKeyBytes[len(publicKeyBytes)-common.AddressLength:]
	}
	pubHex := make([]byte, len(publicKeyBytes)*2+2)
	copy(pubHex[:2], "0x")
	hex.Encode(pubHex[2:], publicKeyBytes)
	pubString := b2s(pubHex)

	return &Wallet{
		Address:    pubString,
		PrivateKey: privString,
	}, nil
}

// b2s converts a byte slice to a string without memory allocation.
func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// PrivateKeyGenerator provides a fn that creates wallets
func PrivateKeyGenerator() func() (*Wallet, error) {
	return func() (*Wallet, error) {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			return nil, errors.WithStack(err)
		}

		wallet, err := NewFromPrivatekey(privateKey)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return wallet, nil
	}
}
