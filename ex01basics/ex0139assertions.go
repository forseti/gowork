package main

import "fmt"

func getStringValue() interface{} {
	return "Oops"
}
func getIntValue() interface{} {
	return 7
}

func assertValue(result interface{}) (string, int) {
	var resultString string
	var resultInt int

	// Safe assertions (won't lead to panic)
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
		resultString = value
	case int:
		fmt.Println("Value", value, "is int")
		resultInt = value
	default:
		fmt.Println("Unknown type")
	}

	return resultString, resultInt
}

func main() {
	var result = getStringValue()
	var resultString string
	var resultInt int

	// Safe assertions
	result = getIntValue()
	resultString, resultInt = assertValue(result)
	fmt.Println("resultString =", resultString)
	fmt.Println("resultInt =", resultInt)

	result = getStringValue()
	resultString, resultInt = assertValue(result)
	fmt.Println("resultString =", resultString)
	fmt.Println("resultInt =", resultInt)

	result = []string{}
	resultString, resultInt = assertValue(result)
	fmt.Println("resultString =", resultString)
	fmt.Println("resultInt =", resultInt)

	// Unsafe assertions lead to panic
	resultString = result.(string)
	fmt.Println("resultString =", resultString)
	resultInt = result.(int)
	fmt.Println("resultInt =", resultInt)
}
