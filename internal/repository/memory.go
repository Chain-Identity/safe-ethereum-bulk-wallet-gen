package repository

import (
	"sync"

	"github.com/chain-identity/walletGenerator/internal/wallet"
)

type InMemoryRepository struct {
	walletsMu sync.Mutex
	wallets   []*wallet.Wallet
}

func NewInMemoryRepository() Repository {
	return &InMemoryRepository{
		wallets: make([]*wallet.Wallet, 0),
	}
}

func (r *InMemoryRepository) Insert(wallet *wallet.Wallet) error {
	r.walletsMu.Lock()
	defer r.walletsMu.Unlock()
	r.wallets = append(r.wallets, wallet)
	return nil
}

func (r *InMemoryRepository) Result() []*wallet.Wallet {
	return r.wallets
}

func (r *InMemoryRepository) Close() error {
	return nil
}
