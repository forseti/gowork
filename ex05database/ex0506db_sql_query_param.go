package main

import (
	"context"
	"fmt"
	"github.com/forseti/gowork/ex05database/helper"
	"log"
)

func main() {
	db := helper.GetConnection()
	defer helper.CloseConnection(db)

	ctx := context.Background()

	username := "admin'; #"
	password := "this_is_wrong_password"

	// Use query params to stop SQL injection.
	query := "SELECT username FROM user WHERE username = ? and password = ? LIMIT 1"
	fmt.Println(query)
	log.Printf("Query '%s'", query)
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Login successful")
	} else {
		fmt.Println("Login failed")
	}

	log.Println("Successfully query table customer")
}
