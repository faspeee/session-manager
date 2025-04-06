package service

import (
	"encoding/json"
	"session-manager/model/request"
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

func Registry(user request.User) (bool, error) {

}
