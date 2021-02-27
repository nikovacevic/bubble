package message

import (
	"github.com/nikovacevic/bubble/pkg/bubble"
)

// Store provides interactions with a repository of Messages
type Store interface {
	// FindMessage should return the Message belonging to the given ID
	FindMessage(id int) (*bubble.Message, error)

	// SaveMessage should insert or update the given Message
	SaveMessage(Message *bubble.Message) error
}
