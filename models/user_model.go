package models

import (
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct
type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName,omitempty" bson:"firstName"`
	LastName  string             `json:"lastName,omitempty" bson:"lastName"`
	Email     string             `json:"email,omitempty" bson:"email"`
	Mobile    string             `json:"mobile,omitempty" bson:"mobile"`
	Password  string             `json:"password,omitempty" bson:"password"`
	Username  string             `json:"username,omitempty" bson:"username"`
}
