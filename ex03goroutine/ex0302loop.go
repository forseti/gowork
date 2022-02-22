package main

import (
	"github.com/forseti/gowork/ex03goroutine/helper"
	"time"
)

func main() {
	for i := 0; i < 100000; i++ {
		go helper.DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
