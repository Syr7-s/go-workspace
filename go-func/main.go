package main

import "fmt"

func main() {
	doSomething()
}

func doSomething() {
	fmt.Println("We are learning go")
	sum := addValues(5, 9)
	fmt.Println("The sum is ", sum)
	sum1, count := addAllValues(5, 78, 69, 45, 12, 98)
	fmt.Println("The sum is ", sum1)
	fmt.Println("Count of items: ", count)
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
