package oauth2

import (
	"github.com/tamnd/gauth"
	"golang.org/x/oauth2"
)

type Endpoint struct {
	AuthURL  string
	TokenURL string
}

type Provider struct {
	Endpoint Endpoint
	UserFn   func(token *gauth.AccessToken) (*gauth.User, error)
	config   *oauth2.Config
}

func (p *Provider) Init(clientId string, secret string, callbackURL string, scopes ...string) error {
	p.config = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: secret,
		Scopes:       scopes,
		RedirectURL:  callbackURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  p.Endpoint.AuthURL,  // "https://provider.com/o/oauth2/auth"
			TokenURL: p.Endpoint.TokenURL, // "https://provider.com/o/oauth2/token"
		},
	}
	return nil
}

func (p *Provider) AuthURL(state string) (string, string, error) {
	return p.config.AuthCodeURL(state), "", nil
}

func (p *Provider) Authorize(code string, params ...string) (*gauth.AccessToken, error) {
	atoken, err := p.config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}
	return &gauth.AccessToken{
		Token:     atoken.AccessToken,
		Secret:    "",
		Expires:   true,
		ExpiresAt: atoken.Expiry,
	}, nil
}

func (p *Provider) User(token *gauth.AccessToken) (*gauth.User, error) {
	return p.UserFn(token)
}
