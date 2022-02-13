package main

import "fmt"

func main() {
	name := "Eko"
	//name := "Joko"
	//name := "Budi"
	//name := "Kurniawan"
	//name := "KurniawanKurniawan"

	// General
	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello Joko")
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Let's get acquainted")
		fmt.Println("Let's get acquainted")
	}

	// With short stmts
	/*	switch length := len(name); length > 5 {
		case true:
			fmt.Println("Name is too long")
		case false:
			fmt.Println("Name's length is correct")
		}*/

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Name is too long")
	case length > 5:
		fmt.Println("Name is quite long")
	default:
		fmt.Println("Name's length is correct")
	}
}
