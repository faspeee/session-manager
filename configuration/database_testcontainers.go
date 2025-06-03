package configuration

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func CreateAndCloseTestContainerMongo() {
	mongodbContainer, err := mongodb.Run(context.TODO(), "mongo:6")

	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}
	endpoint, err := mongodbContainer.ConnectionString(context.TODO())
	if err != nil {
		log.Printf("failed to get connection string: %s", err)
		return
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(endpoint))
	if err != nil {
		log.Printf("failed to connect to MongoDB: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	MongoClient = client
	defer func() {
		if err := testcontainers.TerminateContainer(mongodbContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
}
