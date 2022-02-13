package main

import (
	"container/ring"
	"fmt"
)

func main() {
	data := ring.New(5)

	names := []string{"Eko", "Budi", "Joko", "Rully", "Sigit"}

	for i := 0; i < data.Len(); i++ {
		data.Value = names[i]
		data = data.Next()
	}

	i := 0
	data.Do(func(value interface{}) {
		fmt.Println("element ", i, " =", value)
		i++
	})
}
