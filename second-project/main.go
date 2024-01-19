package main

import (
	"fmt"
	"sort"
	"strings"
)

const showExpectedResult = true

//const showHints = false

func findLargest(numbers []int) int {
	sort.Ints(numbers)
	return numbers[len(numbers)-1]
}

func findSmallerNumberWithLoop(numbers []int) int {
	smallest := numbers[0]
	for _, number := range numbers {
		if number < smallest {
			smallest = number
		}
	}
	return smallest
}

func stringWorlds(message string) {
	fmt.Println("Message is ", message)
	fmt.Println("Message len is ", len(message))
	msg := strings.Split(message, " ")
	fmt.Println("Message word count is ", len(msg))
	for _, v := range msg {
		fmt.Println(v)
	}
	var count = 0
	for _, v := range message {
		if v == 'a' {
			count++
		}
	}
	fmt.Println("Message a count is ", count)
	var alphabetCount map[string]int = make(map[string]int)

	for i := 0; i < len(message); i++ {
		if string(message[i]) == " " {
			continue
		} else {

			if alphabetCount[string(message[i])] == 0 {
				alphabetCount[string(message[i])] = 1
			} else {
				alphabetCount[string(message[i])]++
			}
		}
	}
	for i := 0; i < len(alphabetCount); i++ {
		if alphabetCount[string(message[i])] != 0 {
			fmt.Println("Alphabet ", string(message[i]), " count is ", alphabetCount[string(message[i])])
		}
	}
}

func main() {
	numbers := []int{7, 17, 13, 19, 5}
	largest := findLargest(numbers)

	if showExpectedResult {
		fmt.Println("Expected result is 19")
	}

	fmt.Println("The largest number is", largest)

	smallNumber := findSmallerNumberWithLoop(numbers)
	fmt.Println("The smallest number is", smallNumber)

	fmt.Println("String len and word count")

	message := "This is a first message from goland"
	fmt.Println("Message is ", message)
	fmt.Println("********************************************")
	stringWorlds(message)
	fmt.Println("********************************************")

	sum := 0
	for j := 0; j < 100; j++ {
		sum += j
	}

}
