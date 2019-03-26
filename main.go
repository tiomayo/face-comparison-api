package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

func main() {
	
	dbConnSuccess := connectMongo()

	if dbConnSuccess {
		log.Println("Success connect to database")
	} else {
		log.Fatal("Cannot connect to database")
	}

	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/", IndexEndPoint).Methods("GET")
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

func IndexEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World !")
}
