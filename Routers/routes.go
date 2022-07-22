package Routers

import (
	"GoMongo/Controllers"
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":3000"
}
func Start() {
	p := getPort()
	router := mux.NewRouter()
	session := Controllers.NewController(GetSession())
	router.HandleFunc("/users", session.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/create-user", session.CreateUsers).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(p, router))
}

// Connection with MongoDB
func GetSession() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://ojas:iamastarboy13@engineeringguide.h2qbook.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}