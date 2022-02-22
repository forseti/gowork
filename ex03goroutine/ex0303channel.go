package main

import (
	"fmt"
	"github.com/forseti/gowork/ex03goroutine/helper"
	"time"
)

func main() {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Eko Kurniawan Khannedy"
		fmt.Println("Finished sending data to channel")
	}()

	data := <-channel
	helper.HelloName(data)

	time.Sleep(5 * time.Second)
}
