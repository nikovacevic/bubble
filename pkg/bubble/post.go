package bubble

import "github.com/google/uuid"

// Post describes a post
// TODO keep as IDs, or do Auth(), Venue(), etc.
type Post interface {
	ID() uuid.UUID
	AuthID() uuid.UUID
	VenueID() uuid.UUID
	IsPublic() bool
	String() string
}
