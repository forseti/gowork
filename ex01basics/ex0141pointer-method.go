package main

import "fmt"

type Man struct {
	Name    string
	married bool
}

func (man *Man) Married() {
	if man.married {
		man.Name = "Mr. " + man.Name
	}
}

func main() {
	eko := Man{"Eko", true}
	eko.Married()
	fmt.Println("Eko's name", eko.Name)

	budi := Man{"Budi", false}
	budi.Married()
	fmt.Println("Budi's name", budi.Name)
}
