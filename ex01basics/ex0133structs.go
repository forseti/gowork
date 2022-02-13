package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	//Married       bool
}

func (customer Customer) sayHelloTo(otherCustomer Customer) {
	fmt.Println("Hello", otherCustomer.Name+",", "My Name is", customer.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko Kurniawan"
	eko.Address = "Indonesia"
	eko.Age = 30

	fmt.Println("eko =", eko)
	fmt.Println("eko.Name =", eko.Name)
	fmt.Println("eko.Address =", eko.Address)
	fmt.Println("eko.Age =", eko.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Cirebon",
		Age:     35,
	}
	fmt.Println("joko =", joko)

	budi := Customer{"Budi", "Jakarta", 35}
	fmt.Println("budi =", budi)

	rully := Customer{Name: "Rully"}
	rully.sayHelloTo(eko)
}
