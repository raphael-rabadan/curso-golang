package conexao

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// OpenConnection Open the connection and return's it
func OpenConnection() *sql.DB {
	db, err := sql.Open("mysql", "admin:dr100fila**@tcp(drsemfiladb.cyzks1l7mtw0.sa-east-1.rds.amazonaws.com:3306)/cursogo")

	if err != nil {
		panic(err)
	}

	return db
}

// CloseConnection Close the connection. If an error occur, it'll stop the program
func CloseConnection(connection *sql.DB) {
	error := connection.Close()
	if error != nil {
		panic(error)
	}
}

func Exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
	return result
}
