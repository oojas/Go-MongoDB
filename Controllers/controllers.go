package Controllers

import (
	"GoMongo/Modals"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type UserController struct {
	session *mongo.Client
}

func NewController(client *mongo.Client) *UserController {
	return &UserController{client}
}
func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	cursor, _ := uc.session.Database("UsersDatabase").Collection("users").Find(context.TODO(), bson.M{})
	var userData []Modals.User
	if err := cursor.All(context.TODO(), &userData); err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
}
func (uc UserController) CreateUsers(w http.ResponseWriter, r *http.Request) {
	user := Modals.User{}
	json.NewDecoder(r.Body).Decode(&user)
	uc.session.Database("UsersDatabase").Collection("users").InsertOne(context.TODO(), user)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
