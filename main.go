package main

import (
	"bufio"
	"estiam/dictionary"
	"fmt"
	"os"
)

func main() {
	d := dictionary.New()

	for {
		fmt.Println("Choisir une action :")
		fmt.Println("1. Ajouté un mot et sa définition")
		fmt.Println("2. Récupérer la définition d'un mot")
		fmt.Println("3. Supprimer un mot")
		fmt.Println("4. Lister les mots et les définitions du dictionnaire")
		fmt.Println("5. Sortir")

		var choice int
		fmt.Print("Faites votre choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			actionAdd(d)
		case 2:
			actionGet(d)
		case 3:
			actionRemove(d)
		case 4:
			actionList(d)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Choix invalide. Choississez un nombre entre 1 et 5.")
		}
	}
}

func actionAdd(d *dictionary.Dictionary) {
	// récupérer le mot
	word := getUserInput("Entrer un mot: ")
	// récupérer la définition
	definition := getUserInput("Entrer la définition du mot: ")
	// ajouté le mot
	d.Add(word, definition)
	fmt.Println("Mot ajouté dans le dictionnaire.")
}

func actionGet(d *dictionary.Dictionary) {
	word := getUserInput("Récupérer un mot: ")

	entry, err := d.Get(word)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("La définition du mot %s est : %s\n", word, entry.Definition)
}

func actionRemove(d *dictionary.Dictionary) {
	word := getUserInput("Quel mot souhaiter vous supprimer : ")

	d.Remove(word)
	fmt.Println("Mot supprimé avec succès.")
}

func actionList(d *dictionary.Dictionary) {
	ordered, entries := d.List()

	if len(entries) == 0 {
		fmt.Println("Le dictionnaire est vide.")
		return
	}

	fmt.Println("Les mots du dictionnaire : ")
	for _, word := range ordered {
		fmt.Printf("%s\n", word)
	}
}

// récupérer une entrée utilisateur
func getUserInput(text string) (result string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(text)
	input, _ := reader.ReadString('\n')
	result = input[:len(input)-1] // supprimer le caractère de retour de ligne

	return
}
