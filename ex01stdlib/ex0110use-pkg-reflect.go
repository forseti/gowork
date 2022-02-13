package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	valid := true

	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tagRequired := field.Tag.Get("required")
		tagMax := field.Tag.Get("max")

		if valid && tagRequired != "" {
			fmt.Println("validate if data is not empty...")
			required, _ := strconv.ParseBool(field.Tag.Get("required"))
			valid = required && reflect.ValueOf(data).Field(0).String() != ""
		}

		if valid && tagMax != "" {
			fmt.Println("validate if data's length <= max...")
			max, _ := strconv.Atoi(field.Tag.Get("max"))
			length := len(reflect.ValueOf(data).Field(0).String())
			valid = !(length > max)
		}
	}

	return valid
}

func main() {
	sample := Sample{"Eko"}
	sampleType := reflect.TypeOf(sample)
	fmt.Println("sample =", sample)
	fmt.Println("sampleType = ", sampleType)
	fmt.Println("Sample, num of fields = ", sampleType.NumField())
	fmt.Println("Sample, name of field #0 = ", sampleType.Field(0).Name)
	fmt.Println("Sample, name of field #0, tag:required =", sampleType.Field(0).Tag.Get("required"))
	fmt.Println("Sample, name of field #0, tag:max =", sampleType.Field(0).Tag.Get("max"))
	fmt.Println("Sample, name of field #0, tag:min =", sampleType.Field(0).Tag.Get("min"))
	fmt.Println("IsValid(sample) =", IsValid(sample))

	sample = Sample{}
	fmt.Println("sample =", sample)
	fmt.Println("IsValid(sample) =", IsValid(sample))

	sample = Sample{"Maximillian"}
	fmt.Println("sample =", sample)
	fmt.Println("IsValid(sample) =", IsValid(sample))
}
