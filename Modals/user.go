package Modals

type User struct {
	Name  string `json:"name" bson:"name"`
	Age   int    `bson:"age" json:"age"`
	State string `bson:"state" json:"state"`
	Count int    `bson:"count" json:"count"`
}
