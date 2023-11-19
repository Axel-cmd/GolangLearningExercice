package dictionary

import (
	"errors"
	"fmt"
	"sort"
)

type Entry struct {
	Definition string
}

func (e Entry) String() string {
	return e.Definition
}

// Définition d'un dictionnaire
type Dictionary struct {
	entries map[string]Entry
}

// contructeur d'un objet Dictionnaire
func New() *Dictionary {
	return &Dictionary{
		entries: make(map[string]Entry),
	}
}

// ajouté un mot dans le dictionnaire
func (d *Dictionary) Add(word string, definition string) {
	entry := Entry{Definition: definition}
	d.entries[word] = entry
}

// récupérer la définition d'un mot dans le dictionnaire
func (d *Dictionary) Get(word string) (Entry, error) {
	entry, found := d.entries[word]
	if !found {
		return Entry{}, errors.New(fmt.Sprintf("Le mot %s n'a pas été trouvé dans le dictionnaire.", word))
	}
	return entry, nil
}

// supprimer un mot dans le dictionnaire
func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
}

func (d *Dictionary) List() ([]string, map[string]Entry) {

	var wordList []string
	for _, word := range d.entries {
		wordList = append(wordList, word.String())
	}
	sort.Strings(wordList)

	var result []string
	for _, word := range wordList {
		result = append(result, fmt.Sprintf("%s: %s", word, d.entries[word].String()))
	}

	return result, d.entries
}
