package bubble

import "github.com/google/uuid"

type Auth interface {
	ID() uuid.UUID
	VenueRoles() VenueRoles
}

type Person struct {
	id    uuid.UUID
	vrs   VenueRoles
	Email string
	Name  string
}

func (p Person) ID() uuid.UUID {
	return p.id
}

func (p Person) VenueRoles() VenueRoles {
	return p.vrs
}

func (p Person) NewMessage(text string, venueID uuid.UUID, isPublic bool) *Message {
	return &Message{
		id:      uuid.New(),
		authID:  p.ID(),
		venueID: venueID,
	}
}
