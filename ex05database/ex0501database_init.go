package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // initialize MySQL driver
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/gowork_db")
	defer db.Close()

	if err != nil {
		panic(err)
	}
}
