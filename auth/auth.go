package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-yee/spot/configs"
	api "github.com/zmb3/spotify/v2"
	apiauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

// NewAPIClient creates an api client handler with methods to access spotify
//
// The oauth2 request implementation is based off of the example from the
// api library. See:
// https://github.com/zmb3/spotify/blob/master/examples/authenticate/pkce/pkce.go
func NewAPIClient(config configs.Config) *api.Client {
	ch := make(chan *api.Client)
	authenticator := apiauth.New(
		apiauth.WithClientID(config.ClientID),
		apiauth.WithClientSecret(config.ClientSecret),
		apiauth.WithRedirectURL(config.RedirectURL),
		apiauth.WithScopes(
			apiauth.ScopeUserReadPrivate,
			apiauth.ScopeUserReadPlaybackState,
		))

	// Randomly generate data for oauth2 request
	state := oauth2.GenerateVerifier()
	codeVerifier := oauth2.GenerateVerifier()
	codeChallenge := oauth2.S256ChallengeFromVerifier(codeVerifier)

	authRequest := func(w http.ResponseWriter, r *http.Request) {
		token, err := authenticator.Token(
			r.Context(),
			state,
			r,
			oauth2.SetAuthURLParam("code_verifier", codeVerifier))
		if err != nil {
			http.Error(w, "Could not get token", http.StatusForbidden)
			log.Fatal(err)
		}
		if formState := r.FormValue("state"); formState != state {
			http.NotFound(w, r)
			log.Fatalf("State mismatch: %s != %s\n", formState, state)
		}
		// use the token to get an authenticated client
		client := api.New(authenticator.Client(r.Context(), token))
		fmt.Fprintf(w, "Login Completed!")
		ch <- client
	}

	// Start http server to handle redirect auth request
	http.HandleFunc("/callback", authRequest)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})

	go http.ListenAndServe(config.RedirectPort(), nil)

	url := authenticator.AuthURL(
		state,
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
	)
	fmt.Printf(
		"Use the following link in your browser to authorize access: \n\n%s\n",
		url)

	// wait for auth to complete
	client := <-ch

	return client
}
