package main

import "fmt"

func main() {
	a := 5
	b := 10

	sum := add(a, b)

	fmt.Println(sum)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}
