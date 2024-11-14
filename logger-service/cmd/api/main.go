package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Client *mongo.Client
}

const mongoURI = "mongodb://mongo:27017/"

func main() {
	client, err := connectMongo(mongoURI)
	if err != nil {
		log.Panic(err)
	}

	_ = Config{
		Client: client,
	}
}

func connectMongo(mongoURI string) (*mongo.Client, error) {
	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}


	return client, nil
}
