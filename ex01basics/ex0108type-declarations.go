package main

import "fmt"

func main() {
	type IdNum string
	type Married bool

	var idNum IdNum = "18741982741897419874"
	var marriedStatus Married = true
	fmt.Println("idNum =", idNum)
	fmt.Println("marriedStatus =", marriedStatus)
}
