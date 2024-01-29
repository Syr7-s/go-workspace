package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("--------------- Math Package ---------------")
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer Sum : ", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum : ", floatSum)
	floatSumRound := math.Round(floatSum)
	fmt.Println("Round Float Sum : ", floatSumRound)
	floatSumFloor := math.Floor(floatSum)
	fmt.Println("Floor Float Sum : ", floatSumFloor)

	circleRadius := 15.56
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumerence: %.2f\n", circumference)

}
