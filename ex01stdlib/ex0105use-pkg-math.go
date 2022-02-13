package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("math.Round(1.7)", math.Round(1.7))
	fmt.Println("math.Round(1.3)", math.Round(1.3))
	fmt.Println("math.Floor(1.7)", math.Floor(1.7))
	fmt.Println("math.Ceil(1.3)", math.Ceil(1.3))
	fmt.Println("math.Max(10, 20)", math.Max(10, 20))
	fmt.Println("math.Min(10, 20)", math.Min(10, 20))
}
