# ead_go_alura_fundamentos-web

> Projeto referente a [este](https://www.alura.com.br/curso-online-go-lang-web) curso.

## Setup

### Banco

```sh
# criar base no postgresql
createdb -U postgres ead_go_alura_fundamentos-web
# senha padrão é 123456

# criar as tabelas
psql -U postgres -d ead_go_alura_fundamentos-web -a -f init.sql

# dicas do postgresql no terminal
# Entrar
psql -U postgres -d ead_go_alura_fundamentos-web

# \?                    exibe ajuda
# \q                    sai
# \l                    lista databases
# \c <databasename>     conecta uma database
# \dt                   lista tables da database
# \d <tablename>        descreve uma tabela
```

## Execução dos scripts

- Terminal com `go run main.go`
