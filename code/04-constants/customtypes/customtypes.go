package customtypes

import "fmt"

type person struct {
	name string
	age  int
}

type customNumber int

type personData map[string]person

func (number customNumber) pow(power int) customNumber {
	result := number

	for i := 1; i < power; i++ {
		result = result * number
	}

	return result
}

func main() {
	var people personData = personData{
		"p1": {"Max", 32},
	}

	fmt.Println(people)

	var dummyNumber customNumber = 5
	poweredNumber := dummyNumber.pow(3)

	fmt.Println(poweredNumber)
}
