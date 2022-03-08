package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/forseti/gowork/ex05database/helper"
	"log"
	"time"
)

func main() {
	db := helper.GetConnection()
	defer helper.CloseConnection(db)

	ctx := context.Background()
	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	log.Printf("Query '%s'", query)
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("id:", id)
		fmt.Println("name:", name)
		if email.Valid {
			fmt.Println("email:", email.String)
		}
		fmt.Println("balance:", balance)
		fmt.Println("rating:", rating)
		if birthDate.Valid {
			fmt.Println("birthDate:", birthDate.Time)
		}
		fmt.Println("married:", married)
		fmt.Println("createdAt:", createdAt)
		fmt.Println("========================")
	}

	log.Println("Successfully query table customer")
}
