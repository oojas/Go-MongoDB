package Controllers

import (
	"GoMongo/Modals"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

type UserController struct {
	session *mongo.Client
}

func NewController(client *mongo.Client) *UserController {
	return &UserController{client}
}
func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 35*time.Second)
	cursor, _ := uc.session.Database("UsersDatabase").Collection("users").Find(ctx, bson.M{})
	var userData []Modals.User
	if err := cursor.All(ctx, &userData); err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
}
func (uc UserController) CreateUsers(w http.ResponseWriter, r *http.Request) {
	var user Modals.User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = primitive.NewObjectID()
	uc.session.Database("UsersDatabase").Collection("users").InsertOne(context.TODO(), user)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := bson.ObjectIdHex(mux.Vars(r)["id"])
	if err, _ := uc.session.Database("UserDatabase").Collection("users").DeleteOne(context.TODO(), id); err != nil {
		w.WriteHeader(404)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User ", id, " ")
}
