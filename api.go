package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	h "github.com/varuuntiwari/galactic-facts/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", h.Index).Methods("GET")
	router.HandleFunc("/api/fact", h.GetFacts).Methods("GET")
	if _, err := os.LookupEnv("PORT"); !err {
		log.Panic("port not specified in env")
	}
	http.Handle("/", router)
	srv := &http.Server {
		Handler:		router,
		Addr:			":" + os.Getenv("PORT"),
		WriteTimeout:	10 * time.Second,
		ReadTimeout:	10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}