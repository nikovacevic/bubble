package user

import (
	"github.com/nikovacevic/bubble/pkg/bubble"
)

// Store provides interactions with a repository of Users
type Store interface {
	// List should provide a the given page from paginated list of Users
	// TODO: sorting (by email, name, etc.)
	// TODO: query (by role, etc.)
	List(per, page int) ([]*bubble.User, error)

	// Find should return the User belonging to the given ID
	Find(id int) (*bubble.User, error)

	// Save should insert or update the given User
	Save(user *bubble.User) error
}
