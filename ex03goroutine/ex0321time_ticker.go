package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool, 1)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		done <- true
	}()

	// Use select to avoid deadlock/empty channel
	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t)
		case <-done:
			return
		}
	}

	// Will result in empty channel/deadlock
	//for t := range ticker.C {
	//	fmt.Println(t)
	//}
}
