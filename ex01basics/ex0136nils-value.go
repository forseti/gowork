package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}

	return map[string]string{
		"name": name,
	}
}

func showPerson(person map[string]string) {
	if person == nil {
		fmt.Println("Empty data")
	} else {
		fmt.Println(person)
	}
}

func main() {
	var person map[string]string
	showPerson(person)
	person = NewMap("")
	showPerson(person)
	person = NewMap("Eko")
	showPerson(person)
	person = nil
	showPerson(person)
}
