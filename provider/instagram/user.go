package instagram

import (
	"github.com/tamnd/gauth"
	"github.com/tamnd/httpclient"
)

var endpointProfile = "https://api.instagram.com/v1/users/self"

func User(token *gauth.AccessToken) (*gauth.User, error) {
	URL := endpointProfile + "?access_token=" + token.Token
	u := struct {
		Data struct {
			Bio    string `json:"bio"`
			Counts struct {
				FollowedBy int `json:"followed_by"`
				Follows    int `json:"follows"`
				Media      int `json:"media"`
			} `json:"counts"`
			FullName       string `json:"full_name"`
			ID             string `json:"id"`
			ProfilePicture string `json:"profile_picture"`
			Username       string `json:"username"`
			Website        string `json:"website"`
		} `json:"data"`
	}{}
	err := httpclient.JSON(URL, &u)
	if err != nil {
		return nil, err
	}

	data := u.Data
	user := gauth.User{
		ID:          data.ID,
		Name:        data.FullName,
		Username:    data.Username,
		Description: data.Bio,
		Avatar:      data.ProfilePicture,
		Raw:         u,
	}

	return &user, nil
}
