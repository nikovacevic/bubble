package bubble

import "time"

// AuthenticationToken is a random byte value used to authenticate a request
// for an IdentityToken. It is a single-use password, essentially.
type AuthenticationToken struct {
	Value  []byte
	Expiry time.Time
}

// IdentityToken is a longer-lasting, multiple-use token for authenticating
// application requests by a User who has already successfullly authenticated.
type IdentityToken struct {
	Value  []byte
	Expiry time.Time
}
