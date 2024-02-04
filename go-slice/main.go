package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println("Colors : ", colors, "Len of Colors : ", len(colors))
	colors = append(colors, "Purple", "Yellow", "Black")
	fmt.Println("Colors : ", colors, "Len of Colors : ", len(colors))
	colors = append(colors[2:len(colors)])
	fmt.Println("Colors : ", colors, "Len of Colors : ", len(colors))
	colors = append(colors[:len(colors)-1]) // delete last item
	fmt.Println("Colors : ", colors, "Len of Colors : ", len(colors))

	numbers := make([]int, 5)
	numbers[0] = rand.Intn(150)
	numbers[1] = rand.Intn(150)
	numbers[2] = rand.Intn(150)
	numbers[3] = rand.Intn(150)
	numbers[4] = rand.Intn(150)

	fmt.Println("Numbers : ", numbers, "Len of Numbers : ", len(numbers), "Capacity of Numbers : ", cap(numbers))
	numbers = append(numbers, rand.Intn(150))
	fmt.Println("Numbers : ", numbers, "Len of Numbers : ", len(numbers), "Capacity of Numbers : ", cap(numbers))

	//sort of numbers

	sort.Ints(numbers)
	fmt.Println(numbers)

}
