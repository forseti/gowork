package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	// Using time.NewTimer, where you need to get the channel from the timer object.
	log.Println("Using time.NewTimer")
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("before 5s:", time.Now())

	t := <-timer.C
	fmt.Println("after 5s:", t)

	// Using time.After which return a channel.
	log.Println("Using time.After")
	timeChan := time.After(5 * time.Second)
	fmt.Println("before 5s", time.Now())

	t = <-timeChan
	fmt.Println("after 5s", t)

	// Using time.AfterFunc
	log.Println("Using time.AfterFunc")
	wg := sync.WaitGroup{}
	wg.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("after 5s", time.Now()) // After 5 secs
		wg.Done()
	})
	fmt.Println("before 5s", time.Now()) // Before 5 secs
	wg.Wait()
}
