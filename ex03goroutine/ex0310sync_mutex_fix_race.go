package main

import (
	"fmt"
	"github.com/forseti/gowork/ex03goroutine/helper"
	"sync"
	"time"
)

func main() {
	x := 0
	var mutex sync.Mutex

	helper.AsyncLoop(1000, func(index int) {
		for j := 0; j < 100; j++ {
			mutex.Lock()
			x++
			mutex.Unlock()
		}
	})

	time.Sleep(5 * time.Second)
	fmt.Println("x: ", x)
}
