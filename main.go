package main

import (
	"estiam/dictionary"
	"estiam/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var d *dictionary.Dictionary

func main() {

	var err error
	// créer l'objet dictionnaire
	d, err = dictionary.New()
	if err != nil {
		// s'il y a une erreur on log et l'application s'arrête
		log.Fatal(err)
	}

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
