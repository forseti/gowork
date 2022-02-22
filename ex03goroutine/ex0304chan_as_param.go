package main

import (
	"github.com/forseti/gowork/ex03goroutine/helper"
	"time"
)

func main() {
	channel := make(chan string)
	defer close(channel)

	go helper.GiveMeResponse(channel)

	data := <-channel
	helper.HelloName(data)

	time.Sleep(5 * time.Second)
}
