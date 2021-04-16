package main

import (
	"html/template"
	"net/http"

	"github.com/nenitf/ead_go_alura_fundamentos-web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}
