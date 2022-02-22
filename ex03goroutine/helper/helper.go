package helper

import (
	"fmt"
	"time"
)

func HelloName(name string) {
	fmt.Println("Hello", name+"!")
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Eko Kurniawan Khannedy"
	fmt.Println("Response has been send to channel")
}

func AsyncLoop(times int, doSomething func(index int)) {
	for i := 0; i < times; i++ {
		go doSomething(i)
	}
}
