package main

import (
	"github.com/mschwarzmueller/monsterslayer/interaction"
)

func main() {
	startGame()

	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {}

func endGame() {}
