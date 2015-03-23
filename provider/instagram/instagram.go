package instagram

import (
	"github.com/tamnd/gauth"
	"github.com/tamnd/gauth/oauth2"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://api.instagram.com/oauth/authorize",
	TokenURL: "https://api.instagram.com/oauth/access_token",
}

func New(clientId string, secret string, callbackURL string, scopes ...string) gauth.Provider {
	provider := &oauth2.Provider{
		Endpoint: Endpoint,
		UserFn:   User,
	}
	provider.Init(clientId, secret, callbackURL)
	return provider
}
