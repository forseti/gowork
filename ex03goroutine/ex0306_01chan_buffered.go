package main

import (
	"fmt"
	"time"
)

func main() {
	buffChan := make(chan string, 3)
	go func() {
		defer close(buffChan)
		buffChan <- "Eko"
		buffChan <- "Kurniawan"
		buffChan <- "Khannedy"
	}()

	go func() {
		fmt.Println(<-buffChan, <-buffChan, <-buffChan)
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}
