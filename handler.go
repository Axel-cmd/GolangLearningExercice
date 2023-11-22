package main

import (
	"encoding/json"
	"estiam/dictionary"
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
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(entry)
}

func HandleAddWord(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)

	var entry dictionary.Entry
	json.Unmarshal(reqBody, &entry)

	d.Add(entry)

	json.NewEncoder(w).Encode(entry)
}

func HandleDeleteWord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	word := vars["word"]

	d.Remove(word)
	json.NewEncoder(w).Encode("Supprimer avec succès")
}
