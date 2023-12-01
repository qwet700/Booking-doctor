package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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

// type Calender struct {
// 	OrderID     User               `json:"_orderid" bson:"_orderid"`
// 	DoctorID    Doctor             `json:"_doctorid" bson:"_doctorid"`
// 	BookingDate primitive.DateTime `json:"bookingdate" bson:"bookingdate"`
// }
