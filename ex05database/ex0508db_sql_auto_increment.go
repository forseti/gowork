package main

import (
	"context"
	"github.com/forseti/gowork/ex05database/helper"
	"log"
)

func main() {
	db := helper.GetConnection()
	defer helper.CloseConnection(db)

	ctx := context.Background()

	data := map[string]string{
		"eko@gmail.com":  "Test komen",
		"budi@gmail.com": "hello, eko",
		"joko@gmail.com": "hello, budi and eko",
	}

	for email, comment := range data {

		query := "INSERT INTO comment(email, comment) VALUES(?, ?)"
		log.Printf("Execute '%s'", query)
		result, err := db.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		log.Println("Successfully insert a new comment with id", insertId)

	}
}
