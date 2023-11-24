package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Dbconnect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://root:root1@localhost:27017/maxPoolSize=20&w=majority?authSource=admin")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Connection Failed to Database")
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Connection Failed to Database")
		log.Fatal(err)
	}
	fmt.Println("Connected to Database")
	return client
}
