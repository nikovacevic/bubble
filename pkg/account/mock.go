package account

import (
	"fmt"
	"log"
	"sync"

	"github.com/nikovacevic/bubble/pkg/bubble"
)

// MockStore implements the Store interface for testing
type MockStore struct {
	sync.RWMutex
	accounts map[int]*bubble.Account
	nextID   int
}

// NewMockStore creates a new MockStore with 10 hard-coded accounts
func NewMockStore() *MockStore {
	return &MockStore{
		accounts: map[int]*bubble.Account{},
		nextID:   1,
	}
}

// Seed saves a set of mock Accounts
func (ms *MockStore) Seed() *MockStore {
	accounts := []*bubble.Account{
		bubble.NewAccount("adam@getbubble.org", "Adam"),
		bubble.NewAccount("beth@getbubble.org", "Beth"),
		bubble.NewAccount("chip@getbubble.org", "Chip"),
		bubble.NewAccount("dave@getbubble.org", "Dave"),
		bubble.NewAccount("elen@getbubble.org", "Elen"),
		bubble.NewAccount("faye@getbubble.org", "Faye"),
		bubble.NewAccount("gary@getbubble.org", "Gary"),
		bubble.NewAccount("hugh@getbubble.org", "Hugh"),
		bubble.NewAccount("inga@getbubble.org", "Inga"),
		bubble.NewAccount("jake@getbubble.org", "Jake"),
	}

	for _, account := range accounts {
		err := ms.SaveAccount(account)
		if err != nil {
			log.Fatalf("MockStore.Seed: %s", err)
		}
	}

	return ms
}

// ListAccounts returns the Accounts in the MockStore
func (ms *MockStore) ListAccounts() ([]*bubble.Account, error) {
	ms.RLock()
	defer ms.RUnlock()

	accounts := []*bubble.Account{}
	for _, account := range ms.accounts {
		accounts = append(accounts, account)
	}

	return accounts, nil
}

// FindAccount should return the Account belonging to the given ID
func (ms *MockStore) FindAccount(id int) (*bubble.Account, error) {
	ms.RLock()
	defer ms.RUnlock()

	if account, ok := ms.accounts[id]; ok {
		return account, nil
	}

	return nil, fmt.Errorf("no account with ID: %d", id)
}

// SaveAccount should insert or update the given Account
func (ms *MockStore) SaveAccount(account *bubble.Account) error {
	ms.Lock()
	defer ms.Unlock()

	if account == nil {
		return fmt.Errorf("cannot save nil account")
	}

	if account.ID > 0 {
		ms.accounts[account.ID] = account
	} else {
		account.ID = len(ms.accounts)
		ms.accounts[account.ID] = account
	}

	return nil
}
