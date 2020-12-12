package bubble

import "github.com/google/uuid"

// VenueType describes the type of venue
type VenueType int

const (
	// PageVenue corresponds to a page venue
	PageVenue VenueType = iota

	// GroupVenue corresponds to a group chat venue
	GroupVenue

	// EventVenue corresponds to an event venue
	EventVenue

	// PollVenue corresponds to a poll venue
	PollVenue
)

// Venue is a repository for Messages
type Venue interface {
	ID() uuid.UUID
	Name() string
	IsPublic() bool
}

// Page is a Venue that acts as a document, receiving messages of type Post.
// It can be public or private. If no Users except the Owner have access, it
// can be thought of as a static resource.
type Page struct {
	id       uuid.UUID
	Name     string
	Content  []byte
	isPublic bool
}

// TODO:
// Group: private repository for direct messages among select Users
// Event: public or private calendar event, repository for Invitations and RSVPs (also a Message)
// Poll: public or private respository for votes (also a Message)
