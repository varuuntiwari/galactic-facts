package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	h "github.com/varuuntiwari/galactic-facts/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", h.Index).Methods("GET")
	router.HandleFunc("/api/fact", h.GetFacts).Methods("GET")

	http.Handle("/", router)
	srv := &http.Server {
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}