package bubble

import "github.com/google/uuid"

type Venue struct {
	id          uuid.UUID
	Name        string
	Description string
	IsPublic    bool
}

type Role int

const (
	NoRole Role = iota
	ReadRole
	WriteRole
	EditRole
	OwnerRole
)

type VenueRole struct {
	VenueID uuid.UUID
	Role    Role
}

type VenueRoles map[uuid.UUID]VenueRole
