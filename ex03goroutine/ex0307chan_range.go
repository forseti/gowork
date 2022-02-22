package main

import (
	"fmt"
	"strconv"
)

func main() {
	channel := make(chan string)

	go func() {
		defer close(channel)
		for i := 0; i < 10; i++ {
			channel <- "Iteration - " + strconv.Itoa(i)
		}
	}()

	for data := range channel {
		fmt.Println("Accepting data:", data)
	}
}
