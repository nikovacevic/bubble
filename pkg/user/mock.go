package user

import (
	"fmt"
	"sort"

	"github.com/nikovacevic/bubble/pkg/bubble"
)

// MockStore implements the Store interface for testing
type MockStore struct {
	users map[int]*bubble.User
}

// NewMockStore creates a new MockStore with 10 hard-coded users
func NewMockStore() *MockStore {
	users := []*bubble.User{}
	users = append(users, bubble.NewUser("adam@getbubble.org", "Adam"))
	users = append(users, bubble.NewUser("beth@getbubble.org", "Beth"))
	users = append(users, bubble.NewUser("chip@getbubble.org", "Chip"))
	users = append(users, bubble.NewUser("dave@getbubble.org", "Dave"))
	users = append(users, bubble.NewUser("elen@getbubble.org", "Elen"))
	users = append(users, bubble.NewUser("faye@getbubble.org", "Faye"))
	users = append(users, bubble.NewUser("gary@getbubble.org", "Gary"))
	users = append(users, bubble.NewUser("hugh@getbubble.org", "Hugh"))
	users = append(users, bubble.NewUser("inga@getbubble.org", "Inga"))
	users = append(users, bubble.NewUser("jake@getbubble.org", "Jake"))

	userMap := map[int]*bubble.User{}
	for _, user := range users {
		userMap[user.ID()] = user
	}

	return &MockStore{users: userMap}
}

// List should provide a the given page from paginated list of Users
func (ms *MockStore) List(per, page int) ([]*bubble.User, error) {
	usersByID := map[int]*bubble.User{}
	ids := []int{}

	for _, user := range ms.users {
		ids = append(ids, user.ID())
		usersByID[user.ID()] = user.Clone()
	}

	sort.Ints(ids)

	list := []*bubble.User{}
	start := (page - 1) * per
	end := (page) * per
	for i := start; i < end; i++ {
		if i >= len(ids) {
			break
		}
		if user, ok := usersByID[ids[i]]; ok {
			list = append(list, user)
		}
	}

	return list, nil
}

// Find should return the User belonging to the given ID
func (ms *MockStore) Find(id int) (*bubble.User, error) {
	if user, ok := ms.users[id]; ok {
		return user, nil
	}

	return nil, fmt.Errorf("no user with ID: %d", id)
}

// Save should insert or update the given User
func (ms *MockStore) Save(user *bubble.User) error {
	if user == nil {
		return fmt.Errorf("cannot save nil user")
	}

	if user.ID() <= 0 {
		return fmt.Errorf("cannot save user with invalid ID: %d", user.ID())
	}

	ms.users[user.ID()] = user
	return nil
}
