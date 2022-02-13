package main

import (
	"fmt"
	"github.com/forseti/gowork/ex01basics/database"
	_ "github.com/forseti/gowork/ex01basics/unused" // use blank identifier if package is unused
)

func main() {
	var database = database.GetDatabase()
	fmt.Println("database =", database)
}
