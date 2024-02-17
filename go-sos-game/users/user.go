package users

import (
	"fmt"
	"math/rand"
	"time"
)

const LETTERS string = "S"
const LETERO string = "O"

func DecideToTheStartGame() (bool, bool) {
	rand.Seed(time.Now().UnixNano())
	isStart := rand.Intn(2) != 0
	if isStart {
		fmt.Println("Player will start the game")
		return true, false
	} else {
		fmt.Println("Computer will start the game")
		return false, true
	}
}

func DecideToTheCharacter() (string, string) {
	rand.Seed(time.Now().UnixNano())
	takeCh := rand.Intn(2) != 0
	if takeCh {
		return LETTERS, LETERO
	}
	return LETERO, LETTERS
}
