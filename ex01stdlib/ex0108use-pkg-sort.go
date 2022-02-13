package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type Users []User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := &Users{
		User{"Budi", 35},
		User{"Eko", 30},
		User{"Joko", 22},
		User{"Sigit", 37},
		User{"Rudi", 27},
	}

	fmt.Println("Before sorting, users =", *users)
	sort.Sort(users)
	fmt.Println("After sorting, users =", *users)
}
