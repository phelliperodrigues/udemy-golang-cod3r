package main

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"name"`
}

func UsuarioHandle(writer http.ResponseWriter, request *http.Request) {
	sid := strings.TrimPrefix(request.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)
	switch {
	case request.Method == "GET" && id > 0:
		usuarioPorId(writer, request, id)
	case request.Method == "GET":
		usuarioTodos(writer, request)
	default:
		writer.WriteHeader(http.StatusNotFound)
		_, err := fmt.Fprintf(writer, "Desculpa.... :(")
		if err != nil {
			return
		}
	}
}

func usuarioPorId(writer http.ResponseWriter, _ *http.Request, id int) {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var usuario Usuario

	err = db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&usuario.ID, &usuario.Nome)
	if err != nil {
		return
	}

	json, _ := json2.Marshal(usuario)

	writer.Header().Set("Content-Type", "application/json")
	_, err = fmt.Fprintf(writer, string(json))
	if err != nil {
		return
	}
}

func usuarioTodos(write http.ResponseWriter, _ *http.Request) {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	rows, _ := db.Query("select id, nome from usuarios")
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var usuarios []Usuario

	for rows.Next() {
		var usuario Usuario
		err := rows.Scan(&usuario.ID, &usuario.Nome)
		if err != nil {
			return
		}
		usuarios = append(usuarios, usuario)
	}

	json, _ := json2.Marshal(usuarios)
	write.Header().Set("Content-Type", "application/json")
	_, err = fmt.Fprintf(write, string(json))
	if err != nil {
		return
	}
}
