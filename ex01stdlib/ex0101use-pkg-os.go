package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("args =", args)

	switch len(args) {
	case 1:
		fmt.Println("No args provided")
	case 2:
		fmt.Println("args[1] =", args[1])
	case 3:
		fmt.Println("args[1] =", args[1])
		fmt.Println("args[2] =", args[2])
	case 4:
		fmt.Println("args[1] =", args[1])
		fmt.Println("args[2] =", args[2])
		fmt.Println("args[3] =", args[3])
	default:
		fmt.Println("Too many args")
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("hostname =", hostname)
	} else {
		fmt.Println("err = ", err.Error())
	}

	gopath := os.Getenv("GOPATH")
	username := os.Getenv("USERNAME")
	println("gopath =", gopath)
	println("username =", username)
}
