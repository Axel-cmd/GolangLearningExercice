package middleware

import (
	"log"
	"net/http"
	"os"
)

var File *os.File

func LoggingMiddleware(next http.Handler) http.Handler {
	var err error
	// Ouvrir le fichier de log en mode append
	File, err = os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Utiliser le fichier de log comme sortie pour le logger
	log.SetOutput(File)

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("[%s] %s %s %s", req.Method, req.RemoteAddr, req.URL, req.UserAgent())
		next.ServeHTTP(w, req)
	})

}
