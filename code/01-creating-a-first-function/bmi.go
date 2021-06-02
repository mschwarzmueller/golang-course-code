package main

import (
	"fmt"

	"github.com/mschwarzmueller/bmi/info"
)

func main() {
	// Output information
	info.PrintWelcome()

	weight, height := getUserMetrics()

	// Calculate the BMI (weight / (height * height))
	bmi := weight / (height * height)

	// Output the calculated BMI
	fmt.Printf("Your BMI: %.2f", bmi)
}
