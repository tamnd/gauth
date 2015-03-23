# GAuth

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Build Status](https://img.shields.io/travis/tamnd/gauth/master.svg?style=flat-square)](https://travis-ci.org/tamnd/gauth)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/tamnd/gauth)

Clean way to write authentication for [Go](http://www.golang.org).


This package was inspired by https://github.com/intridea/omniauth (written in Ruby) and http://github.com/markbates/goth (written in Go). 


## Features

- Support both OAuth 1.0a and OAuth2
- Have a long list of providers and is easy to write new ones.


## Install 

```
$ go get github.com/tamnd/gauth
```

## List of Providers
Here are the list of supported providers. Would you love to see more providers? Feel free to contribute ones by create your own repository or create pull requests!

**Please keep the list in alphabetical order**

| Provider | Description | Author | Is Official?
| --- | ---- | --- | ---
| [Dropbox](http://github.com/tamnd/gauth/tree/master/provider/dropbox) | Authenticate to [Dropbox](http://www.dropbox.com) using OAuth2. | tamnd | No
| [Facebook](http://github.com/tamnd/gauth/tree/master/provider/facebook) | Authenticate to [Facebook](http://www.facebook.com) using OAuth2. | tamnd | No
| [Foursquare](http://github.com/tamnd/gauth/tree/master/provider/foursquare) | Authenticate to [Foursquare](http://ww.foursquare.com) using OAuth2. | tamnd | No
| [Instagram](http://github.com/tamnd/gauth/tree/master/provider/instagram) | Authenticate to [Instagram](http://www.instagram.com) using OAuth2. | tamnd | No
| [Meetup](http://github.com/tamnd/gauth/tree/master/provider/meetup) | Authenticate to [Meetup](http://www.meetup.com) using OAuth2. | tamnd | No

## Usage

First, import GAuth and create OAuth provider

```go
import (
	"github.com/tamnd/gauth"
	"github.com/tamnd/gauth/provider/facebook"
)

// Create facebook provider
var provider = facebook.New("FACEBOOK_KEY", "FACEBOOK_CLIENT_SECRET", "http://localhost:8080/auth/facebook/callback")
```

Next, define a handler for `GET /auth/facebook`. This will redirect the user to the facebook login page.

```go
func login(w http.ResponseWriter, r *http.Request) {
	authURL, _, err := provider.AuthURL("state")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}
```

Then define the handler for callback request at `/auth/facebook/callback`:

```go
func callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	// Get access token from verification code.
	token, err := provider.Authorize(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the authenticating user
	user, err := provider.User(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal user info and write to the response.
	output, _ := json.Marshal(user)
	w.Write([]byte(output))
}
```

Finally, we register those handlers and start the server: 

```go
func main() {
	fmt.Println("Start listening on :8080")

	http.Handle("/auth/facebook", http.HandlerFunc(login))
	http.Handle("/auth/facebook/callback", http.HandlerFunc(callback))
	http.ListenAndServe(":8080", nil)
}
```

See the full source code at [examples/facebook.go](https://github.com/tamnd/gauth/tree/master/examples/facebook.go).


## Write new provider 

Writing new provider is pretty easy. You just need to define OAuth endpoints:

```go
package yoursite

import "github.com/tamnd/gauth/oauth2"

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://yoursite.com/oauth2/auth",
	TokenURL: "https://yoursite.com/oauth2/token",
}
``` 

and define a function `User` to get user info from an access token: 

```go
package yoursite

import "github.com/tamnd/gauth"

func User(token *gauth.AccessToken) (*gauth.User, error) {
	// Get the user info here...
}
```

For example, here is the Facebook Authentication Provider:

[provider/facebook/facebook.go](http://www.github.com/tamnd/gauth/tree/master/provider/facebook/facebook.go)
```go
package facebook

import (
	"github.com/tamnd/gauth"
	"github.com/tamnd/gauth/oauth2"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://www.facebook.com/dialog/oauth",
	TokenURL: "https://graph.facebook.com/oauth/access_token",
}

func New(clientId string, secret string, callbackURL string, scopes ...string) gauth.Provider {
	provider := &oauth2.Provider{
		Endpoint: Endpoint,
		UserFn:   User,
	}
	provider.Init(clientId, secret, callbackURL)
	return provider
}
```

[provider/facebook/user.go](http://www.github.com/tamnd/gauth/tree/master/provider/facebook/user.go)
```go
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
```


## Contribute

- Fork repository
- Create a feature branch
- Open a new pull request
- Create an issue for bug report or feature request

## Contact

- Nguyen Duc Tam
- [tamnd87@gmail.com](mailto:tamnd87@gmail.com)
- [http://twitter.com/tamnd87](http://twitter.com/tamnd87)

## License
The MIT License (MIT). Please see [LICENSE](LICENSE) for more information.

Copyright (c) 2015 Nguyen Duc Tam, tamnd87@gmail.com

