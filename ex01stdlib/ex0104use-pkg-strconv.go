package main

import (
	"fmt"
	"strconv"
)

func main() {
	status, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("strconv.ParseBool(\"true\") =", status)
	} else {
		fmt.Println("err =", err)
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println("strconv.ParseInt(\"1000000\", 10, 64) =", number)
	} else {
		fmt.Println("err =", err)
	}

	for _, base := range []int{2, 8, 10, 16} {
		value := strconv.FormatInt(number, base)
		fmt.Println("base =", strconv.FormatInt(int64(base), 10)+", value =", value)
	}
}
