package main

import "fmt"

func main() {
	//const firstName string = "Eko"
	//const lastName = "Khannedy"
	//const value = 1000

	const (
		firstName string = "Eko"
		lastName         = "Khannedy"
		value            = 1000
	)

	fmt.Println("firstName =", firstName)
	fmt.Println("lastName =", lastName)
	fmt.Println("value =", value)
}
