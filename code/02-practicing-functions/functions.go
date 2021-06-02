package main

import "fmt"

func main() {
	a := 5
	b := 10

	sum := add(a, b)

	printNumber(sum)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func printNumber(number int) {
	fmt.Printf("The number is %v", number)
}
