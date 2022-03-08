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

	data := map[string]string{
		"eko":  "Eko",
		"budi": "Budi",
		"joko": "Joko",
	}

	for id, name := range data {
		query := fmt.Sprintf("INSERT INTO customer(id, name) VALUES('%s', '%s')", id, name)

		log.Printf("Execute '%s'", query)

		_, err := db.ExecContext(ctx, query)
		if err != nil {
			panic(err)
		}
	}

	log.Println("Successfully insert a new customer")
}
