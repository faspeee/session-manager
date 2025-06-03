package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
)

var oauth2Config = oauth2.Config{
	RedirectURL: "http://localhost:8080/callback", // Callback URL
	Scopes:      []string{"read", "write"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "http://localhost:8081/authorize", // Authorization server URL
		TokenURL: "http://localhost:8081/token",     // Token server URL
	},
}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)

	log.Println("OAuth2 Client running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Main page with register link
func handleMain(w http.ResponseWriter, r *http.Request) {
	// Register client dynamically
	clientID, clientSecret := registerClient()

	// Save these credentials in the oauth2Config
	oauth2Config.ClientID = clientID
	oauth2Config.ClientSecret = clientSecret

	// Display the login link
	fmt.Fprintf(w, `<html><body><a href="/login">Login with Custom OAuth2</a></body></html>`)
}

// Register a client dynamically by calling the registration endpoint
func registerClient() (string, string) {
	// Make a request to the OAuth2 provider's /register endpoint to get Client ID and Secret
	resp, err := http.Get("http://localhost:8081/register")
	if err != nil {
		log.Fatal("Failed to register client:", err)
	}
	defer resp.Body.Close()

	// Read the response and extract Client ID and Client Secret
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response:", err)
	}

	// For simplicity, assume the response contains the Client ID and Secret in plain text
	clientID := string(body[:16])     // Assuming first 16 characters is Client ID
	clientSecret := string(body[17:]) // Assuming the rest is Client Secret

	return clientID, clientSecret
}

// Redirect to OAuth2 provider for login
func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Redirect to the OAuth2 authorization server
	url := oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

// Handle callback from OAuth2 provider
func handleCallback(w http.ResponseWriter, r *http.Request) {
	// Get the authorization code from the URL
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	// Exchange the authorization code for an access token
	token, err := oauth2Config.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to exchange code for token: %v", err), http.StatusInternalServerError)
		return
	}

	// Display the access token
	fmt.Fprintf(w, "Access Token: %s", token.AccessToken)
}
