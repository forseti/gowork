package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(10, 90, 30, 50, 40)
	fmt.Println("total =", total)
	fmt.Println("sumAll() =", sumAll())
	slice := []int{10, 20, 30, 40, 50}
	total = sumAll(slice...)
	fmt.Println("sumAll(slice...) =", total)
}
