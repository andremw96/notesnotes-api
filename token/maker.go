package token

import "time"

// interface for managing token
type Maker interface {
	// create token for specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// check if token valid or not
	VerifyToken(token string) (*Payload, error)
}
