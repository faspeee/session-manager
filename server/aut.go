package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Store for registered clients and authorization codes
var clients = make(map[string]string) // client_id -> client_secret
var authCode = ""                     // authorization code (in a real app, you'd store this securely)

// Generate a random string for Client ID and Client Secret
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < length; i++ {
		result += string(charset[rand.Intn(len(charset))])
	}
	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Routes for the OAuth2 provider
	http.HandleFunc("/authorize", handleAuthorize) // Handle authorization requests
	http.HandleFunc("/token", handleToken)         // Handle token requests
	http.HandleFunc("/register", handleRegister)   // Handle client registration requests
	http.HandleFunc("/", handleMain)               // Main page (a simple registration link)

	log.Println("OAuth2 Authorization Server running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// Handle main page (simple link to register new client)
func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html><body><a href="/register">Register a new client</a></body></html>`)
}

// Handle client registration (returns Client ID and Client Secret)
func handleRegister(w http.ResponseWriter, r *http.Request) {
	// Generate a random Client ID and Client Secret
	clientID := generateRandomString(16)
	clientSecret := generateRandomString(32)

	// Store these in the clients map (simulating a database)
	clients[clientID] = clientSecret

	// Respond with the generated credentials
	fmt.Fprintf(w, "Client registered successfully! Use the following credentials:\n")
	fmt.Fprintf(w, "Client ID: %s\n", clientID)
	fmt.Fprintf(w, "Client Secret: %s\n", clientSecret)
}

// Handle authorization request (display login form or authenticate)
func handleAuthorize(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	redirectURI := r.URL.Query().Get("redirect_uri")
	responseType := r.URL.Query().Get("response_type")
	state := r.URL.Query().Get("state")

	// Check that the necessary parameters are provided
	if clientID == "" || redirectURI == "" || responseType != "code" {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	// Validate the client ID (check if it's in the clients map)
	if _, ok := clients[clientID]; !ok {
		http.Error(w, "Invalid client", http.StatusUnauthorized)
		return
	}

	// If method is GET, show login form
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, `<html><body>
		<form method="POST" action="/authorize">
			Username: <input type="text" name="username" /><br />
			Password: <input type="password" name="password" /><br />
			<input type="submit" value="Login" />
		</form></body></html>`)
		return
	}

	// Handle login form submission and validate user credentials
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username != "user" || password != "pass" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Simulate issuing an authorization code
	authCode = "auth_code_123"

	// Redirect back to the client with the authorization code and state
	http.Redirect(w, r, fmt.Sprintf("%s?code=%s&state=%s", redirectURI, authCode, state), http.StatusFound)
}

// Handle token exchange request
func handleToken(w http.ResponseWriter, r *http.Request) {
	// Parse the form values
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")
	code := r.FormValue("code")
	redirectURI := r.FormValue("redirect_uri")
	grantType := r.FormValue("grant_type")

	// Validate client credentials
	if expectedSecret, ok := clients[clientID]; !ok || expectedSecret != clientSecret {
		http.Error(w, "Invalid client credentials", http.StatusUnauthorized)
		return
	}

	// Ensure the grant type is 'authorization_code'
	if grantType != "authorization_code" {
		http.Error(w, "Invalid grant type", http.StatusBadRequest)
		return
	}

	// Validate the authorization code
	if code != authCode {
		http.Error(w, "Invalid authorization code", http.StatusBadRequest)
		return
	}

	// Validate the redirect URI
	if redirectURI != r.URL.Query().Get("redirect_uri") {
		http.Error(w, "Invalid redirect URI", http.StatusBadRequest)
		return
	}

	// Issue the access token
	token := map[string]interface{}{
		"access_token":  "access_token_123",
		"token_type":    "bearer",
		"expires_in":    3600,
		"scope":         "read write",
		"refresh_token": "refresh_token_123",
	}

	tokenJSON, _ := json.Marshal(token)

	// Send the token as JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(tokenJSON)
}
