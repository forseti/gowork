package main

import (
	"errors"
	"fmt"
)

func Divide(nilai int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("divisor cannot be 0")
	} else {
		result := nilai / divisor
		return result, nil
	}
}

func showResult(result int, myError error) {
	if myError == nil {
		fmt.Println("Result", result)
	} else {
		fmt.Println("Error", myError)
	}
}

func main() {
	var errorExample = errors.New("Oops error")
	fmt.Println(errorExample.Error())

	var myError error
	result, myError := Divide(100, 10)
	showResult(result, myError)

	result, myError = Divide(100, 0)
	showResult(result, myError)
}
