package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	value1 := "10"
	value2 := "5.5"
	result := calculate(value1, value2)
	fmt.Println("Result : ", result)
}

func calculate(value1 string, value2 string) float64 {
	// Your code goes here.
	first, _ := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	// Convert the first string to a float64
	second, _ := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	// Convert the second string to a float64

	// Calculate and return the result

	return first + second
}
