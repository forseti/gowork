package main

import "fmt"

func getFullName() (string, string, string) {
	return "Eko", "Kurniawan", "Khannedy"
}

func main() {
	firstName, _ /*middleName*/, lastName := getFullName()
	fmt.Println("firstName =", firstName)
	//fmt.Println(middleName)
	fmt.Println("lastName =", lastName)
}
