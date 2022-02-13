package helper

import "fmt"

var version = 1.0
var Application = "Learning Go"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}
