package main

import (
	"fmt"
	"github.com/forseti/gowork/ex03goroutine/helper"
	"time"
)

func main() {
	go helper.HelloName("World")
	fmt.Println("Oops")
	time.Sleep(1 * time.Second)
}
