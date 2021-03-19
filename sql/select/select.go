package main

import (
	"fmt"

	"github.com/raphael-rabadan/curso-golang/sql/conexao"
)

type usuario struct {
	id   int
	nome string
}

func main() {

	db := conexao.OpenConnection()
	defer conexao.CloseConnection(db)

	rows, _ := db.Query("SELECT id, nome FROM usuarios")
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}

}
