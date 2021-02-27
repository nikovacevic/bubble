package bubble

// DestinationType describes the type of venue
type DestinationType int

const (
	// PageDestination corresponds to a page venue
	PageDestination DestinationType = iota

	// GroupDestination corresponds to a group chat venue
	GroupDestination

	// EventDestination corresponds to an event venue
	EventDestination

	// PollDestination corresponds to a poll venue
	PollDestination
)

// Destination is a repository for Messages
type Destination interface {
	ID() int
	Name() string
	IsPublic() bool
}

// Page is a Destination that acts as a document, receiving messages of type Post.
// It can be public or private. If no Accounts except the Owner have access, it
// can be thought of as a static resource.
type Page struct {
	id       int
	Name     string
	Content  []byte
	isPublic bool
}

// TODO:
// Group: private repository for direct messages among select Accounts
// Event: public or private calendar event, repository for Invitations and RSVPs (also a Message)
// Poll: public or private respository for votes (also a Message)
