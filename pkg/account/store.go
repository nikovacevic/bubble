package account

import (
	"github.com/nikovacevic/bubble/pkg/bubble"
)

// Store provides interactions with a repository of Accounts
type Store interface {
	// ListAccount should return a list of Accounts
	ListAccounts() ([]*bubble.Account, error)

	// FindAccount should return the Account belonging to the given ID
	FindAccount(id int) (*bubble.Account, error)

	// SaveAccount should insert or update the given Account
	SaveAccount(account *bubble.Account) error
}
