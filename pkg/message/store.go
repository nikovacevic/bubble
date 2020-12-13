package message

import (
	"github.com/google/uuid"
	"github.com/nikovacevic/bubble/pkg/bubble"
)

// Store provides interactions with a repository of Messages
type Store interface {
	// Find should return the Message belonging to the given ID
	Find(id uuid.UUID) (*bubble.Message, error)

	// Save should insert or update the given Message
	Save(Message *bubble.Message) error
}