package meetup

import (
	"github.com/tamnd/gouth"
	"github.com/tamnd/gouth/oauth2"
	"github.com/tamnd/httpclient"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://secure.meetup.com/oauth2/authorize",
	TokenURL: "https://secure.meetup.com/oauth2/access",
}

func New(clientId string, secret string, callbackURL string, scopes ...string) gouth.Provider {
	provider := &oauth2.Provider{
		Endpoint: Endpoint,
		Getuser:  getUser,
	}
	provider.Init(clientId, secret, callbackURL)
	return provider
}

var endpointProfile = "https://api.meetup.com/2/member/self"

func getUser(token *gouth.AccessToken) (*gouth.User, error) {
	URL := endpointProfile + "?access_token=" + token.Token

	content, err := httpclient.String(URL)
	if err != nil {
		return nil, err
	}

	user := gouth.User{
		Raw: content,
	}

	return &user, nil
}
