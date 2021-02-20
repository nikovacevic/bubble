package bubble

// Message describes an object that can be sent to a Venue by a User.
type Message interface {
	// ID is the unique identifier of the message, itself.
	ID() int

	// UserID is the unique indentifier of the sender.
	UserID() int

	// VenueID is the unique identifier of the venue to which the message
	// has been sent.
	// TODO: should this be plural, VenueIDs?
	VenueID() int

	// IsPublic, if true, means the message can be seen by anyone, even those
	// who are not authenticated.
	IsPublic() bool

	// String is a basic text representation of the message.
	String() string
}

// Direct implements Message for DMs. Currently just text, but could include
// images, etc. eventually.
type Direct struct {
	id       int
	userID   int
	venueID  int
	isPublic bool
	Text     string
}

// ID is the unique identifier of the message, itself.
func (d *Direct) ID() int {
	return d.id
}

// UserID is the unique indentifier of the sender.
func (d *Direct) UserID() int {
	return d.userID
}

// VenueID is the unique identifier of the venue to which the message
// has been sent.
func (d *Direct) VenueID() int {
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
