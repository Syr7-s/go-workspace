package main

import (
	"fmt"
	"strconv"
)

func main() {
	value1 := "10"
	value2 := "5.5"
	result := calculate(value1, value2)
	fmt.Println("Result : ", result)
}

func calculate(value1 string, value2 string) float64 {
	// Your code goes here.
	first, _ := strconv.ParseFloat(value1, 64)
	// Convert the first string to a float64
	second, _ := strconv.ParseFloat(value2, 64)
	// Convert the second string to a float64

	// Calculate and return the result

	return first + second
}
