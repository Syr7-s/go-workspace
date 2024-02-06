package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	value1 := "10"
	value2 := "5.5"
	operation := "*"
	result := calculate(value1, value2, operation)

	fmt.Printf("value1: %s\nvalue2: %s\noperation: %s\nreturn: %f", value1, value2, operation, result)
}

func calculate(input1 string, input2 string, operation string) float64 {
	// Your code goes here.
	value1 := convertInputToValue(input1)
	value2 := convertInputToValue(input2)
	var result float64 = 0
	switch operation {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subtractValues(value1, value2)
	case "*":
		result = multiplyValues(value1, value2)
	case "/":
		result = divideValues(value1, value2)
	default:
		result = 0
		panic("Invalid operation")
	}
	return result
}

func convertInputToValue(input string) float64 {
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		//message := fmt.Sprintf("%v must be a number",input)
		//panic(message)
		log.Panic("This string can't parse")
	}
	return value
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	if value2 == 0 {
		log.Panic("Zero Division error")
	}
	return value1 / value2
}
