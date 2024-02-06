package main

import "fmt"

func main() {
	poodle := Dog{"Poodle", 10, "Woof!"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Speak()
	poodle.Sound = "Arf!"
	poodle.Speak()
	// Go not support override
	// Each method has a unique name
	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes()
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
