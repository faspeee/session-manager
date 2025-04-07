package service

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"session-manager/model/collection"
	"session-manager/model/request"
	"session-manager/model/response"
	"session-manager/repository"
)

func Login(info request.LoginInfo) (string, error) {
	user, err := repository.GetUserByNameAndPassword(info.Username, info.Password)
	if err != nil {
		return "", err
	}
	jsonData, _ := json.Marshal(user)
	return string(jsonData), nil
}
func CheckToken(token string) bool {
	return true
}

func Registry(user request.UserRequest) (response.UserResponse, error) {
	clientId, clientSecret, err := generateClientCredentials()
	userCollection := createUserCollection(user, clientId, clientSecret)
	if err != nil {
		return response.UserResponse{}, errors.New("i have an error when create client credentials")
	}
	err = repository.RegistryUser(userCollection)
	if err != nil {
		return response.UserResponse{}, errors.New("i have an error when registry new user")
	}
	return response.UserResponse{Username: user.Username, Email: user.Email}, nil
}

// generateRandomBytes creates a slice of securely generated random bytes.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	// crypto/rand.Read returns the number of bytes read and an error
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

// generateSecureString returns a securely generated random string in hexadecimal form.
func generateSecureString(byteLength int) (string, error) {
	bytes, err := generateRandomBytes(byteLength)
	if err != nil {
		return "", err
	}
	// Hex encoding doubles the length of the output string.
	return hex.EncodeToString(bytes), nil
}

// generateClientCredentials generates both client id and client secret.
func generateClientCredentials() (string, string, error) {
	// Use 32 bytes for client ID (64 hex characters)
	clientID, err := generateSecureString(32)
	if err != nil {
		return "", "", err
	}
	// Use 64 bytes for client secret (128 hex characters)
	clientSecret, err := generateSecureString(64)
	if err != nil {
		return "", "", err
	}
	return clientID, clientSecret, nil
}
func createUserCollection(userDto request.UserRequest, clientId string, clientSecret string) collection.User {
	user := collection.User{
		Username:     userDto.Username,
		Password:     userDto.Password,
		Birthday:     userDto.Birthday,
		Country:      collection.Country(userDto.Country),
		Email:        userDto.Email,
		ClientID:     clientId,     // Set default or generate as needed
		ClientSecret: clientSecret, // Set default or generate as needed
	}
	return user
}
