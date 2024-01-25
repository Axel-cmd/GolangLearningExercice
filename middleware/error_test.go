package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorMiddlewareWithPanic(t *testing.T) {
	// Créez un gestionnaire factice qui panique
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("Panic simulé dans le handler")
	})

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Créez un enregistrement de réponse HTTP simulé
	rr := httptest.NewRecorder()

	// Appliquez le middleware d'erreur à la requête
	ErrorMiddleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Le code d'erreur devrait être égale à 500")

}

func TestErrorMiddlewareWithExplicitError(t *testing.T) {
	// Créez un gestionnaire factice qui renvoie une erreur explicite
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RespondWithError(w, http.StatusNotFound, "Resource not found")
	})

	// Créez une requête HTTP simulée
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Créez un enregistrement de réponse HTTP simulé
	rr := httptest.NewRecorder()

	// Appliquez le middleware d'erreur à la requête
	ErrorMiddleware(handler).ServeHTTP(rr, req)

	// Vérifiez que le code de statut de la réponse est Not Found (404)
	assert.Equal(t, http.StatusNotFound, rr.Code)

	var apiError APIError
	err = json.Unmarshal(rr.Body.Bytes(), &apiError)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, apiError.Code)
	assert.Equal(t, "Resource not found", apiError.Message)
}
