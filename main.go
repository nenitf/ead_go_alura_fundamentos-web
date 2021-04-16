package main

import (
	"net/http"

	"github.com/nenitf/ead_go_alura_fundamentos-web/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
