package main

import "github.com/raphael-rabadan/curso-golang/sql/conexao"

func main() {

	db := conexao.OpenConnection()
	defer conexao.CloseConnection(db)

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Uóxiton Istive", 1)
	stmt.Exec("Sheristone Uasleska", 2)

	stmt2, _ := db.Prepare("DELETE FROM usuarios WHERE id = ?")

	stmt2.Exec(1)
	stmt2.Exec(2)
}
