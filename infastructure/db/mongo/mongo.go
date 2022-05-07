package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

// Connecting to mongo database

func NewConnection(mongoUri string) (*mongo.Client, error) {
	log.Printf("Connecting to MongoDB on: %v", mongoUri)

	// Connecting to mongo uri
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))

	// Defining a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")

	return client, nil
}
