package main

import "fmt"

func main() {
	name := "Eko"
	//name := "Joko"
	//name := "Budi"
	//name := "Rully"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hi, let's get acquainted")
	}

	// With short stmts
	if length := len(name); length > 5 {
		fmt.Println("Name is too long")
	} else {
		fmt.Println("Nama's length is correct")
	}
}
