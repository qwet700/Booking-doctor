package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/qwet700/Booking-doctor/pkg/db"
	"github.com/qwet700/Booking-doctor/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client = db.Dbconnect()

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	collection := client.Database("user").Collection("users")
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := result.InsertedID
	user.UserID = id.(primitive.ObjectID)
	json.NewEncoder(w).Encode(&user)
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

func GetUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")              // lấy giá trị id từ URL, với go-chi
	id, err := primitive.ObjectIDFromHex(idParam) // chuyển đổi id từ string sang ObjectID
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid ID")
		return
	}

	var user models.User
	collection := client.Database("user").Collection("users")

	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
		return
	}
	json.NewEncoder(w).Encode(user)
}

// func GetUserName(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	nameParam := chi.URLParam(r, "name") // lấy giá trị id từ URL, với go-chi
// 	var err error
// 	var user models.User
// 	collection := client.Database("user").Collection("users")
// 	err = collection.FindOne(context.TODO(), bson.M{"name": nameParam}).Decode(&user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("User not found")
// 		return
// 	}

// 	json.NewEncoder(w).Encode(user)
// }

// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	idParam := chi.URLParam(r, "id")
// 	id, err := primitive.ObjectIDFromHex(idParam)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode("Invalid ID")
// 		return
// 	}

// 	var user models.User
// 	collection := client.Database("user").Collection("users")

// 	err = collection.FindOneAndDelete(context.TODO(), bson.M{"_id": id}).Decode(&user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("User not found or no longer exist")
// 		return
// 	}
// 	json.NewEncoder(w).Encode(user)
// }

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := chi.URLParam(r, "id")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database("user").Collection("users")
	opts := options.Update().SetUpsert(true)
	filter := bson.D{primitive.E{Key: "_id", Value: userID}}
	update := bson.M{"$set": user}
	_, err = collection.UpdateOne(context.TODO(), filter, update, opts) // result, err :=
	if err != nil {
		log.Fatal(err)
	}
	// if result.MatchedCount != 0 {
	// 	fmt.Println("matched and replaced an existing document")
	// 	return
	// }
	// if result.UpsertedCount != 0 {
	// 	fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	// }
	json.NewEncoder(w).Encode(user)
}
