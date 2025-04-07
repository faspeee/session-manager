package configuration

import "go.mongodb.org/mongo-driver/mongo"

func GetCollectionByDb(collection string, database string) *mongo.Collection {
	return MongoClient.Database(database).Collection(collection)
}
