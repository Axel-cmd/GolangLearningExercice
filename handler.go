package main

import (
	"encoding/json"
	"estiam/dictionary"
	"estiam/middleware"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// handler pour récupérer la définition d'un mot
func HandleGetWord(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	word := vars["word"]

	entry, err := d.Get(word)

	if err != nil {
		middleware.RespondWithError(w, err.Code, err.Message)
		return
	}
	json.NewEncoder(w).Encode(entry)
}

func HandleAddWord(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)

	var entry dictionary.Entry
	json.Unmarshal(reqBody, &entry)

	errorChan := make(chan *middleware.APIError, 1)

	go d.HandleAdd(entry, errorChan)

	// Attendre la réponse du canal d'erreur
	if err := <-errorChan; err != nil {
		// Si une erreur est reçue, renvoyer une réponse d'erreur
		middleware.RespondWithError(w, err.Code, err.Message)
		return
	}

	json.NewEncoder(w).Encode(entry)
}

func HandleDeleteWord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	word := vars["word"]

	errorChan := make(chan *middleware.APIError, 1)

	go d.HandleRemove(word, errorChan)

	// Attendre la réponse du canal d'erreur
	if err := <-errorChan; err != nil {
		// Si une erreur est reçue, renvoyer une réponse d'erreur
		middleware.RespondWithError(w, err.Code, err.Message)
		return
	}

	json.NewEncoder(w).Encode("Supprimer avec succès")
}
