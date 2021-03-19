package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/raphael-rabadan/curso-golang/sql/conexao"
)

func main() {

	db := conexao.OpenConnection()
	defer conexao.CloseConnection(db)

	conexao.Exec(db, "create database if not exists cursogo")
	conexao.Exec(db, "use cursogo")
	// exec(db, "drop table if exists usuarios")
	conexao.Exec(db, `CREATE TABLE if not exists usuarios (
				id integer auto_increment,
				nome varchar(80),
				PRIMARY KEY (id)
			)`)

}
