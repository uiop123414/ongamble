package main

import (
	"context"
	"fmt"
	"log"
	"logger-service/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Client *mongo.Client
	Models data.Models
}

const (
	mongoURI = "mongodb://mongo:27017/"
	webPort = "80"
	gRpcPort = "50001"
)

func main() {
	client, err := connectMongo(mongoURI)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connected to mongodb")
	ctx, cancel := context.WithTimeout(context.Background(), 15 *time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Client: client,
		Models: data.New(client),
	}

	go app.gRPCListen()

	log.Println("Starting service on port ", webPort)
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func connectMongo(mongoURI string) (*mongo.Client, error) {
	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI(mongoURI)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

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
