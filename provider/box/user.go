package box

import (
	"fmt"

	"github.com/tamnd/gauth"
	"github.com/tamnd/httpclient"
)

var endpointProfile = "https://api.box.com/2.0/users/me"

func User(token *gauth.AccessToken) (*gauth.User, error) {
	URL := endpointProfile + "&oauth_token=" + token.Token
	fmt.Println(URL)
	content, err := httpclient.String(URL)
	if err != nil {
		return nil, err
	}

	user := gauth.User{
		Raw: content,
	}

	return &user, nil
}
