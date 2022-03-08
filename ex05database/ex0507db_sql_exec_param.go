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
		"eko":                         "password123",
		"budi":                        "password456",
		"joko":                        "password789",
		"hacker'; DROP TABLE USER; #": "whatever",
	}

	for username, password := range data {
		query := "INSERT INTO user(username, password) VALUES(?, ?)"

		log.Printf("Execute '%s'", query)

		_, err := db.ExecContext(ctx, query, username, password)
		if err != nil {
			panic(err)
		}
	}

	log.Println("Successfully insert a new customer")
}
