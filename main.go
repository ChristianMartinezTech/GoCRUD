package main

import (
	"log"
	"net/http"
	"text/template"
)

var templ = template.Must(template.ParseGlob("templ/*"))

func Start(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "init", nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "create", nil)
}

func main() {
	http.HandleFunc("/", Start)
	http.HandleFunc("/create", Create)
	log.Println("Server runnning...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
