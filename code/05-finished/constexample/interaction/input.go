package interaction

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Constants are used here to store the different player choice (1, 2 & 3)
// in a central place - this could avoid errors (typos) as these values don't
// have to be repeated in various places throughout the program
const (
	PlayerChoiceAttack        = 1
	PlayerChoiceHeal          = 2
	PlayerChoiceSpecialAttack = 3
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAttackIsAvailable bool) string {
	for {
		playerChoice, _ := getPlayerInput()

		if playerChoice == fmt.Sprint(PlayerChoiceAttack) {
			return "ATTACK"
		} else if playerChoice == fmt.Sprint(PlayerChoiceHeal) {
			return "HEAL"
		} else if playerChoice == fmt.Sprint(PlayerChoiceSpecialAttack) && specialAttackIsAvailable {
			return "SPECIAL_ATTACK"
		}
		fmt.Println("Fetching the user input failed. Please try again.")
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")

	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	if runtime.GOOS == "windows" {
		userInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		userInput = strings.Replace(userInput, "\n", "", -1)
	}

	return userInput, nil
}
