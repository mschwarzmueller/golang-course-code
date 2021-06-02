package main

import "fmt"

func main() {
	age := 32

	fmt.Println(age)

	myAge := &age
	// var myAge *int
	// myAge = &age

	fmt.Println(myAge)
}
