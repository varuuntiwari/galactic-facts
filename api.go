package main

import (
	"net/http"

	"github.com/gorilla/mux"
	h "github.com/varuuntiwari/galactic-facts/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", h.GetFacts).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}