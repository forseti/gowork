package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		group.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	time.Sleep(5 * time.Second)
	fmt.Println("x: ", x)
}
