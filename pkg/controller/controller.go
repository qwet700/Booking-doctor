package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/qwet700/Booking-doctor/pkg/models"
	"github.com/qwet700/Booking-doctor/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if client == nil {
		log.Println("MongoDB client is nil")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	collection := client.Database("user").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := result.InsertedID
	user.ID = id.(primitive.ObjectID)
	json.NewEncoder(w).Encode(&user)
	utils.JsonResponse(w, http.StatusCreated, user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	collection := client.Database("user").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(users)
}

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	idParam := chi.URLParam(r, "id")              // lấy giá trị id từ URL, với go-chi
// 	id, err := primitive.ObjectIDFromHex(idParam) // chuyển đổi id từ string sang ObjectID
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode("Invalid ID")
// 		return
// 	}

// 	var user models.User
// 	collection := client.Database("user").Collection("users")
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("User not found")
// 		return
// 	}

// 	json.NewEncoder(w).Encode(user)
// }
