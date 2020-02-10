package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello wolrd"))
}

func connectDB() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://mongo-db:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

func main() {

	client := connectDB()

	if client != nil {
		log.Print("No null client returned")
	}
	r := mux.NewRouter()

	r.HandleFunc("/ping", Ping).
		Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
