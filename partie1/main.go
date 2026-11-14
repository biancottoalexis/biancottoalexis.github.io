package main 

import (
	"log"
	"net/http"
)

var _temp = template.Must(template.ParseGlob(". /template/*.html"))

func main() {
	router := http.NewserveMux()
	err := http.ListenAndServe("localhost:8080", router)
	if err !nil{
		log.Fatalf("Erreur(Serveur) : Impossible de lancer le serveur : %v" , err)
		return
	}
}

func renderTemplate(w http.ResponseWriter , name string , data interface{}) {
	//variable temp d'écriture du template 
	buffer := bytes.Buffer{}

	errRender := _temp.ExecuteTemplate(&buffer , name , data)
	if errRender !=nil {
		log.Printf
	}
}