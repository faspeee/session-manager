package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"session-manager/configuration"
)

func GetUserByNameAndPassword(username string, password string) (bson.M, error) {

	var user bson.M
	err := configuration.MongoClient.Database("local").Collection("user").FindOne(context.TODO(),
		bson.D{{"username", "test"}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, err
}
