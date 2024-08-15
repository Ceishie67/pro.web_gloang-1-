package main

import (
	"net/http"
	"webcligo/handlers"
)

func main() {
	// Route pour afficher la page de login
	http.HandleFunc("/login", handlers.LoginHandler)

	// Route pour traiter le formulaire de login
	http.HandleFunc("/auth", handlers.AuthHandler)

	// Route pour afficher la page de register
	http.HandleFunc("/register", handlers.RegisterHandler)

	// Sert les fichiers statiques (CSS, images, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Redirect root URL to /login
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/register", http.StatusFound)
	})

	// Démarre le serveur sur le port 8080
	http.ListenAndServe(":8080", nil)
}
