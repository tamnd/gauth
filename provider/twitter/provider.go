package twitter

import (
	"tamnd/misc/gauth"
	"tamnd/misc/gauth/oauth"
)

var provider = &oauth.Provider{
	Endpoint: oauth.Endpoint{
		AuthURL:        "https://api.twitter.com/oauth/authorize",
		RequestURL:     "https://api.twitter.com/oauth/request_token",
		AccessTokenURL: "https://api.twitter.com/oauth/access_token",
	},
	Code: "oauth_verifier",
	User: getUser,
}

func getUser(token *oauth.AccessToken) (*gauth.User, error) {

}

func init() {
	gauth.Register("twitter", provider)
}
