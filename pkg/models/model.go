package models

import (
	"context"
	"log"

	"github.com/qwet700/Booking-doctor/pkg/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Phone  string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Age    int                `json:"age,omitempty" bson:"age,omitempty"`
}

type Doctor struct {
	DocID primitive.ObjectID `json:"_docid,omitempty" bson:"_docid,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Prof  string             `json:"prof" bson:"prof"`
	Age   int                `json:"age" bson:"age"`
	Phone string             `json:"phone" bson:"phone"`
}

//	type Calender struct {
//		OrderID     User               `json:"_orderid" bson:"_orderid"`
//		DoctorID    Doctor             `json:"_doctorid" bson:"_doctorid"`
//		BookingDate primitive.DateTime `json:"bookingdate" bson:"bookingdate"`
//	}
var client = db.Dbconnect()

func NewUser(Name, Phone, Age string) {
	var user User
	collection := client.Database("user").Collection("users")

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err)
		return
	}
}
