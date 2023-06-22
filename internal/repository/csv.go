package repository

import (
	"encoding/csv"
	"os"
	"sync"

	"github.com/chain-identity/walletGenerator/internal/wallet"
)

type FileRepository struct {
	walletsMu sync.Mutex
	wallets   []*wallet.Wallet
	writer    csv.Writer
}

func NewFileRepository(path string) Repository {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		panic("cant open csv")
	}

	writer := csv.NewWriter(f)
	return &FileRepository{
		wallets: make([]*wallet.Wallet, 0),
		writer:  *writer,
	}
}

func (r *FileRepository) Insert(wallet *wallet.Wallet) error {
	r.walletsMu.Lock()
	defer r.walletsMu.Unlock()

	return r.writer.Write([]string{wallet.Address, wallet.PrivateKey})
}

func (r *FileRepository) Result() []*wallet.Wallet {
	return r.wallets
}

func (r *FileRepository) Close() error {
	r.writer.Flush()
	return r.writer.Error()
}
