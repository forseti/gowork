package main

import (
	"fmt"
	"github.com/forseti/gowork/ex03goroutine/helper"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go helper.GiveMeResponse(channel1)
	go helper.GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data from channel1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data from channel2", data)
			counter++
		default:
			fmt.Println("Waiting for data")
			time.Sleep(2 * time.Second)
		}

		if counter == 2 {
			break
		}
	}
}
