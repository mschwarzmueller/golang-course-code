package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your age: ")
	userAgeInput, _ := reader.ReadString('\n')
	if runtime.GOOS == "windows" {
		userAgeInput = strings.Replace(userAgeInput, "\r\n", "", -1)
	} else {
		userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
	}
	userAge, _ := strconv.ParseInt(userAgeInput, 0, 64)

	if userAge >= 18 {
		fmt.Println("Welcome to the club!")
	}
}
