package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	}

	return "Hello " + name
}

func main() {
	// Returning a value
	hello := getHello("Eko")
	fmt.Println("getHello(\"Eko\") =", hello)
	fmt.Println("getHello(\"\") =", getHello(""))
}
