package scopes

import "fmt"

var userName = "Max"

func main() {
	shouldContinue := true
	userName := "Maximilian" // shadow the higher-level variable

	if shouldContinue {
		userAge := 32

		fmt.Printf("Name: %v, Age: %v", userName, userAge)
	}

	// fmt.Println(userAge) => Doesn't work
}

func printData() {
	fmt.Println(userName)
	// fmt.Println(shouldContinue) => Also doesn't work
}
