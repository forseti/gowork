package main

import "fmt"

func main() {
	var name string

	name = "Eko Kurniawan"
	fmt.Println("name =", name)

	name = "Eko Khannedy"
	fmt.Println("name =", name)

	var friendName = "Budi"
	fmt.Println("friendName =", friendName)

	var age = 30
	fmt.Println("age =", age)

	country := "Indonesia"
	fmt.Println("country =", country)

	var (
		firstName = "Eko Kurniawan"
		lastName  = "Khannedy"
	)

	fmt.Println("firstName =", firstName)
	fmt.Println("lastName =", lastName)
}
