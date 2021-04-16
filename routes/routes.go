package routes

import (
	"net/http"

	"github.com/nenitf/ead_go_alura_fundamentos-web/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
