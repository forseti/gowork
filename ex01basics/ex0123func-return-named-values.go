package main

import "fmt"

func getFullName2() (firstName, middleName, lastName string) {
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khannedy"

	return
}

func main() {
	a, b, c := getFullName2()
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
}
