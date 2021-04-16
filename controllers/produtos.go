package controllers

import (
	"net/http"
	"text/template"

	"github.com/nenitf/ead_go_alura_fundamentos-web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}
