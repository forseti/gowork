package main

import (
	"context"
	"fmt"
	"github.com/forseti/gowork/ex04context/helper"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Total Goroutine (before):", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	dest := helper.CreateCounter(ctx, 0)
	fmt.Println("Total Goroutine (after CreateCounter(ctx)):", runtime.NumGoroutine())

	for n := range dest {
		fmt.Println("Counter:", n)
		if n >= 10 {
			break
		}
	}

	cancel()

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine (after cancel()):", runtime.NumGoroutine())
}
