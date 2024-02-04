package main

import "fmt"

func main() {
	//m := new(map[string]int) invalid operation
	//m["theAnswer"] = 42
	//fmt.Println(m)

	m := make(map[string]int)
	m["theAnswer"] = 50
	fmt.Println("The Answeer Value: ", m["theAnswer"])
}
