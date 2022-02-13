package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := name
	fmt.Println("Hello " + filter(nameFiltered))
}

// Accepting multiple function parameters
func sayHelloWithFilters(name string, filters ...Filter) {
	nameFiltered := name

	for _, filter := range filters {
		nameFiltered = filter(nameFiltered)
	}

	fmt.Println("Hello " + nameFiltered)
}

func spamFilter(name string) string {
	if name == "#@!#*%" {
		return "..."
	}

	return name
}

func uppercaseFilter(name string) string {
	return strings.ToUpper(name)
}

func main() {
	filter1 := spamFilter
	sayHelloWithFilters("Eko", filter1)
	sayHelloWithFilters("#@!#*%", filter1)

	filter2 := uppercaseFilter
	sayHelloWithFilters("Eko", filter1, filter2)
}
