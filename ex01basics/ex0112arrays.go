package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khannedy"

	// Below will throw errors
	// names[3] = "ok"
	fmt.Println("names[0] =", names[0])
	fmt.Println("names[1] =", names[1])
	fmt.Println("names[2] =", names[2])

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println("values =", values)
	fmt.Println("values[0] =", values[0])
	fmt.Println("values[1] =", values[1])
	fmt.Println("values[2] =", values[2])

	fmt.Println("len(names) =", len(names))
	fmt.Println("len(values) =", len(values))
	
	var strArray [10]string
	fmt.Println("len(strArray)", len(strArray))
	fmt.Println("strArray", strArray)
}
