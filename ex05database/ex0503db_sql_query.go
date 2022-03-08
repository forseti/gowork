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
	query := "SELECT id, name FROM customer"
	log.Printf("Query '%s'", query)
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Printf("id: %s, name: %s\n", id, name)
	}

	log.Println("Successfully query table customer")
}
