package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://root:root1@localhost:27017/maxPoolSize=20&w=majority?authSource=admin"

func main() {
	var client *mongo.Client
	var err error

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Disconnect the connection
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	defer fmt.Println("Disconnected")

	r := chi.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}
