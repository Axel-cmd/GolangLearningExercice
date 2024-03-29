package main

import (
	"estiam/dictionary"
	"estiam/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var d *dictionary.Dictionary

//var c *redis.Client

func main() {

	// créer l'objet dictionnaire
	d = dictionary.New()

	r := mux.NewRouter().StrictSlash(true)

	// middleware de logging
	r.Use(middleware.LoggingMiddleware)

	// middleware d'authentication
	r.Use(middleware.AuthMiddelware)

	r.HandleFunc("/words/{word}", HandleGetWord).Methods("GET")
	r.HandleFunc("/words", HandleAddWord).Methods("POST")
	r.HandleFunc("/words/{word}", HandleDeleteWord).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))

	defer middleware.File.Close()
}
