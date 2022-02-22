package main

import (
	"github.com/forseti/gowork/ex03goroutine/helper"
	"log"
	"time"
)

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	log.Println("OnlyIn")
	channel <- "Eko Kurniawan Khannedy"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	log.Println("OnlyOut")
	helper.HelloName(data)
}

func main() {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}
