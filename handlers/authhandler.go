package handlers

import (
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	// Vérifie que la méthode est POST
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Simple validation (à remplacer par une vraie validation en production)
		if username == "admin" && password == "password" {
			http.Redirect(w, r, "/welcome", http.StatusSeeOther)
		} else {
			http.Error(w, "Nom d'utilisateur ou mot de passe incorrect", http.StatusUnauthorized)
		}
	}
}
