package venue

import (
	"github.com/google/uuid"
	"github.com/nikovacevic/bubble/pkg/bubble"
)

// Store provides interactions with a repository of Venues
type Store interface {
	// Find should return the Venue belonging to the given ID
	Find(id uuid.UUID) (*bubble.Venue, error)

	// Save should insert or update the given Venue
	Save(venue *bubble.Venue) error
}
