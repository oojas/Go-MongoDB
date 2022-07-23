package Modals

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `json:"name" bson:"name"`
	Age   int                `bson:"age" json:"age"`
	State string             `bson:"state" json:"state"`
	Count int                `bson:"count" json:"count"`
}
