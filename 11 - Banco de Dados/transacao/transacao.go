package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

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

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values (?, ?)")

	_, _ = stmt.Exec(2000, "Phellipe")
	_, _ = stmt.Exec(2001, "Patty")
	_, err = stmt.Exec(1, "Ping") // chave duplicada

	if err != nil {
		_ = tx.Rollback()

		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		return
	}
}
