package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"session-manager/model/request"
	"session-manager/service"
)

// Login handles user authentication requests.
// It reads the JSON-encoded login information from the request body,
// unmarshal it into a request.LoginInfo struct, and attempts to authenticate the user.
// If authentication is successful, it returns a login response; otherwise, it logs the error.
func Login(context *gin.Context) {
	jsonData, _ := io.ReadAll(context.Request.Body)
	var loginInfo request.LoginInfo
	// Unmarshal the JSON data into the loginInfo struct.
	err := json.Unmarshal(jsonData, &loginInfo)
	if err != nil {
		log.Fatal(err)
	}
	loginResponse, err := service.Login(loginInfo)
	context.JSON(http.StatusOK, loginResponse)
}

// CheckToken validates the provided authentication token.
// It reads the token from the request body and checks its validity using the service.CheckToken function.
// The function returns a JSON response indicating whether the token is valid.
func CheckToken(context *gin.Context) {
	jsonData, _ := io.ReadAll(context.Request.Body)
	isValid := service.CheckToken(string(jsonData))
	if isValid {
		context.JSON(http.StatusOK, true)
	} else {
		context.JSON(http.StatusForbidden, false)
	}
}

// Register handles new user registration requests.
// It reads the JSON-encoded user information from the request body,
// unmarshal it into a request.UserRequest struct, and attempts to register the user.
// Upon successful registration, it returns a registry response; otherwise, it logs the error.
func Register(context *gin.Context) {
	jsonData, _ := io.ReadAll(context.Request.Body)
	var userInfo request.UserRequest
	// Unmarshal the JSON data into the userInfo struct.
	err := json.Unmarshal(jsonData, &userInfo)
	if err != nil {
		log.Fatal(err)
	}
	registryResponse, err := service.Registry(userInfo)
	context.JSON(http.StatusOK, registryResponse)
}
