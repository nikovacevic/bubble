package bubble

import "github.com/google/uuid"

// User is an identity that can be authenticated and can interact with
// resources (Venues, Messages, etc.)
type User struct {
	ID          uuid.UUID
	Permissions Permissions
	Email       string
	Name        string
	idToken     *IdentityToken
	authToken   *AuthenticationToken
}
