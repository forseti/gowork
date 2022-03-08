package main

import (
	"context"
	"fmt"
	"github.com/forseti/gowork/ex05database/helper"
	"strconv"
)

func main() {
	db := helper.GetConnection()
	defer helper.CloseConnection(db)

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comment (email, comment) VALUES (?, ?)"

	for i := 0; i < 10; i++ {
		email := "budi" + strconv.Itoa(i) + "@gmail.com"
		comment := "Comment #" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id =", id)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
