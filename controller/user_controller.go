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

func Login(context *gin.Context) {
	jsonData, _ := io.ReadAll(context.Request.Body)
	var loginInfo request.LoginInfo
	// Use json.Unmarshal to parse the JSON into the struct.
	err := json.Unmarshal(jsonData, &loginInfo)
	if err != nil {
		log.Fatal(err)
	}
	loginResponse, err := service.Login(loginInfo)
	context.JSON(http.StatusOK, loginResponse)
}
func CheckToken(context *gin.Context) {
	jsonData, _ := io.ReadAll(context.Request.Body)
	isValid := service.CheckToken(string(jsonData))
	if isValid {
		context.JSON(http.StatusOK, true)
	} else {
		context.JSON(http.StatusOK, false)
	}
}
