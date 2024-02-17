package main

import (
	"fmt"
	game_area "go-sos-game/game-area"
	users "go-sos-game/users"
	"math/rand"
	"os"
	"time"
)

var gameAreaSize *int
var playerScore int
var computerScore int
var numberOfMatrixElements int
var playerOrder bool
var computerOrder bool
var sosGameMatrix [][]string
var playerEnterCharacter string
var computerPlayerCharacter string
var pIndexX *int
var pIndexY *int
var cIndexX *int
var cIndexY *int
var indexState bool
var gameMatrix [][]string
var score int

func main() {
	fmt.Println("*********************Sos-Game*********************")
	time.Sleep(5000)
	gameAreaSize = game_area.GetGameAreaDimension()
	sosGameMatrix = game_area.SetDefaultValue(make([][]string, *gameAreaSize))
	game_area.DrawGameArea(sosGameMatrix)

	playerOrder, computerOrder = users.DecideToTheStartGame()
	fmt.Println(playerOrder, computerOrder)
	playerEnterCharacter, computerPlayerCharacter = users.DecideToTheCharacter()
	fmt.Printf("Player will start with : %s letter to the game.\nComputer will start with : %s letter to the game\n", playerEnterCharacter, computerPlayerCharacter)
	game_area.ScoreBoard(playerScore, computerScore)
	StartGame()
}

func NumberOfMatrixElementsControl(numberOfMatrixElements int) {
	if numberOfMatrixElements <= 0 {
		fmt.Println("There are no characters left to enter the panel, the game ends")
		WhoIsTheGameWin()
		os.Exit(0)
	}
}

func WhoIsTheGameWin() {
	fmt.Println("--------------------------------------------")
	if playerScore < computerScore {
		fmt.Println("The Computer won game")
	} else if playerScore > computerScore {
		fmt.Println("The Player won game")
	} else {
		fmt.Println("The game ended in a draw")
	}
	fmt.Println("--------------------------------------------")
	game_area.ScoreBoard(playerScore, computerScore)
	fmt.Println("--------------------------------------------")
	fmt.Println("----------------------------The Game Over----------------------------------")
}

func IsValidIndexXAndIndexY(indexX int, indexY int) bool {
	if indexX < 0 || indexX >= *gameAreaSize {
		if indexY < 0 || indexY >= *gameAreaSize {
			fmt.Println("Re-enter indexX and indexY values")
			return false
		} else {
			fmt.Println("Re enter indexY values")
			return false
		}
	} else {
		if indexY < 0 || indexY >= *gameAreaSize {
			fmt.Println("Re-enter indexX values")
			return false
		} else {
			return true
		}
	}
}

func PlayerEnterIndexValues() {
	var playerIndexX int
	var playerIndexY int
	for {
		fmt.Printf("Enter indexX value (between 0 and %d):", *gameAreaSize-1)
		_, _ = fmt.Scanln(&playerIndexX)
		fmt.Printf("Enter indexY value (between 0 and %d):", *gameAreaSize-1)
		_, _ = fmt.Scanln(&playerIndexY)
		result := IsValidIndexXAndIndexY(playerIndexX, playerIndexY)
		if result {
			pIndexX = &playerIndexX
			pIndexY = &playerIndexY
			break
		}

	}
}

func ComputerEnterIndexValues() {
	var computerIndexX = rand.Intn(*gameAreaSize)
	var computerIndexY = rand.Intn(*gameAreaSize)
	fmt.Println("ComputerIndexX : ", computerIndexX, "CompTuerIndexY : ", computerIndexY)
	cIndexX = &computerIndexX
	cIndexY = &computerIndexY
}

func IsMatrixIndexNull(indexX *int, indexY *int) bool {
	fmt.Println("IndexX  : ", *indexX, "IndexY : ", *indexY)
	if gameMatrix[*indexX][*indexY] == "None" {
		return true
	} else if gameMatrix[*indexX][*indexY] == "S" {
		fmt.Println("Re-enter Index value")
		return false
	} else {
		fmt.Println("Re-enter Index value")
		return false
	}
}

func EnterCharacterToPanel(indexX *int, indexY *int, character *string, playerOrder bool, computerOrder bool) {
	fmt.Println("Character : ", *character)
	gameMatrix[*indexX][*indexY] = *character
	fmt.Println("Value : ", gameMatrix[*indexX][*indexY])
	if gameMatrix[*indexX][*indexY] == "S" && *character == "S" && playerOrder {
		PlayerOrderAndSLetterControl(indexX, indexY)
		score = 0
	} else if gameMatrix[*indexX][*indexY] == "S" && *character == "S" && computerOrder {
		ComputerOrderAndSLetterControl(indexX, indexY)
		score = 0
	} else if gameMatrix[*indexX][*indexY] == "O" && *character == "O" && playerOrder {
		PlayerOrderAndOLetterControl(indexX, indexY)
		score = 0
	} else if gameMatrix[*indexX][*indexY] == "O" && *character == "O" && computerOrder {
		ComputerOrderAndOLetterControl(indexX, indexY)
		score = 0
	}

}

func CharacterSControl(indexX *int, indexY *int) int {
	fmt.Println("Selected Value : ", gameMatrix[*indexX][*indexY])
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if gameMatrix[*indexX][*indexY] == "S" {
				if *indexX+2*i >= 0 && *indexX+2*i < len(gameMatrix) && *indexY+2*j >= 0 && *indexY+2*j < len(gameMatrix) {
					if gameMatrix[*indexX+i][*indexY+j] == "O" && gameMatrix[*indexX+2*i][*indexY+2*j] == "S" {
						score++
					}
				}
			}

		}
	}
	fmt.Println("Score : ", score)
	return score
}

func CharacterOControl(indexX *int, indexY *int) int {
	fmt.Println("Selected Value : ", gameMatrix[*indexX][*indexY])
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if gameMatrix[*indexX][*indexY] == "O" {
				if (*indexX+i >= 0 && *indexX+i < len(gameMatrix) && *indexY+j >= 0 && *indexY+j < len(gameMatrix)) && (*indexX-i >= 0 && *indexX-i < len(gameMatrix) && *indexY-j >= 0 && *indexY-j < len(gameMatrix)) {
					if gameMatrix[*indexX+i][*indexY+j] == "S" && gameMatrix[*indexX-i][*indexY-j] == "S" {
						score++
					}
				}
			}
		}
	}
	return score
}

func PlayerOrderAndSLetterControl(indexX *int, indexY *int) {
	sosCount := CharacterSControl(indexX, indexY)
	fmt.Println("SosCount : ", sosCount)
	if sosCount > 0 {
		fmt.Println("Player won ", sosCount, " point.Thus, the Player won one more game")
		playerScore += score
	} else {
		playerOrder = false
		computerOrder = true
	}
}

func PlayerOrderAndOLetterControl(indexX *int, indexY *int) {
	sosCount := CharacterOControl(indexX, indexY)
	if sosCount > 0 && sosCount <= 2 {
		fmt.Println("Player won 1 point.Thus, The Player  one more game")
		playerScore++
	} else if sosCount > 2 {
		fmt.Println("Player won ,", sosCount-2, " point.Thus, The player one more game")
		playerScore += 2
	} else if sosCount == 0 {
		playerOrder = false
		computerOrder = true
	}
}

func ComputerOrderAndSLetterControl(indexX *int, indexY *int) {
	sosCount := CharacterSControl(indexX, indexY)
	if sosCount > 0 {
		fmt.Println("Computer won ", sosCount, " point.Thus, the Player won one more game")
		computerScore += score
	} else {
		playerOrder = true
		computerOrder = false
	}
}

func ComputerOrderAndOLetterControl(indexX *int, indexY *int) {
	sosCount := CharacterOControl(indexX, indexY)
	if sosCount > 0 && sosCount <= 2 {
		fmt.Println("Computer won 1 point.Thus, The Player  one more game")
		computerScore++
	} else if sosCount > 2 {
		fmt.Println("Computer won ,", sosCount-2, " point.Thus, The player one more game")
		computerScore += 2
	} else if sosCount == 0 {
		playerOrder = true
		computerOrder = false
	}
}

func PlayerContinueGame(numberOfMatrixElements *int, sosGameMatrix [][]string, pEnterCharacter *string, playerOrder bool) {
	fmt.Println("NumberOfMatrixElements : in Player ", *numberOfMatrixElements, " Address : ", &numberOfMatrixElements)
	*numberOfMatrixElements--
	gameMatrix = sosGameMatrix
	fmt.Println("Player Index State : ", indexState)
	for !indexState {
		PlayerEnterIndexValues()
		indexState = IsMatrixIndexNull(pIndexX, pIndexY)
	}
	indexState = false
	EnterCharacterToPanel(pIndexX, pIndexY, pEnterCharacter, playerOrder, computerOrder)
	game_area.DrawGameArea(gameMatrix)
	game_area.ScoreBoard(playerScore, computerScore)
	NumberOfMatrixElementsControl(*numberOfMatrixElements)

}

func ComputerContinueGame(numberOfMatrixElements *int, sosGameMatrix [][]string, cEnterCharacter *string, computerOrder bool) {
	fmt.Println("NumberOfMatrixElements : in Player ", *numberOfMatrixElements, " Address : ", &numberOfMatrixElements)
	*numberOfMatrixElements--
	gameMatrix = sosGameMatrix
	for !indexState {
		ComputerEnterIndexValues()
		indexState = IsMatrixIndexNull(cIndexX, cIndexY)
	}
	indexState = false
	EnterCharacterToPanel(cIndexX, cIndexY, cEnterCharacter, playerOrder, computerOrder)
	game_area.DrawGameArea(gameMatrix)
	game_area.ScoreBoard(playerScore, computerScore)
	NumberOfMatrixElementsControl(*numberOfMatrixElements)
}

func Game() {
	var gameMatrix = &sosGameMatrix
	var noMatrixElements = &numberOfMatrixElements
	if playerOrder {
		PlayerContinueGame(noMatrixElements, *gameMatrix, &playerEnterCharacter, playerOrder)
	} else {
		ComputerContinueGame(noMatrixElements, *gameMatrix, &computerPlayerCharacter, computerOrder)
	}
}

func StartGame() {
	numberOfMatrixElements = *gameAreaSize * *gameAreaSize
	fmt.Println("NumberOfMatrixElements : ", numberOfMatrixElements, "  Address ", &numberOfMatrixElements)
	for {
		fmt.Println("Finish Game. Press Enter Y\nContinue Game. Press Enter N\nIf is not left null field on the game board. Press Enter: C")
		var decision string
		_, err := fmt.Scanln(&decision)
		if err != nil {
			panic("Unexpected error occurred")
		}
		switch decision {
		case "Y":
			fmt.Println("The Game has been exited.")
			return
		case "N":
			fmt.Println("The game will continue")
			Game()
			//numberOfMatrixElements++
		case "C":
			NumberOfMatrixElementsControl(numberOfMatrixElements)
		}
	}
}
