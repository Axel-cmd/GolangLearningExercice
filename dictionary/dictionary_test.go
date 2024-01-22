package dictionary

import (
	"estiam/db"
	"estiam/middleware"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHandleAddFunction struct{}

func (m *MockHandleAddFunction) HandleAdd(entry Entry, errorChan chan<- *middleware.APIError) {
	errorChan <- &middleware.APIError{Code: http.StatusInternalServerError, Message: "Erreur simulée lors de l'ajout d'une entrée en bdd"}
}

func TestHandleAdd(t *testing.T) {

	dict := &Dictionary{
		db: db.NewDatabaseClient(),
	}

	// test1 : ajout réussi
	entry := Entry{Word: "test", Definition: "test"}
	errorChan := make(chan *middleware.APIError, 1)
	go dict.HandleAdd(entry, errorChan)
	err := <-errorChan

	assert.Nil(t, err, "Erreur lors de l'ajout d'une entrée")

	testEntry, _ := dict.Get(entry.Word)
	assert.Equal(t, entry.Definition, testEntry.Definition, "L'entrée devrait être ajouté dans la bdd")

}

func TestHandleAddFail(t *testing.T) {
	// Cas de test 2: Échec d'ajout avec erreur de sauvegarde dans le fichier
	// Simuler une erreur lors de la sauvegarde dans le fichier (par exemple, en forçant une erreur)
	dict2 := &Dictionary{
		db:          db.NewDatabaseClient(),
		addFunction: &MockHandleAddFunction{},
	}

	entryWithError := Entry{Word: "test", Definition: "test"}
	errorChanWithError := make(chan *middleware.APIError, 1)
	go dict2.addFunction.HandleAdd(entryWithError, errorChanWithError)
	errWithError := <-errorChanWithError

	assert.NotNil(t, errWithError, "Une erreur devrait se produire lors de l'ajout avec une erreur de sauvegarde dans le fichier")
	assert.Equal(t, http.StatusInternalServerError, errWithError.Code, "Le code d'erreur devrait être 500")
}

func TestGet(t *testing.T) {
	dict := &Dictionary{
		db: db.NewDatabaseClient(),
	}

	entry := Entry{Word: "test", Definition: "test"}
	errorChan := make(chan *middleware.APIError, 1)
	go dict.HandleAdd(entry, errorChan)
	err := <-errorChan

	if err != nil {
		t.Fatalf("Erreur lors de l'ajout d'une entrée en bdd")
	}

	// cas 1 : Récupération d'un élément qui existe
	value, err := dict.Get(entry.Word)
	assert.Nil(t, err, "Erreur innatendu lors de la récupération de l'entrée existante")
	assert.Equal(t, entry.Word, value.Word, "La valeur récupérer ne correspond pas à la valeur attendue")

	//cas 2 : Récupération d'un élément inexistant
	keyNotExist := "notExist"
	valueNoExit, errNotExist := dict.Get(keyNotExist)
	assert.NotNil(t, errNotExist, "Une erreur devrait se produire si la clé n'existe pas")
	assert.Equal(t, "", valueNoExit.Word, "La valeur attendu pour une entrée inexistante doit être vide")

}

func TestDelete(t *testing.T) {

	dict := &Dictionary{
		db: db.NewDatabaseClient(),
	}

	entry := Entry{Word: "test", Definition: "test"}
	errorChan := make(chan *middleware.APIError, 1)

	go dict.HandleAdd(entry, errorChan)
	addErr := <-errorChan

	if addErr != nil {
		t.Fatalf("Erreur lors de l'ajout d'une clé dans la bdd")
	}

	// 	// Cas de test 1: Suppression réussie d'une entrée existante
	keyToDelete := "test"

	errorChanDel := make(chan *middleware.APIError, 1)

	go dict.HandleRemove(keyToDelete, errorChanDel)

	delErr := <-errorChanDel

	assert.Nil(t, delErr, "Erreur inattendue lors de la suppression de l'entrée existante")

	// Vérifier que l'entrée a été correctement supprimée
	_, isDeleteErr := dict.Get(keyToDelete)
	assert.NotNil(t, isDeleteErr, "L'entrée supprimée devrait être introuvable après la suppression")

	// 	// Cas de test 2: Suppression d'une entrée inexistante
	keyToDeleteNotExist := "notexist"

	errorChanDelNoExist := make(chan *middleware.APIError, 1)

	go dict.HandleRemove(keyToDeleteNotExist, errorChanDelNoExist)

	delNotExistErr := <-errorChanDelNoExist

	assert.NotNil(t, delNotExistErr, "Une erreur devrait se produire lors de la suppression d'une entrée inexistante")

}
