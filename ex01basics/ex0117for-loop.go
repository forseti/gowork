package main

import "fmt"

func main() {
	//counter := 1
	//
	//for counter <= 10 {
	//	fmt.Println("Iteration", counter)
	//	counter++
	//}

	// With stmts
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Iteration", counter)
	}

	// loop the slice
	names := []string{"Eko", "Kurniawan", "Khannedy"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// With range
	for i, name := range names {
		fmt.Println("Index", i, "=", name)
	}

	// With range
	for _, name := range names {
		fmt.Println(name)
	}

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"
	for key, value := range person {
		fmt.Println("Key", key, "=", value)
	}
}
