package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Phone string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Age   int                `json:"age,omitempty" bson:"age,omitempty"`
}

type Doctor struct {
	ID    primitive.ObjectID `json:"_docid" bson:"_docid"`
	Name  string             `json:"name" bson:"name"`
	Prof  string             `json:"prof" bson:"prof"`
	Age   int                `json:"age" bson:"age"`
	Phone string             `json:"phone" bson:"phone"`
}

type Calender struct {
	OrderID     int                `json:"orderID" bson:"orderID"`
	DoctorID    Doctor             `json:"doctorid" bson:"doctorid"`
	BookingDate primitive.DateTime `json:"bookingdate" bson:"bookingdate"`
}
