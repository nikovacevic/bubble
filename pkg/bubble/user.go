package bubble

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// User is an identity that can be authenticated and can interact with
// resources (Venues, Messages, etc.)
type User struct {
	Email       string               `json:"email"`
	Name        string               `json:"name"`
	Permissions Permissions          `json:"permissions"`
	id          uuid.UUID            `json:"-"`
	idToken     *IdentityToken       `json:"-"`
	authToken   *AuthenticationToken `json:"-"`
}

// NewUser instantiates a new user with a new UUID and no auth/id tokens
func NewUser(email, name string) *User {
	return &User{
		id:          uuid.New(),
		Permissions: Permissions{},
		Email:       email,
		Name:        name,
	}
}

// Clone returns a clone of the User
func (u *User) Clone() *User {
	if u == nil {
		return nil
	}

	id := &uuid.UUID{}
	idText, err := u.id.MarshalText()
	if err != nil {
		newID := uuid.New()
		id = &newID
	}
	err = id.UnmarshalText(idText)

	return &User{
		id:          *id,
		Permissions: u.Permissions,
		Email:       u.Email,
		Name:        u.Name,
	}
}

// ID returns a User's ID
func (u *User) ID() uuid.UUID {
	if u == nil {
		return uuid.Nil
	}

	return u.id
}

// String returns a string representation of the user
func (u *User) String() string {
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
// application requests by a User who has already successfullly authenticated.
type IdentityToken struct {
	Value  []byte
	Expiry time.Time
}

// Role describes a set of permissions for interacting with a Venue.
type Role int

const (
	// NoRole does not allow a User to interact with a Venue.
	NoRole Role = iota

	// ReaderRole allows a User read-only access to a Venue.
	ReaderRole

	// WriterRole allows a User to send messages to a Venue, but not edit it,
	// invite Users to it, or delete it.
	WriterRole

	// EditorRole allows a User to edit and send messages to a Venue, as well
	// as to invite Users. Editors cannot delete a Venue.
	EditorRole

	// OwnerRole should be unique to one User. It is given to the creator of a
	// Venue, but can be transferred. Confers access to all features.
	OwnerRole
)

// Permission describes a tuple of a Role and a Venue, allowing the described
// access to the identified Venue. Permissions should belong to Users.
type Permission struct {
	VenueID uuid.UUID
	Role    Role
}

// Permissions are a set of Permission instances, keyed by VenueID
type Permissions map[uuid.UUID]Permission
