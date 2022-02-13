package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Learning Go-Lang"
	book["author"] = "Eko"
	book["oops"] = "Wrong"

	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "oops")
	fmt.Println(book)
	fmt.Println(len(book))
}
