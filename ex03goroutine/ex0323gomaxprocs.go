package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("num of cpus:", runtime.NumCPU())

	fmt.Println("num of threads (before change):", runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(20)
	fmt.Println("num of threads (after change):", runtime.GOMAXPROCS(-1))

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			wg.Done()
		}()
	}

	fmt.Println("num of goroutines:", runtime.NumGoroutine())

	wg.Wait()
}
