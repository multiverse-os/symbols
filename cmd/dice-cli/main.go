package main

import (
	"fmt"
	"math/rand"
	"time"

	dice "github.com/multiverse-os/symbols/games/dice"
)

func randomDiceNumber() int {
	randomNumber := rand.Intn(6)
	fmt.Printf("randomNumber: %v", randomNumber)
	return randomNumber
}

func rollDice() string {
	switch randomDiceNumber() {
	case 1:
		return dice.Symbols["one"]
	case 2:
		return dice.Symbols["two"]
	case 3:
		return dice.Symbols["three"]
	case 4:
		return dice.Symbols["four"]
	case 5:
		return dice.Symbols["five"]
	default: // 6
		return dice.Symbols["six"]
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("dice-cli\n")
	fmt.Printf("========\n")

	fmt.Printf("=> %s + %s\n", rollDice(), rollDice())

}
