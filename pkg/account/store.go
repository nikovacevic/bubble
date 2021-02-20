package account

import (
	"github.com/nikovacevic/bubble/pkg/bubble"
)

// Store provides interactions with a repository of Accounts
type Store interface {
	// List should provide a the given page from paginated list of Accounts
	// TODO: sorting (by email, name, etc.)
	// TODO: query (by role, etc.)
	List(per, page int) ([]*bubble.Account, error)

	// Find should return the Account belonging to the given ID
	Find(id int) (*bubble.Account, error)

	// Save should insert or update the given Account
	Save(account *bubble.Account) error
}
