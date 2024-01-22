package main

import (
	"bytes"
	"encoding/json"
	"estiam/db"
	"estiam/dictionary"
	"estiam/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetWordsRoute(t *testing.T) {
	// créer le dictionnaire
	d = &dictionary.Dictionary{
		Db: db.NewDatabaseClient(),
	}

	// créer une entrée dans le dictionnaire
	entry := dictionary.Entry{
		Word:       "example",
		Definition: "example",
	}
	errChan := make(chan *middleware.APIError, 1)

	go d.HandleAdd(entry, errChan)
	errAdd := <-errChan
	if errAdd != nil {
		t.Fatalf("Erreur lors de l'ajout d'une entrée !")
	}

	// cas 1 : réussite

	router := mux.NewRouter()

	router.HandleFunc("/words/{word}", HandleGetWord).Methods("GET")

	req, err := http.NewRequest("GET", "/words/example", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response dictionary.Entry
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, entry, response, "La Valeur devrait être égale")

	// cas 2 : 404 entity not found

	req, err = http.NewRequest("GET", "/words/fail_entries", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	rr = httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.NotEqual(t, dictionary.Entry{}, response, "La Valeur ne devrait pas être égale")
}

func TestPostWordsRoute(t *testing.T) {
	d = &dictionary.Dictionary{
		Db: db.NewDatabaseClient(),
	}

	// cas 1 : réussite

	router := mux.NewRouter()

	router.HandleFunc("/words", HandleAddWord).Methods("POST")

	// Créez des données JSON pour le test
	requestBody := []byte(`{"Word": "example", "Definition": "example"}`)

	req, err := http.NewRequest("POST", "/words", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response dictionary.Entry
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Vérifiez le contenu de la réponse
	expectedResponse := dictionary.Entry{
		Word:       "example",
		Definition: "example",
	}

	assert.Equal(t, expectedResponse, response)

}

func TestDeleteWordsRoute(t *testing.T) {

	d = &dictionary.Dictionary{
		Db: db.NewDatabaseClient(),
	}

	// créer une entrée dans le dictionnaire
	entry := dictionary.Entry{
		Word:       "example",
		Definition: "example",
	}
	errChan := make(chan *middleware.APIError, 1)

	go d.HandleAdd(entry, errChan)
	errAdd := <-errChan
	if errAdd != nil {
		t.Fatalf("Erreur lors de l'ajout d'une entrée !")
	}

	// cas 1 : suppression réussi

	router := mux.NewRouter()

	router.HandleFunc("/words/{word}", HandleDeleteWord).Methods("DELETE")

	req, err := http.NewRequest("DELETE", "/words/example", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}
