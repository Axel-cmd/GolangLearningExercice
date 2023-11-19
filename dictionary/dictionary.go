package dictionary

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
)

// Représente une entrée dans le dictionnaire
type Entry struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

func (e Entry) String() string {
	return e.Definition
}

// Définition d'un dictionnaire
type Dictionary struct {
	entries []Entry // liste des entrées du dictionnaire
}

// contructeur d'un objet Dictionnaire
func New() *Dictionary {
	d := &Dictionary{}
	d.loadFromFile() // charger les données depuis le fichier
	return d
}

// ajouté un mot dans le dictionnaire
func (d *Dictionary) Add(word, definition string) {
	entry := Entry{Word: word, Definition: definition}
	d.entries = append(d.entries, entry)
	d.saveToFile()
}

// récupérer la définition d'un mot dans le dictionnaire
func (d *Dictionary) Get(word string) (Entry, error) {
	for _, entry := range d.entries {
		if entry.Word == word {
			return entry, nil
		}
	}

	return Entry{}, errors.New(fmt.Sprintf("Le mot %s n'a pas été trouvé dans le dictionnaire.", word))
}

// supprimer un mot dans le dictionnaire
func (d *Dictionary) Remove(word string) {
	for i, entry := range d.entries {
		if entry.Word == word {
			d.entries = append(d.entries[:i], d.entries[i+1:]...)
			d.saveToFile()
			return
		}
	}
}

func (d *Dictionary) List() ([]string, map[string]Entry) {
	sort.Slice(d.entries, func(i, j int) bool {
		return d.entries[i].Word < d.entries[j].Word
	})

	words := make([]string, len(d.entries))
	entriesMap := make(map[string]Entry)

	for i, entry := range d.entries {
		words[i] = entry.Word
		entriesMap[entry.Word] = entry
	}

	return words, entriesMap
}

func (d *Dictionary) loadFromFile() {
	data, err := os.ReadFile("dictionary.json")
	if err != nil {
		// si aucun fichier
		return
	}

	err = json.Unmarshal(data, &d.entries)
	if err != nil {
		panic(err)
	}
}

func (d *Dictionary) saveToFile() {
	data, err := json.MarshalIndent(d.entries, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("dictionary.json", data, 0644)
	if err != nil {
		panic(err)
	}
}
