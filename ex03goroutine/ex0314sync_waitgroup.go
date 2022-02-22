package main

import (
	"fmt"
	"sync"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup, i int) {
	defer group.Done()
	group.Add(1)
	fmt.Println("Hello", i)
	time.Sleep(1 * time.Second)
}

func main() {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group, i)
	}

	group.Wait()

	fmt.Println("Done")
}
