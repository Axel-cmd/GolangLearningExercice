package dictionary

import (
	"errors"
	"estiam/middleware"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockSaveFunction struct{}

func (m *MockSaveFunction) saveToFile() error {
	return errors.New("Erreur simulée lors de la sauvegarde dans le fichier")
}

func TestHandleAdd(t *testing.T) {
	// Créer un fichier de test
	tempFile, fileErr := os.CreateTemp("", "dictionary_test.json")
	if fileErr != nil {
		t.Fatalf("Erreur lors de la création du fichier de test: %v", fileErr)
	}
	defer os.Remove(tempFile.Name()) // Supprimer le fichier de test après le test

	dict := &Dictionary{
		file: tempFile.Name(),
	}

	// test1 : ajout réussi
	entry := Entry{Word: "test", Definition: "test"}
	errorChan := make(chan *middleware.APIError, 1)
	go dict.HandleAdd(entry, errorChan)
	err := <-errorChan

	assert.Nil(t, err, "Erreur lors de l'ajout d'une entrée")
	assert.Len(t, dict.entries, 1, "Le nombre d'entrée dans le dictionnaire devrait être 1")

	// vérifier le contenu du fichier json
	content, readErr := os.ReadFile(tempFile.Name())
	if readErr != nil {
		t.Fatalf("Erreur lors de la lecture du fichier de test: %v", err)
	}
	// ce qui est attendu
	expectedJSON := `[{"word":"test","definition":"test"}]`
	assert.JSONEq(t, expectedJSON, string(content), "Le contenu du fichier JSON ne correspond pas à ce qui était attendu")

	// Cas de test 2: Échec d'ajout avec erreur de sauvegarde dans le fichier
	// Simuler une erreur lors de la sauvegarde dans le fichier (par exemple, en forçant une erreur)
	dict = &Dictionary{
		saveFn: &MockSaveFunction{},
	}

	entryWithError := Entry{Word: "test", Definition: "test"}
	errorChanWithError := make(chan *middleware.APIError, 1)
	go dict.HandleAdd(entryWithError, errorChanWithError)
	errWithError := <-errorChanWithError
	assert.NotNil(t, errWithError, "Une erreur devrait se produire lors de l'ajout avec une erreur de sauvegarde dans le fichier")
	assert.Equal(t, http.StatusInternalServerError, errWithError.Code, "Le code d'erreur devrait être 500")
	assert.Len(t, dict.entries, 1, "Le nombre d'entrées dans le dictionnaire ne devrait pas changer en cas d'erreur")
}

func TestGet(t *testing.T) {
	dict := &Dictionary{
		entries: []Entry{
			{Word: "test", Definition: "test"},
		},
	}

	// cas 1 : Récupération d'un élément qui existe
	key := "test"
	value, err := dict.Get(key)
	assert.Nil(t, err, "Erreur innatendu lors de la récupération de l'entrée existante")
	assert.Equal(t, key, value.Word, "La valeur récupérer ne correspond pas à la valeur attendue")

	//cas 2 : Récupération d'un élément inexistant
	keyNotExist := "notExist"
	valueNoExit, errNotExist := dict.Get(keyNotExist)
	assert.NotNil(t, errNotExist, "Une erreur devrait se produire si la clé n'existe pas")
	assert.Equal(t, "", valueNoExit.Word, "La valeur attendu pour une entrée inexistante doit être vide")

}

func TestDelete(t *testing.T) {
	tempFile, tempErr := os.CreateTemp("", "dictionary_test.json")
	if tempErr != nil {
		t.Fatalf("Erreur lors de la création du fichier de test: %v", tempErr)
	}
	defer os.Remove(tempFile.Name()) // Supprimer le fichier de test après le test

	dict := &Dictionary{
		file: tempFile.Name(),
	}

	entry := Entry{Word: "test", Definition: "test"}
	errorChan := make(chan *middleware.APIError, 1)
	go dict.HandleAdd(entry, errorChan)
	go dict.HandleAdd(Entry{Word: "exemple", Definition: "exemple"}, errorChan)
	addErr := <-errorChan
	if addErr != nil {
		t.Fatalf("Erreur lors de l'ajout d'une clé dans le fichier %v", tempFile.Name())
	}

	// Cas de test 1: Suppression réussie d'une entrée existante
	keyToDelete := "test"

	errorChanDel := make(chan *middleware.APIError, 1)

	go dict.HandleRemove(keyToDelete, errorChanDel)

	delErr := <-errorChanDel

	assert.Nil(t, delErr, "Erreur inattendue lors de la suppression de l'entrée existante")
	assert.Len(t, dict.entries, 1, "Le nombre d'entrées dans le dictionnaire devrait être ajusté après la suppression")

	// Vérifier que l'entrée a été correctement supprimée
	_, isDeleteErr := dict.Get(keyToDelete)
	assert.NotNil(t, isDeleteErr, "L'entrée supprimée devrait être introuvable après la suppression")

	// Cas de test 2: Suppression d'une entrée inexistante
	keyToDeleteNotExist := "notexist"

	errorChanDelNoExist := make(chan *middleware.APIError, 1)

	go dict.HandleRemove(keyToDeleteNotExist, errorChanDelNoExist)

	delNotExistErr := <-errorChanDelNoExist

	assert.NotNil(t, delNotExistErr, "Une erreur devrait se produire lors de la suppression d'une entrée inexistante")
	assert.Len(t, dict.entries, 1, "Le nombre d'entrées dans le dictionnaire ne devrait pas changer après la suppression d'une entrée inexistante")

}
