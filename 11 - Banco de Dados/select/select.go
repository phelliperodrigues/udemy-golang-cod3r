package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")

	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	rows, _ := db.Query("select id, nome from usuarios where id > ? ", 5)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var u usuario
		err := rows.Scan(&u.id, &u.nome)
		if err != nil {
			return
		}
		fmt.Println(u)
	}
}
