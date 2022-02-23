package main

import (
	"context"
	"fmt"
)

func main() {
	background := context.Background()
	fmt.Println("background:", background)

	todo := context.TODO()
	fmt.Println("todo:", todo)
}
