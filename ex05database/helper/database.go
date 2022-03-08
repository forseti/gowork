package helper

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // initialize MySQL driver
	"log"
	"time"
)

func GetConnection() *sql.DB {
	// use 'localhost' instead of '127.0.0.1'
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/gowork_db?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func CloseConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully close entity")
}
