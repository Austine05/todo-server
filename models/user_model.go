package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct
type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Mobile    string             `json:"mobile,omitempty" bson:"mobile,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty"`
}
