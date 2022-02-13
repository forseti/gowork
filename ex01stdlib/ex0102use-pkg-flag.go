package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Database's hostname")
	port := flag.Uint("port", 5432, "Database's port")
	user := flag.String("user", "root", "Database's username")
	password := flag.String("password", "password", "Database's password")

	flag.Parse()
	//if flag.NFlag() == 0 {
	//	flag.Usage()
	//} else {
	fmt.Println("host =", *host)
	fmt.Println("port =", *port)
	fmt.Println("user =", *user)
	fmt.Println("password =", *password)
	//}
}
