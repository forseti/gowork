package main

import (
	"fmt"
	"github.com/forseti/gowork/ex03goroutine/helper"
	"strconv"
	"sync"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)

	data.Store(value, value)
}

func main() {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	helper.AsyncLoop(100, func(index int) {
		AddToMap(data, index, group)
	})

	group.Wait()

	i := 0
	data.Range(func(key, value interface{}) bool {
		fmt.Println("item #"+strconv.Itoa(i), "=", value)
		i++
		return true
	})

	fmt.Println("total items in map = ", i)
}
