package main

import "fmt"

func endApp() {
	message := recover()

	if message != nil {
		fmt.Println("Error with message: ", message)
	}

	fmt.Println("Application ended")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Application running")
}

func main() {
	runApp(true)
	fmt.Println("Eko")
}
