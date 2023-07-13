package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Todo struct
type Todo struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	Deadline    time.Time          `json:"deadline,omitempty" bson:"deadline,omitempty"`
	UserID		primitive.ObjectID `json:"userId,oi=mitempty" bson:"userId,omitempty`
}

