package dropbox

import (
	"strconv"

	"github.com/tamnd/gauth"
	"github.com/tamnd/httpclient"
)

var endpointProfile = "https://api.dropbox.com/1/account/info"

func User(token *gauth.AccessToken) (*gauth.User, error) {
	URL := endpointProfile + "?access_token=" + token.Token
	u := struct {
		Country       string `json:"country"`
		DisplayName   string `json:"display_name"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		IsPaired      bool   `json:"is_paired"`
		Locale        string `json:"locale"`
		NameDetails   struct {
			FamiliarName string `json:"familiar_name"`
			GivenName    string `json:"given_name"`
			Surname      string `json:"surname"`
		} `json:"name_details"`
		QuotaInfo struct {
			Datastores int `json:"datastores"`
			Normal     int `json:"normal"`
			Quota      int `json:"quota"`
			Shared     int `json:"shared"`
		} `json:"quota_info"`
		ReferralLink string      `json:"referral_link"`
		Team         interface{} `json:"team"`
		Uid          int         `json:"uid"`
	}{}

	err := httpclient.JSON(URL, &u)
	if err != nil {
		return nil, err
	}

	user := gauth.User{
		ID:        strconv.Itoa(u.Uid),
		Name:      u.DisplayName,
		Firstname: u.NameDetails.Surname,
		Lastname:  u.NameDetails.FamiliarName,
		Email:     u.Email,
		Location:  u.Country,
		Raw:       u,
	}

	return &user, nil
}
