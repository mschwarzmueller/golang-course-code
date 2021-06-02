package main

import "fmt"

func main() {
	age := 32

	fmt.Println(age)

	myAge := &age
	// var myAge *int
	// myAge = &age

	fmt.Println(*myAge)

	*myAge = 33

	fmt.Println(*myAge)
	fmt.Println(age)
}
