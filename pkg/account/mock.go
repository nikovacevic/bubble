package account

import (
	"fmt"
	"sort"

	"github.com/nikovacevic/bubble/pkg/bubble"
)

// MockStore implements the Store interface for testing
type MockStore struct {
	accounts map[int]*bubble.Account
}

// NewMockStore creates a new MockStore with 10 hard-coded accounts
func NewMockStore() *MockStore {
	accounts := []*bubble.Account{}
	accounts = append(accounts, bubble.NewAccount("adam@getbubble.org", "Adam"))
	accounts = append(accounts, bubble.NewAccount("beth@getbubble.org", "Beth"))
	accounts = append(accounts, bubble.NewAccount("chip@getbubble.org", "Chip"))
	accounts = append(accounts, bubble.NewAccount("dave@getbubble.org", "Dave"))
	accounts = append(accounts, bubble.NewAccount("elen@getbubble.org", "Elen"))
	accounts = append(accounts, bubble.NewAccount("faye@getbubble.org", "Faye"))
	accounts = append(accounts, bubble.NewAccount("gary@getbubble.org", "Gary"))
	accounts = append(accounts, bubble.NewAccount("hugh@getbubble.org", "Hugh"))
	accounts = append(accounts, bubble.NewAccount("inga@getbubble.org", "Inga"))
	accounts = append(accounts, bubble.NewAccount("jake@getbubble.org", "Jake"))

	accountMap := map[int]*bubble.Account{}
	for _, account := range accounts {
		accountMap[account.ID()] = account
	}

	return &MockStore{accounts: accountMap}
}

// List should provide a the given page from paginated list of Accounts
func (ms *MockStore) List(per, page int) ([]*bubble.Account, error) {
	accountsByID := map[int]*bubble.Account{}
	ids := []int{}

	for _, account := range ms.accounts {
		ids = append(ids, account.ID())
		accountsByID[account.ID()] = account.Clone()
	}

	sort.Ints(ids)

	list := []*bubble.Account{}
	start := (page - 1) * per
	end := (page) * per
	for i := start; i < end; i++ {
		if i >= len(ids) {
			break
		}
		if account, ok := accountsByID[ids[i]]; ok {
			list = append(list, account)
		}
	}

	return list, nil
}

// Find should return the Account belonging to the given ID
func (ms *MockStore) Find(id int) (*bubble.Account, error) {
	if account, ok := ms.accounts[id]; ok {
		return account, nil
	}

	return nil, fmt.Errorf("no account with ID: %d", id)
}

// Save should insert or update the given Account
func (ms *MockStore) Save(account *bubble.Account) error {
	if account == nil {
		return fmt.Errorf("cannot save nil account")
	}

	if account.ID() <= 0 {
		return fmt.Errorf("cannot save account with invalid ID: %d", account.ID())
	}

	ms.accounts[account.ID()] = account
	return nil
}
