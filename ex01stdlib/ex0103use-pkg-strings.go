package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Eko Kurniawan"
	fmt.Println("name =", name)
	fmt.Println("strings.Contains(name, \"Eko\") =", strings.Contains(name, "Eko"))
	fmt.Println("strings.Contains(name, \"Budi\") =", strings.Contains(name, "Budi"))

	name = "Eko Kurniawan Khannedy"
	fmt.Println("name =", name)
	fmt.Println("strings.Split(name, \" \") =", strings.Split(name, " "))

	fmt.Println("strings.ToLower(name) =", strings.ToLower(name))
	fmt.Println("strings.ToUpper(name) =", strings.ToUpper(name))

	name = "eko kurniawan khannedy"
	fmt.Println("name =", name)
	fmt.Println("strings.Title(name) =", strings.Title(name))

	name = " 	 Eko Kurniawan    			    "
	fmt.Println("name =", name+".")
	fmt.Println("strings.TrimSpace(name) =", strings.TrimSpace(name)+".")

	name = "Eko Eko Eko Eko Eko"
	fmt.Println("name =", name)
	fmt.Println("strings.ReplaceAll(name, \"Joko\", \"Budi\") =", strings.ReplaceAll(name, "Joko", "Budi"))
	fmt.Println("strings.ReplaceAll(name, \"Eko\", \"Budi\") =", strings.ReplaceAll(name, "Eko", "Budi"))
}
