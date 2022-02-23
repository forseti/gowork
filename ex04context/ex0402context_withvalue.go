package main

import (
	"context"
	"fmt"
)

func main() {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println("contextA", contextA)
	fmt.Println("contextB", contextB)
	fmt.Println("contextC", contextC)
	fmt.Println("contextD", contextD)
	fmt.Println("contextE", contextE)
	fmt.Println("contextF", contextF)
	fmt.Println("contextG", contextG)

	fmt.Println("Does contextF have value with key 'f'?", contextF.Value("f"))
	fmt.Println("Does contextF have value with key 'c'?", contextF.Value("c"))
	fmt.Println("Does contextF have value with key 'b'?", contextF.Value("b"))
	fmt.Println("Does contextA have value with key 'b'?", contextA.Value("b"))
}
