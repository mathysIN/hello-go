package main

import (
	"fmt"
	"net/http"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

var state string

func main() {
	redirectURL := "https://localhost:3000"
	auth := spotifyauth.New(spotifyauth.WithRedirectURL(redirectURL), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))

	url := auth.AuthURL(state)
	fmt.Println(url)
}

func redirectHandler(w http.ResponseWriter, r *http.Request, auth *spotifyauth.Authenticator) {
	token, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusNotFound)
		return
	}
	client := spotify.New(auth.Client(r.Context(), token))
	user, err := client.CurrentUser(r.Context())
	if err != nil {
		http.Error(w, "Couldn't get user", http.StatusNotFound)
		return
	}
	fmt.Println(user.DisplayName)
}
