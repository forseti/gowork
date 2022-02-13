package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}

	fmt.Println("address1 := Address{\"Subang\", \"Jawa Barat\", \"Indonesia\"}")
	fmt.Println("\nExample #1 - Changing the value of address1 (Invalid)")
	fmt.Println("address2 := address1")
	fmt.Println("address2.City = \"Bandung\"")
	address2 := address1
	address2.City = "Bandung"
	fmt.Println("Address 1", address1) // Remains unchanged
	fmt.Println("Address 2", address2)

	fmt.Println("\nExample #2 - Changing the value of the address1 (Successful)")
	fmt.Println("address3 := &address1")
	fmt.Println("address3.City = \"Bandung\"")
	address3 := &address1 // same like: var address3 *Address = &Address1
	address3.City = "Bandung"
	fmt.Println("Address 1", address1) // Changed (note: only address value changed)
	fmt.Println("Address 3", address3)

	fmt.Println("\nExample #3 - Initialize address1 as a new object (Invalid)")
	fmt.Println("address4 := &address1")
	fmt.Println("address3 = &Address{\"Malang\", \"Jawa Timur\", \"Indonesia\"}")
	address4 := &address1
	address3 = &Address{"Malang", "Jawa Timur", "Indonesia"}
	fmt.Println("Address 1", address1) // Still refers to original reference
	fmt.Println("Address 3", address3)
	fmt.Println("Address 4", address4)

	fmt.Println("\nExample #4 - Initialize address1 as a new object (Successful)")
	fmt.Println("address5 := &address1")
	fmt.Println("*address5 = Address{\"Surabaya\", \"Jawa Timur\", \"Indonesia\"}")
	address5 := &address1
	*address5 = Address{"Surabaya", "Jawa Timur", "Indonesia"}
	fmt.Println("Address 1", address1) // Changed to new reference
	fmt.Println("Address 4", address4)
	fmt.Println("Address 5", address5)

	fmt.Println("\nExample #5 - Using 'new'")
	fmt.Println("var address6 = new(Address)")
	var address6 = new(Address)
	fmt.Println("Address 6", address6)
	fmt.Println("address6.City = \"Jakarta\"")
	address6.City = "Jakarta"
	fmt.Println("Address 6", address6)

	fmt.Println("\nExample #7 - function - pass by reference")
	fmt.Println(
		`address7 := Address{
			City:     "Subang",
			Province: "Jawa Barat",
			Country:  "",
		}`,
	)
	address7 := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	fmt.Println("Address 7", address7)
	fmt.Println("ChangeCountryToIndonesia(address7)")
	ChangeCountryToIndonesia(&address7)
	fmt.Println("Address 7", address7)
}
