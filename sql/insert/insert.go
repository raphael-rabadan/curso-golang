package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raphael-rabadan/curso-golang/sql/conexao"
)

func main() {
	db := conexao.OpenConnection()
	defer conexao.CloseConnection(db)

	stmt, _ := db.Prepare("INSERT INTO usuarios (nome) VALUES (?)")

	stmt.Exec("Pedro")
	stmt.Exec("Maria")
	res, error := stmt.Exec("João")
	if error != nil {
		panic(error)
	}

	id, _ := res.LastInsertId()
	fmt.Println("Inserido o id ", id)

	lines, _ := res.RowsAffected()
	fmt.Println("Linhas afetadas", lines)

	stmt.Close()

}
