package handlers

import (
	_ "html/template"
	"net/http"
)

// HomeHandler Handler pour la page d'accueil (index.html)
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "templates/index.html")
}
