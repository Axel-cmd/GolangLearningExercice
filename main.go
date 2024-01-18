package main

import (
	"estiam/db"
	"estiam/dictionary"
	"estiam/middleware"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var d *dictionary.Dictionary
var client *redis.Client

func main() {

	client = db.NewDatabaseClient()

	// err := client.Set(context.Background(), "example_key", "example_value", 10*time.Second).Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Exemple : obtenir la valeur de la clé depuis Redis
	// val, err := client.Get(context.Background(), "example_key").Result()
	// if err == redis.Nil {
	// 	fmt.Println("La clé n'existe pas")
	// } else if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println("Valeur de la clé:", val)
	// }

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
