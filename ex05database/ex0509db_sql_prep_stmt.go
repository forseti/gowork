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
	script := "INSERT INTO comment (email, comment) VALUES (?, ?)"
	stmt, err := db.PrepareContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Comment #" + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id =", id)
	}
}
