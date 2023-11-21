package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Phone     string             `json:"phone,omitempty" bson:"phone,omitempty"`
	OrderNumb int                `json:"ordernumb,omitempty" bson:"ordernumb,omitempty"`
}
