package bubble

import "github.com/google/uuid"

type Message struct {
	id       uuid.UUID
	authID   uuid.UUID
	venueID  uuid.UUID
	isPublic bool
	Text     string
}

func (m *Message) ID() uuid.UUID {
	return m.id
}

func (m *Message) AuthID() uuid.UUID {
	return m.authID
}

func (m *Message) VenueID() uuid.UUID {
	return m.venueID
}

func (m *Message) IsPublic() bool {
	return m.isPublic
}

func (m *Message) String() string {
	return m.Text
}
