package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/tamnd/gauth"
	"github.com/tamnd/gauth/provider/facebook"
)

var provider gauth.Provider

func init() {
	provider = facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), "http://localhost:8080/auth/facebook/callback")
}

func main() {
	fmt.Println("Start listening on :8080")

	http.Handle("/auth/facebook", http.HandlerFunc(login))
	http.Handle("/auth/facebook/callback", http.HandlerFunc(callback))
	http.ListenAndServe(":8080", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	authURL, _, err := provider.AuthURL("state")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

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
