package bubble

import (
	"fmt"
)

// Account is an identity that can be authenticated and can interact with
// resources (Venues, Messages, etc.)
type Account struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	Role      Role   `json:"role"`
	ID        int    `json:"-"`
	idToken   *Token `json:"-"`
	authToken *Token `json:"-"`
}

// NewAccount instantiates a new Account with a new UUID and no auth/id tokens
func NewAccount(email, name string, role Role) *Account {
	return &Account{
		Email: email,
		Name:  name,
		Role:  role,
	}
}

// String returns a string representation of the Account
func (a *Account) String() string {
	if a == nil {
		return "<nil>"
	}

	return fmt.Sprintf("%s %s", a.Name, a.Email)
}

// Role describes a set of permissions for interacting with things. It's the
// default level of access for any given Account to interact with any given
// object, but a specific Role for a specific object can be granted, which
// overrides the Account's default Role.
type Role int

const (
	// VisitorRole limits an account to at most read access, and only for
	// public things.
	VisitorRole Role = iota

	// MemberRole gives an account read access to most things, and write access
	// to some things, but no moderation access.
	MemberRole

	// ModeratorRole gives an account read, write, and moderation access to
	// most things, but no administrative access.
	ModeratorRole

	// AdministratorRole allows an account to read, write, and moderate most
	// things, as well as to carry out administrative tasks like creating new
	// destinations, inviting new accounts, and granting roles.
	AdministratorRole

	// OwnerRole should be unique to one account per bubble, and allows any
	// action to be taken. Owner is the only role that an administrator cannot
	// assign or unassign, so an owner must transfer the role if a new owner is
	// to adopt ownership of the bubble.
	OwnerRole
)

func (r Role) String() string {
	switch r {
	case OwnerRole:
		return "Owner"
	case AdministratorRole:
		return "Administrator"
	case ModeratorRole:
		return "Moderator"
	case MemberRole:
		return "Member"
	}
	return "Visitor"
}
