package main

import "fmt"

func main() {
	var colors [4]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	colors[3] = "Yellow"

	fmt.Println("Colors : ", colors)

	fmt.Println("Color[0] : ", colors[0])
	var number = [5]int{5, 3, 1, 2, 6}
	fmt.Println("Numbers : ", number)
	fmt.Println("Number[2] : ", number[2])
	fmt.Println("Number of Colors  : ", len(colors))
	fmt.Println("Number of Numbers  : ", len(number))

}
