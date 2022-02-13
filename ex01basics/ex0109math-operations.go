package main

import "fmt"

func main() {

	var result = 10 + 10
	fmt.Println("result =", result)

	var (
		a = 10
		b = 10
		c = a * b
	)
	fmt.Println("c =", c)

	var i = 10
	i += 10
	fmt.Println("i =", i)

	i++
	fmt.Println("i =", i)

	var negative = -100
	var positive = +100
	fmt.Println("negative =", negative)
	fmt.Println("positive =", positive)
}
