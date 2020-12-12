package bubble

import "github.com/google/uuid"

// Message describes an object that can be sent to a Venue by a User.
type Message interface {
	// ID is the unique identifier of the message, itself.
	ID() uuid.UUID

	// UserID is the unique indentifier of the sender.
	UserID() uuid.UUID

	// VenueID is the unique identifier of the venue to which the message
	// has been sent.
	// TODO: should this be plural, VenueIDs?
	VenueID() uuid.UUID

	// IsPublic, if true, means the message can be seen by anyone, even those
	// who are not authenticated.
	IsPublic() bool

	// String is a basic text representation of the message.
	String() string
}

// Direct implements Message for DMs. Currently just text, but could include
// images, etc. eventually.
type Direct struct {
	id       uuid.UUID
	userID   uuid.UUID
	venueID  uuid.UUID
	isPublic bool
	Text     string
}

// ID is the unique identifier of the message, itself.
func (d *Direct) ID() uuid.UUID {
	return d.id
}

// UserID is the unique indentifier of the sender.
func (d *Direct) UserID() uuid.UUID {
	return d.userID
}

// VenueID is the unique identifier of the venue to which the message
// has been sent.
func (d *Direct) VenueID() uuid.UUID {
	return d.venueID
}

// IsPublic determines whether the message can be seen by anyone, even those
// who are not authenticated.
func (d *Direct) IsPublic() bool {
	return d.isPublic
}

// String is a basic text representation of the message.
func (d *Direct) String() string {
	return d.Text
}

// TODO:
// Post: public or private message belonging to a page or event
// Event: public or private calendar event (also a Venue)
// Invitation: public or private invitation to join an event
// RSVP: public or private message belonging to an event
// Poll: public or private message with vote options (also a Venue)
// Vote: public or private message belonging to a poll
