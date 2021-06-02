package main

import "fmt"

func main() {
	// var greetingText string
	// greetingText = "Hello World!!!!!"

	// var greetingText string = "Hello World!!!!!"

	greetingText := "Hello World!!!!!"
	luckyNumber := 17
	evenMoreLuckyNumber := luckyNumber + 5

	fmt.Println(greetingText)
	fmt.Println(greetingText)
	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)

	var newNumber float64 = float64(luckyNumber) / 3

	fmt.Println(newNumber)

	var defaultFloat float64 = 1.23456789123456789123456
	var smallFloat float32 = 1.23456789123456789123456

	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)

	var firstRune rune = '€'
	fmt.Println(firstRune)
	fmt.Println(string(firstRune))

	var firstByte byte = 'a'
	fmt.Println(firstByte)

	firstName := "Max"
	lastName := "Schwarzmüller"
	fullName := fmt.Sprintf("%v %v", firstName, lastName)
	age := 32

	fmt.Printf("Hi, I am %v and I am %v (Type: %T) years old.", fullName, age, age)

	// fmt.Println("9" + 1)
}
