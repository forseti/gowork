package main

import (
	"fmt"
	"time"
)

func main() {
	// Like ticker but with a channel instead of a struct.
	tick := time.Tick(1 * time.Second)
	done := make(chan bool, 1)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for {
		select {
		case t := <-tick:
			fmt.Println(t)
		case <-done:
			return
		}
	}
}
