package main

import "fmt"

type Apapun interface{}

func Oops(i int) Apapun {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Oops"
	}
}

func main() {
	var data = Oops(1)
	fmt.Println("Oops(1) =", data)
	data = Oops(2)
	fmt.Println("Oops(2) =", data)
	data = Oops(3)
	fmt.Println("Oops(3) =", data)
}
