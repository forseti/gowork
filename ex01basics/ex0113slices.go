package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println("slice1 =", slice1)
	fmt.Println("len(slice1) =", len(slice1))
	fmt.Println("cap(slice1) =", cap(slice1))

	// Modifying values of the slice/array will change the values of original array.
	//months[5] = "Changed"
	//fmt.Println(slice1)
	//
	//slice1[0] = "May again"
	//fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println("slice2 =", slice2)

	var slice3 = append(slice2, "Eko")
	fmt.Println("slice3 =", slice3)

	slice3[1] = "Bukan Desember"
	fmt.Println("slice3 =", slice3)

	// The reason why modifying slice3 didn't affect slice2 and months
	// because append(slice2, "Eko") created new array since slice2 have not enough capacity
	// to accommodate "Eko" at the end of the array.
	fmt.Println("slice2 =", slice2)
	fmt.Println("months =", months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"
	fmt.Println("newSlice =", newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("copySlice =", copySlice)

	thisIsArray := [5]int{1, 2, 3, 4, 5}
	thisIsSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("thisIsArray =", thisIsArray)
	fmt.Println("thisIsSlice =", thisIsSlice)
}
