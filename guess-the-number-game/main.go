package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	var randomNumber int = rand.Intn(3) + 1
	var chance int = 0
	for i := chance; chance <= 2; chance++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a Number : ")
		input, _ := reader.ReadString('\n')
		fmt.Println("You Entered : ", input)
		convertNumber, err := strconv.ParseInt(strings.TrimSpace(input), 0, 32)
		if err != nil {
			fmt.Println(err)

		}

		fmt.Println("Converted Number : ", convertNumber)
		b := int(convertNumber) == randomNumber
		if b {
			fmt.Println("You Win : i ", i, " th chance . Your Number : ", convertNumber, " Random Number : ", randomNumber)
			return
		} else {
			fmt.Println("Try Again")
		}
	}
	if chance == 3 {
		fmt.Println("You Lose : randomNumber ", randomNumber)

	}
}
