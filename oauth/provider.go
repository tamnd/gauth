package oauth

import (
	"net/url"
	"tamnd/misc/gauth"

	"github.com/mrjones/oauth"
)

type Endpoint struct {
	AuthURL        string
	RequestURL     string
	AccessTokenURL string
}

type Provider struct {
	Endpoint Endpoint
	callbackURL string
	consumer *oauth.Consumer
}

func (p *Provider) Init(clientId string, secret string, callbackURL string) error {
	p.consumer := oauth.NewConsumer(
		clientId,
		secret,
		oauth.ServiceProvider{
			AuthorizeTokenUrl: p.Endpoint.AuthURL,
			RequestTokenUrl:   p.Endpoint.RequestURL,
			AccessTokenUrl:    p.Endpoint.AccessTokenURL,
		})
	return nil
}


func (p *Provider) AuthURL(state string) (loginURL string, requestToken string, err error) {
	requestToken, loginURL, err = p.consumer.GetRequestTokenAndUrl(p.callbackURL)
	return
}

// Authorize get an access token from authorization code.
func (p *Provider) Authorize(code string, params ...string) (*gauth.AccessToken, error) {
	requestToken := params[0]
	atoken, err := p.consumer.AuthorizeToken(requestToken, code)
	if err != nil {
		return nil, err
	}
}
