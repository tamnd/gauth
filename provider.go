package gauth

import "time"

type Provider interface {
	// AuthURL returns the login URL.
	AuthURL(state string) (requestToken string, loginURL string, err error)

	// Authorize gets an access token from verification code.
	Authorize(code string, params ...string) (*AccessToken, error)

	// User returns the authenticating user from the given access token.
	User(token *AccessToken) (*User, error)
}

type AccessToken struct {
	// The Token
	Token     string
	Secret    string
	Expires   bool
	ExpiresAt time.Time
}
