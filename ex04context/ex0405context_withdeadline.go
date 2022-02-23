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
	// context.WithTimeout(parent, duration) is actually context.WithDeadline(parent, time.Now().Add(duration))
	// However, context.WithDeadline(parent, duration) is more flexible as you are not tied to time.Now()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	dest := helper.CreateCounter(ctx, 1)
	fmt.Println("Total Goroutine (after CreateCounter(ctx)):", runtime.NumGoroutine())

	for n := range dest {
		fmt.Println("Counter:", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine (after cancel()):", runtime.NumGoroutine())
}
