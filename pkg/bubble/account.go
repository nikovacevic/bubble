package bubble

import (
	"fmt"
	"time"
)

// Account is an identity that can be authenticated and can interact with
// resources (Venues, Messages, etc.)
type Account struct {
	Email       string               `json:"email"`
	Name        string               `json:"name"`
	Permissions Permissions          `json:"permissions"`
	id          int                  `json:"-"`
	idToken     *IdentityToken       `json:"-"`
	authToken   *AuthenticationToken `json:"-"`
}

// NewAccount instantiates a new Account with a new UUID and no auth/id tokens
func NewAccount(email, name string) *Account {
	return &Account{
		Permissions: Permissions{},
		Email:       email,
		Name:        name,
	}
}

// Clone returns a clone of the Account
func (u *Account) Clone() *Account {
	if u == nil {
		return nil
	}

	return &Account{
		id:          u.id,
		Permissions: u.Permissions,
		Email:       u.Email,
		Name:        u.Name,
	}
}

// ID returns a Account's ID
func (u *Account) ID() int {
	if u == nil {
		return 0
	}

	return u.id
}

// String returns a string representation of the Account
func (u *Account) String() string {
	if u == nil {
		return "<nil>"
	}

	return fmt.Sprintf("%s %s", u.Name, u.Email)
}

// AuthenticationToken is a random byte value used to authenticate a request
// for an IdentityToken. It is a single-use password, essentially.
type AuthenticationToken struct {
	Value  []byte
	Expiry time.Time
}

// IdentityToken is a longer-lasting, multiple-use token for authenticating
// application requests by a Account who has already successfullly authenticated.
type IdentityToken struct {
	Value  []byte
	Expiry time.Time
}

// Role describes a set of permissions for interacting with a Venue.
type Role int

const (
	// NoRole does not allow a Account to interact with a Venue.
	NoRole Role = iota

	// ReaderRole allows a Account read-only access to a Venue.
	ReaderRole

	// WriterRole allows a Account to send messages to a Venue, but not edit it,
	// invite Accounts to it, or delete it.
	WriterRole

	// EditorRole allows a Account to edit and send messages to a Venue, as well
	// as to invite Accounts. Editors cannot delete a Venue.
	EditorRole

	// OwnerRole should be unique to one Account. It is given to the creator of a
	// Venue, but can be transferred. Confers access to all features.
	OwnerRole
)

// Permission describes a tuple of a Role and a Venue, allowing the described
// access to the identified Venue. Permissions should belong to Accounts.
type Permission struct {
	VenueID int
	Role    Role
}

// Permissions are a set of Permission instances, keyed by VenueID
type Permissions map[int]Permission
