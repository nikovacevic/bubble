package destination

import (
	"github.com/nikovacevic/bubble/pkg/bubble"
)

// Store provides interactions with a repository of Destinations
type Store interface {
	// FindDestination should return the Destination belonging to the given ID
	FindDestination(id int) (*bubble.Destination, error)

	// SaveDestination should insert or update the given Destination
	SaveDestination(venue *bubble.Destination) error
}
