package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/qwet700/Booking-doctor/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var doc models.Doctor
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	collection := client.Database("doc").Collection("docs")
	result, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := result.InsertedID
	doc.DocID = id.(primitive.ObjectID)
	json.NewEncoder(w).Encode(&doc)
}

func GetDoctorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := chi.URLParam(r, "docid")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid ID")
		return
	}

	var doc models.Doctor
	collection := client.Database("doc").Collection("docs")

	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&doc)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Doctor not found")
		return
	}
	json.NewEncoder(w).Encode(doc)
}

func UpdateDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	docID := chi.URLParam(r, "docid")

	var doc models.Doctor
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database("doc").Collection("docs")
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "_docid", Value: docID}}
	update := bson.M{"$set": doc} // auto convert ObjectID to string
	_, err = collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(doc)
}

func DeleteDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := chi.URLParam(r, "docid")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid ID")
		return
	}

	var doc models.Doctor
	collection := client.Database("doc").Collection("docs")

	err = collection.FindOneAndDelete(context.TODO(), bson.M{"_docid": id}).Decode(&doc)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("doc not found or no longer exist")
		return
	}
	json.NewEncoder(w).Encode(doc)
}
