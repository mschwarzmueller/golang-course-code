package main

import "fmt"

func main() {
	firstName := "Maximilian"
	var lastName string = "Schwarzmüller"

	fmt.Println(firstName)
	fmt.Println(lastName)

	currentYear := 2021
	var birthYear int
	birthYear = 1989

	age := currentYear - birthYear

	fmt.Println(age)

	currentYear = currentYear + 1

	age = currentYear - birthYear

	fmt.Println(age)
}
