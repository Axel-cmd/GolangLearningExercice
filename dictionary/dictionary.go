package dictionary

import (
	"context"
	"estiam/db"
	"estiam/middleware"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

type AddFunction interface {
	HandleAdd(entry Entry, errorChan chan<- *middleware.APIError)
}

// Représente une entrée dans le dictionnaire
type Entry struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

// Définition d'un dictionnaire
type Dictionary struct {
	Db          *redis.Client
	AddFunction AddFunction
}

// contructeur d'un objet Dictionnaire
func New() *Dictionary {
	return &Dictionary{
		Db: db.NewDatabaseClient(),
	}
}

func (d *Dictionary) HandleAdd(entry Entry, errorChan chan<- *middleware.APIError) {
	// ajouté l'entrée dans la base
	err := d.Db.Set(context.Background(), entry.Word, entry.Definition, 10*time.Second).Err()
	// si erreur on renvoie une erreur
	if err != nil {
		errorChan <- &middleware.APIError{Code: http.StatusInternalServerError, Message: "erreur lors de la tentative d'ajout d'un mot dans le dictionnaire"}
		return
	}
	// sinon nil
	errorChan <- nil
}

// récupérer la définition d'un mot dans le dictionnaire
func (d *Dictionary) Get(word string) (Entry, *middleware.APIError) {
	val, err := d.Db.Get(context.Background(), word).Result()

	if err == redis.Nil {
		return Entry{}, &middleware.APIError{Code: http.StatusNotFound, Message: "le mot n'a pas été trouvé dans le dictionnaire"}
	} else if err != nil {
		log.Fatal(err)
	}

	return Entry{Word: word, Definition: val}, nil
}

func (d *Dictionary) HandleRemove(word string, errorChan chan<- *middleware.APIError) {
	result := d.Db.Del(context.Background(), word)

	if result.Err() != nil {
		errorChan <- &middleware.APIError{Code: http.StatusInternalServerError, Message: "erreur lors de la tentative de suppression du mot dans le dictionnaire"}
		return
	}

	if result.Val() < 1 {
		errorChan <- &middleware.APIError{Code: http.StatusNotFound, Message: "l'élément que vous souhaitez supprimé n'existe pas"}
		return
	}

	errorChan <- nil
}
