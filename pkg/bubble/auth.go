package bubble

import "time"

// Token is a random byte value used to authenticate requests.
type Token struct {
	Value  []byte
	Expiry time.Time
}

// TODO
type Session struct {
	AccountID int
	Token     Token
	Expires   time.Time
}
