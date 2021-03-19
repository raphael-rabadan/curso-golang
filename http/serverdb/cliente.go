package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/raphael-rabadan/curso-golang/sql/conexao"
)

type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa...:(")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db := conexao.OpenConnection()
	defer conexao.CloseConnection(db)

	var u Usuario
	error := db.QueryRow("select id,nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)

	jsonUsuarios := []byte{}
	if error == nil {
		jsonUsuarios, _ = json.Marshal(u)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonUsuarios))

}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db := conexao.OpenConnection()
	defer conexao.CloseConnection(db)

	rows, _ := db.Query("select id, nome from usuarios")

	var usuarios []Usuario

	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
