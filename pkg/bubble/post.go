package bubble

import "github.com/google/uuid"

// Post describes a post
type Post interface {
	AuthID() uuid.UUID
}
