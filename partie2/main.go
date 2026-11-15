
package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Projet struct {
	ID string `json:"id"`
	Titre string `json:"titre"`
	Image string `json:"image"`
	Technologies []string `json:"technologies"`
	Introduction string `json:"introduction"`
	Description string `json:"description"`
	Fonctionnalites []string `json:"fonctionnalites"`
	Depot string `json:"depot"`
	Video string `json:"video"`
}

func lire() []Projet {
	fichier, _ := os.ReadFile("projets.json")

	var liste []Projet
	json.Unmarshal(fichier, &liste)

	return liste
}

func accueil(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	page, _ := template.ParseFiles("templates/index.html")
	page.Execute(w, nil)
}

func projets(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("templates/projects.html")
	page.Execute(w, lire())
}

func detail(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/projects/")

	for _, p := range lire() {
		if p.ID == id {
			page, _ := template.ParseFiles("templates/detail.html")
			page.Execute(w, p)
			return
		}
	}

	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/", accueil)
	http.HandleFunc("/projects", projets)
	http.HandleFunc("/projects/", detail)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets"))))

	println("Serveur : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
