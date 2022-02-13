package main

import (
	"fmt"
	"regexp"
)

func main() {
	// For details, check: https://github.com/google/re2/wiki/Syntax

	pattern := "[Ee][A-Za-z][Oo]"
	regex := regexp.MustCompile(pattern)

	fmt.Println("regex =", regex)
	fmt.Println("regex.MatchString(\"eko\") =", regex.MatchString("eko"))
	fmt.Println("regex.MatchString(\"eto\") =", regex.MatchString("eto"))
	fmt.Println("regex.MatchString(\"eDo\") =", regex.MatchString("eDo"))
	fmt.Println("regex.MatchString(\"EyO\") =", regex.MatchString("EyO"))
	fmt.Println("regex.MatchString(\"EKo\") =", regex.MatchString("EKo"))

	results := regex.FindAllString("eko eto xyZ eDo AbC  EyO EKo", 1)
	fmt.Println("results =", results)

	results = regex.FindAllString("eko eto xyZ eDo AbC  EyO EKo", 3)
	fmt.Println("results =", results)

	results = regex.FindAllString("eko eto xyZ eDo AbC  EyO EKo", -1)
	fmt.Println("results =", results)
}
