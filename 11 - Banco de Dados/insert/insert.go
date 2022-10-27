package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")
	_, _ = stmt.Exec("Maria")
	_, _ = stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")
	id, _ := res.LastInsertId()
	fmt.Println("Id criado", id)

	linhas, _ := res.RowsAffected()
	fmt.Println("Quantidade de linhas afetadas", linhas)
}
