package main

import (
	"log"

	"github.com/raphael-rabadan/curso-golang/sql/conexao"
)

func main() {
	db := conexao.OpenConnection()
	defer conexao.CloseConnection(db)

	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("INSERT INTO usuarios (id, nome) VALUES (?, ?)")
	stmt.Exec(2002, "Bia")
	stmt.Exec(2003, "Carlos")

	_, err := stmt.Exec(1, "Tiago") //Duplicado forçando errado

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
