package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/index", IndexEndPoint).Methods("GET")
	fmt.Println("Connected to port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}

func IndexEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World !")
}
