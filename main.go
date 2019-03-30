package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/tiomayo/face-comparison-api/controller"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	connectMongo()
}

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/identify", controller.Identify).Methods("POST")
	r.HandleFunc("/go/aisatsu", controller.Aisatsu).Methods("GET")
	// r.HandleFunc("/upload", controller.UploadImage).Methods("POST")
	log.Println("Connected to port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}

func connectMongo() bool {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	errConnect := client.Connect(ctx)

	if errConnect != nil {
		return false
	}

	return true
}
