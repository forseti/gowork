package main

import (
	"context"
	"fmt"
	"github.com/forseti/gowork/ex05database/entity"
	"github.com/forseti/gowork/ex05database/helper"
	"github.com/forseti/gowork/ex05database/repository"
)

func main() {
	db := helper.GetConnection()
	defer helper.CloseConnection(db)
	ctx := context.Background()

	commentRepository := repository.NewCommentRepository(db)

	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "TestRepository",
	}

	// Insert
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	// FindById
	resultById, err := commentRepository.FindById(ctx, result.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println("comment retrieved with FindById", result.Id, "=", resultById)

	// FindAll
	results, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range results {
		fmt.Println("Comment", comment)
	}
}
