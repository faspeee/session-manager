package configuration

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Your MongoDB Atlas Connection String
const uri = "mongodb://host.docker.internal:27017/?directConnection=true&serverSelectionTimeoutMS=3000"

// A global variable that will hold a reference to the MongoDB client
var MongoClient *mongo.Client

// Init The init function will run before our main function to establish a connection to MongoDB. If it cannot connect it will fail and the program will exit.
func Init() {
	if err := connectToMongodb(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
}

// Our implementation logic for connecting to MongoDB
func connectToMongodb() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	MongoClient = client
	return err
}
