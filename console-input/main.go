package main

// Get Console Input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const userNameConst string = "dev123"
const passwordConst string = "1233"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text : ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You  Entered : ", input)

	var chance int = 0
	for i := chance; chance <= 2; chance++ {
		fmt.Print("Enter userName : ")
		userName, _ := reader.ReadString('\n')
		fmt.Print("Enter password : ")
		password, _ := reader.ReadString('\n')

		isValidUser := isValidUserFunc(userName, password)
		fmt.Println("Valid : ", isValidUser)
		if isValidUser {
			fmt.Printf("%d th chance ", i)
			fmt.Print("Is Valid User : ")
			return
		}
	}

	//chance += 1

	if chance == 3 {
		fmt.Printf("UserName password incorrect userName %s password %s", userNameConst, passwordConst)
	}

}

func isValidUserFunc(userName string, password string) bool {
	if strings.Trim(userName, "\n") == userNameConst && strings.Trim(password, "\n") == passwordConst {
		return true
	} else {
		return false
	}
}
