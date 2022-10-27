package main

import (
	"database/sql"
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

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	_, _ = stmt.Exec("Ping", 1)
	_, _ = stmt.Exec("Kimmy", 2)
}
