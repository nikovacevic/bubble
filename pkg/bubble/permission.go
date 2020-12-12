package bubble

import "github.com/google/uuid"

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
