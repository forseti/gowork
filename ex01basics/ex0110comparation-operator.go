package main

import "fmt"

func main() {
	var (
		name1 = "Eko"
		name2 = "Budi"
	)

	var result bool = name1 == name2
	fmt.Println("name 1 == name2 =", result)

	var (
		value1 = 100
		value2 = 200
	)

	fmt.Println("value1 > value2 =", value1 > value2)
	fmt.Println("value1 < value2 =", value1 < value2)
	fmt.Println("value1 == value2 =", value1 == value2)
	fmt.Println("value1 != value2 =", value1 != value2)
}
