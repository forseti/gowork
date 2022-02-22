package main

import (
	"fmt"
	"github.com/forseti/gowork/ex03goroutine/helper"
	"time"
)

func main() {
	x := 0

	helper.AsyncLoop(1000, func(index int) {
		for j := 0; j < 100; j++ {
			x++
		}
	})

	time.Sleep(5 * time.Second)
	fmt.Println("x: ", x)
}
