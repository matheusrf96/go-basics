package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
}

func root(w http.ResponseWriter, r *http.Request) {
	u := user{"Matheus", "matheusrf96@gmail.com"}
	templates.ExecuteTemplate(w, "index.html", u)
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", root)

	log.Println("Server running at: 0.0.0.0:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
