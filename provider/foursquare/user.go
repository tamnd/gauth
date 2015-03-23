package foursquare

import (
	"github.com/tamnd/gauth"
	"github.com/tamnd/httpclient"
)

var endpointProfile = "https://api.foursquare.com/v2/users/self?v=20150101"

func User(token *gauth.AccessToken) (*gauth.User, error) {
	URL := endpointProfile + "&oauth_token=" + token.Token

	content, err := httpclient.String(URL)
	if err != nil {
		return nil, err
	}

	user := gauth.User{
		Raw: content,
	}

	return &user, nil
}
