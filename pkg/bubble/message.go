package bubble

// Message describes an object that can be sent to a Destination by a Account.
type Message interface {
	// ID is the unique identifier of the message, itself.
	ID() int

	// AccountID is the unique indentifier of the sender.
	AccountID() int

	// DestinationID is the unique identifier of the venue to which the message
	// has been sent.
	// TODO: should this be plural, DestinationIDs?
	DestinationID() int

	// IsPublic, if true, means the message can be seen by anyone, even those
	// who are not authenticated.
	IsPublic() bool

	// String is a basic text representation of the message.
	String() string
}

// Direct implements Message for DMs. Currently just text, but could include
// images, etc. eventually.
type Direct struct {
	id        int
	accountID int
	venueID   int
	isPublic  bool
	Text      string
}

// ID is the unique identifier of the message, itself.
func (d *Direct) ID() int {
	return d.id
}

// AccountID is the unique indentifier of the sender.
func (d *Direct) AccountID() int {
	return d.accountID
}

// DestinationID is the unique identifier of the venue to which the message
// has been sent.
func (d *Direct) DestinationID() int {
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
// Event: public or private calendar event (also a Destination)
// Invitation: public or private invitation to join an event
// RSVP: public or private message belonging to an event
// Poll: public or private message with vote options (also a Destination)
// Vote: public or private message belonging to a poll
