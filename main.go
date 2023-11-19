package main

import (
	"fmt"
	"sort"
)

func main() {

	wordMap := make(map[string]string)

	word := "zigzag"
	definition := "Ligne brisée"

	word2 := "parapluie"
	definition2 := "Objet portatif constitué par une étoffe tendue sur une armature pliante à manche, et qui sert d'abri contre la pluie."

	add(wordMap, word, definition)
	fmt.Printf("%s ajouté dans le dictionnaire..\n", word)

	add(wordMap, word2, definition2)
	fmt.Printf("%s ajouté dans le dictionnaire..\n", word2)

	fmt.Println(get(wordMap, word))

	wordList := list(wordMap)
	fmt.Println(wordList)

	remove(wordMap, word)
	fmt.Printf("%s supprimé du dictionnaire...\n", word)

	fmt.Println(get(wordMap, word))

}

// ajouté un mot dans la map
func add(wordMap map[string]string, word string, definition string) {
	wordMap[word] = definition
}

func get(wordMap map[string]string, word string) (result string) {
	value, ok := wordMap[word]

	if !ok {
		result = fmt.Sprintf("Le mot %s n'est pas dans le dictionnaire", word)
	} else {
		result = value
	}
	return
}

func remove(wordMap map[string]string, word string) {
	delete(wordMap, word)
}

func list(wordMap map[string]string) (result []string) {
	var wordList []string
	for word := range wordMap {
		wordList = append(wordList, word)
	}
	sort.Strings(wordList)

	for _, word := range wordList {
		result = append(result, fmt.Sprintf("%s: %s", word, wordMap[word]))
	}

	return result
}
