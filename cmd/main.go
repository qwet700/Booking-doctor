package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/qwet700/Booking-doctor/pkg/controller"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://root:root1@localhost:27017/maxPoolSize=20&w=majority?authSource=admin"

func main() {
	var client *mongo.Client
	var err error
	r := chi.NewRouter()

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

	// Disconnected
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	defer fmt.Println("Disconnected")

	// Routes
	r.Use(middleware.Timeout(60 * time.Second))
	// r.Use(middleware.Logger)

	r.Post("/user", controller.CreateUser)
	r.Get("/users", controller.GetAllUsers)
	// r.Get("/user/{id}", controller.GetUsers)
	http.ListenAndServe(":8000", r)

}
