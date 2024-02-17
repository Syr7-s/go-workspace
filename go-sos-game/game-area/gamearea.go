package game_area

import (
	"fmt"
	"strconv"
)

func GetGameAreaDimension() *int {
	var d int
	var dimension *int
	var err error
	for {
		fmt.Println("***************************************************")
		fmt.Println("Please, Enter the number between 3 and 7 : ")
		var input string
		fmt.Scanln(&input)
		d, err = strconv.Atoi(input)
		dimension = &d
		if err != nil {
			fmt.Println("You must enter values")
			continue
		}
		fmt.Println("***************************************************")
		if *dimension < 3 || *dimension > 7 {
			fmt.Println("Please try again")
		} else {
			break
		}
	}
	fmt.Println("Game Area Dimension : ", dimension)
	return dimension
}

func SetDefaultValue(gameMatrix [][]string) [][]string {
	for i := range gameMatrix {
		gameMatrix[i] = make([]string, len(gameMatrix))
		for j := range gameMatrix[i] {
			gameMatrix[i][j] = "None"
		}
	}
	return gameMatrix
}

func DrawGameArea(gameMatrix [][]string) {
	fmt.Println("--------------------------------------------")
	fmt.Println("------------------SOS Panel-----------------")

	//i := 0
	for j := 0; j < len(gameMatrix); j++ {
		fmt.Print("	", j+1, " ")
	}
	fmt.Println()
	for i, values := range gameMatrix {
		fmt.Printf("%d ", i+1)
		for _, value := range values {
			fmt.Printf("| %s | ", value)
		}
		fmt.Println()
	}

	fmt.Println("--------------------------------------------")
}

func ScoreBoard(playerScore int, computerScore int) {
	fmt.Println("--------------------------------------------")
	fmt.Println("*****************Score Board***************\n*******************************************\n*Players:******************Score***********\n*******************************************\n*Player:   ****************| ", playerScore, " |*********\n*Computer: ****************| ", computerScore, " |*********\n*******************************************")
	fmt.Println("--------------------------------------------")

}
