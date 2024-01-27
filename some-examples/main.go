package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	//isPerfectNumber
	//isPerfectNumber()
	//findTheLargestSecondNumberInArray()
	isPalindrome()

}

func isPerfectNumber() {
	fmt.Println("Enter the number : ")
	var number int
	fmt.Scan(&number)
	var sum int = 0
	for i := 1; i < number; i++ {
		if number%i == 0 {
			sum += i
		}
	}
	if number == sum {
		fmt.Printf("The %d number is perfect number", number)
	} else {
		fmt.Printf("The %d number is not perfect number", number)
	}
}

func findTheLargestSecondNumberInArray() {
	fmt.Println("Enter array len : ")
	var arrLen int
	fmt.Scan(&arrLen)
	randomArr := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		randomArr[i] = int(rand.Int31n(1000))
	}
	for i := 0; i < arrLen; i++ {
		fmt.Println(randomArr[i])
	}
	fmt.Println("-----------------------------")
	for i := 0; i < arrLen; i++ {
		for j := i; j < arrLen; j++ {
			if randomArr[i] < randomArr[j] {
				tempValue := randomArr[j]
				randomArr[j] = randomArr[i]
				randomArr[i] = tempValue
			}
		}
	}
	fmt.Printf("Find the largest second number : %d", randomArr[1])

}

func isPalindrome() {
	fmt.Println("Enter text : ")
	var text string
	//fmt.Scan(&text)
	reader := bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')
	var count int
	text = strings.TrimSpace(text)
	for i := 0; i < len(text); i++ {
		if text[i] == text[len(text)-1-i] {
			count += 1
		}
	}

	if count == len(text) {
		fmt.Printf("The %s word is palindrome", text)
	} else {
		fmt.Printf("The %s word is not palindrome", text)
	}
}
