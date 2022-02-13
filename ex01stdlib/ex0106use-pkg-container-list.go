package main

import (
	"container/list"
	"fmt"
	"math/rand"
)

func main() {
	numbers := list.New()
	numbers.PushBack(rand.Intn(100))
	numbers.PushBack(rand.Intn(100))
	numbers.PushBack(rand.Intn(100))
	numbers.PushBack(rand.Intn(100))
	numbers.PushBack(rand.Intn(100))
	numbers.PushBack(rand.Intn(100))

	fmt.Println("numbers =", numbers)
	fmt.Println("Print numbers in order...")
	for i, e := 0, numbers.Front(); e != nil; i, e = i+1, e.Next() {
		fmt.Println("element ", i, "of numbers = ", e.Value)
	}
	fmt.Println("Print numbers in reversed order...")
	for i, e := numbers.Len(), numbers.Back(); e != nil; i, e = i-1, e.Prev() {
		fmt.Println("element ", i, "of numbers = ", e.Value)
	}

	names := list.New()
	names.PushBack("Eko")
	names.PushBack("Budi")
	names.PushBack("Joko")
	fmt.Println("names =", names)
	e := names.Front()
	fmt.Println("e := names.Front() =", e.Value)
	e = e.Next()
	fmt.Println("e = e.Next() =", e.Value)
	e = e.Next()
	fmt.Println("e = e.Next() =", e.Value)
	e = e.Prev().Prev()
	fmt.Println("e = e.Prev().Prev()", e.Value)
}
