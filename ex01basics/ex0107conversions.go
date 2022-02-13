package main

import "fmt"

func main() {
	var num32 int32 = 100000
	var num64 = int64(num32)
	var num8 = int8(num64) // Can't do a correct conversion - integer overflow

	fmt.Println("num64 =", num64)
	fmt.Println("num32 =", num32)
	fmt.Println("num8 =", num8)

	var name = "Eko"
	var char = name[0]            // get a byte (or char) of string
	var charString = string(char) // Convert the byte (or the char) to string

	fmt.Println("name =", name)
	fmt.Println("char[0] =", charString)
}
