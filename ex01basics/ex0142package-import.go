package main

import (
	"fmt"
	"github.com/forseti/gowork/ex01basics/helper"
)

func main() {
	helper.SayHello("Eko")
	// ex01helper.sayGoodbye("Eko") // error
	fmt.Println("ex01helper.Application =", helper.Application)
	// fmt.Println(ex01helper.version) // error
}
