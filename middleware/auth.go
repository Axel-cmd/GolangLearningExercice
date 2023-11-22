package middleware

import (
	"net/http"
)

var expectedToken string = "token"

// middleware d'authentication
func AuthMiddelware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Récupérer le jeton d'authentification de l'en-tête
		token := req.Header.Get("Authorization")

		// Vérifier la présence du jeton
		if token == "" {
			http.Error(w, "Token manquant", http.StatusUnauthorized)
			return
		}

		// Vérifier la validité du jeton (vérification basique ici, adaptez selon vos besoins)
		if !isValidToken(token) {
			http.Error(w, "Token invalide", http.StatusUnauthorized)
			return
		}

		// Authentification réussie, passer à la prochaine étape du middleware
		next.ServeHTTP(w, req)
	})
}

// valider le token attendu
func isValidToken(token string) bool {
	return token == expectedToken
}
