package dictionary

import (
	"encoding/json"
	"estiam/middleware"
	"fmt"
	"net/http"
	"os"
)

// Représente une entrée dans le dictionnaire
type Entry struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

type SaveFunction interface {
	saveToFile() error
}

// Définition d'un dictionnaire
type Dictionary struct {
	file       string  // fichier ou est stocké le dictionnaire
	entries    []Entry // liste des entrées du dictionnaire
	addChan    chan Entry
	removeChan chan string
	saveFn     SaveFunction
}

// contructeur d'un objet Dictionnaire
func New() (*Dictionary, error) {
	d := &Dictionary{
		file:       "dictionary.json",
		addChan:    make(chan Entry),
		removeChan: make(chan string),
	}
	err := d.loadFromFile() // charger les données depuis le fichier
	if err != nil {
		return &Dictionary{}, err
	}
	return d, nil
}

func (d *Dictionary) HandleAdd(entry Entry, errorChan chan<- *middleware.APIError) {
	d.entries = append(d.entries, entry)

	err := d.saveToFile()
	if err != nil {
		errorChan <- &middleware.APIError{Code: http.StatusInternalServerError, Message: "erreur lors de la tentative d'ajout d'un mot dans le dictionnaire"}
		return
	}

	errorChan <- nil
}

// récupérer la définition d'un mot dans le dictionnaire
func (d *Dictionary) Get(word string) (Entry, *middleware.APIError) {
	for _, entry := range d.entries {
		if entry.Word == word {
			return entry, nil
		}
	}

	return Entry{}, &middleware.APIError{Code: http.StatusNotFound, Message: "le mot n'a pas été trouvé dans le dictionnaire"}
}

func (d *Dictionary) HandleRemove(word string, errorChan chan<- *middleware.APIError) {
	for i, entry := range d.entries {
		if entry.Word == word {
			d.entries = append(d.entries[:i], d.entries[i+1:]...)
			err := d.saveToFile()
			if err != nil {
				errorChan <- &middleware.APIError{Code: http.StatusInternalServerError, Message: "erreur lors de la tentative de suppression du mot dans le dictionnaire"}
				return
			}
			errorChan <- nil
			return
		}
	}
	errorChan <- &middleware.APIError{Code: http.StatusNotFound, Message: "impossible de supprimer le mot car il n'est pas présent dans le dictionnaire"}
}

func (d *Dictionary) loadFromFile() error {
	data, err := os.ReadFile(d.file)
	if err != nil {
		// si aucun fichier

		return fmt.Errorf("erreur lors de la tentative de lecture du fichier %s", d.file)

	}

	err = json.Unmarshal(data, &d.entries)
	if err != nil {
		return fmt.Errorf("erreur lors de la récupération des données dans le fichier dictionnaire : %s", d.file)
	}

	return nil
}

func (d *Dictionary) saveToFile() error {
	data, err := json.MarshalIndent(d.entries, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(d.file, data, 0644)

	return err
}
