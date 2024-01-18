package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWordsRoute(t *testing.T) {

	// d = &dictionary.Dictionary{
	// 	entries: []dictionary.Entry{
	// 		{Word: "test", Definition: "test"},
	// 	},
	// }

	server := httptest.NewServer(http.HandlerFunc(HandleGetWord))
	defer server.Close()

	response, err := http.Get(server.URL + "/words/test")
	if err != nil {
		t.Fatal(err)
	}

	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode, "Le ")
}

func TestPostWordsRoute(t *testing.T) {

}

func TestDeleteWordsRoute(t *testing.T) {

}
