package main

import (
	"fmt"
	"github.com/forseti/gowork/ex03goroutine/helper"
	"sync"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func main() {
	once := sync.Once{}
	group := sync.WaitGroup{}

	helper.AsyncLoop(100, func(index int) {
		group.Add(1)
		once.Do(OnlyOnce)
		group.Done()
	})

	group.Wait()
	fmt.Println("Done. Counter:", counter)
}
