package main

import (
	"fmt"
	"github.com/forseti/gowork/ex03goroutine/helper"
	"sync"
	"time"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	pool.Put("Eko")
	pool.Put("Kurniawan")
	pool.Put("Khannedy")

	helper.AsyncLoop(10, func(index int) {
		data := pool.Get()
		fmt.Println(data)
		time.Sleep(1 * time.Second)
		pool.Put(data)
	})

	time.Sleep(1 * time.Second)

	fmt.Println("Done")
}
