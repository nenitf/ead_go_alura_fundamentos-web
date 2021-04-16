package db

import (
	"database/sql"

	// go get github.com/lib/pq
	_ "github.com/lib/pq" // é usado em tempo de execução, não no código
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=ead_go_alura_fundamentos-web password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
