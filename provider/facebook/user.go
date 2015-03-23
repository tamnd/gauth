package facebook

import (
	"net/url"

	"github.com/tamnd/gauth"
	"github.com/tamnd/httpclient"
)

var endpointProfile = "https://graph.facebook.com/me?fields=email,first_name,last_name,link,bio,id,name,picture,location"

func User(token *gauth.AccessToken) (*gauth.User, error) {
	URL := endpointProfile + "&access_token=" + url.QueryEscape(token.Token)

	u := struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Bio     string `json:"bio"`
		Name    string `json:"name"`
		Link    string `json:"link"`
		Picture struct {
			Data struct {
				URL string `json:"url"`
			} `json:"data"`
		} `json:"picture"`
		Location struct {
			Name string `json:"name"`
		} `json:"location"`
	}{}

	err := httpclient.JSON(URL, &u)
	if err != nil {
		return nil, err
	}

	user := gauth.User{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		Username:    u.Name,
		Location:    u.Location.Name,
		Description: u.Bio,
		Avatar:      u.Picture.Data.URL,
		Raw:         u,
	}

	return &user, nil
}
