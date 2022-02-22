package main

import "fmt"

func retNameInChan(name string) <-chan string {
	out := make(chan string, 3)
	defer close(out)
	out <- "Eko"
	out <- "Kurniawan"
	out <- "Khannedy"
	return out
}

func main() {
	data := retNameInChan("Eko Kurniawan")
	fmt.Println(<-data, <-data, <-data)
}
