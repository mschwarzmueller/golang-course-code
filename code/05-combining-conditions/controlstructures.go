package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your age: ")
	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
	userAge, _ := strconv.ParseInt(userAgeInput, 0, 64)

	// isOldEnough := userAge >= 18

	if (userAge >= 30 && userAge < 50) || userAge >= 60 {
		fmt.Println("You're eligible for our senior jobs.")
	} else if userAge >= 50 {
		fmt.Println("The best age!")
	} else if userAge >= 18 {
		fmt.Println("Welcome to the club!")
	} else {
		fmt.Println("Sorry, you're not old enough :/")
	}

	fmt.Println("Goodbye!")
}
