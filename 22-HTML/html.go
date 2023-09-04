package main

import (
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	Name  string
	Email string
}

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"Marcio", "zeth.sell@gmail.com"}
		templates.ExecuteTemplate(w, "html.html", u)
	})

	log.Fatal(http.ListenAndServe(":5001", nil))
}
