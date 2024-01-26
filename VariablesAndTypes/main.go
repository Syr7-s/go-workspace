package main

import (
	"fmt"
	"reflect"
)

// Const : Constant  Declaration : const <name> <type> = <value>

const aConstString string = "This is a constant string"

func main() {

	/*
			Storing Data in Variables
			Go is a statically typed language, which means that variables always have a specific type and that type cannot change.
			The type of a variable is always known at compile time.
			The := syntax is shorthand for declaring and initializing a variable in one line.
			All variables have assigned types
			Variables can be declared without a value
			You can set types explicitly or implicitly

			Variable Types
			- bool
			- string
			- int  int8  int16  int32  int64
			- uint uint8 uint16 uint32 uint64 uintptr
			- byte // alias for uint8
			- rune // alias for int32
			- float32 float64
			- complex64 complex128

			Zero Values
			Variables declared without an explicit initial value are given their zero value.

				Data Collections
				- Arrays : fixed length
				- Slices : dynamic length
				- Maps : key value pairs
				- Structs : collection of fields

				Language Organization
					- Functions
					- Interfaces
					- Channels
			Data Management
				- Pointers : reference to memory address. Used to pass data by reference.
		Go also supports pointers, reference variables that point to an address in memory to refer to another value.
		These are the built-in types in Go but they're just the beginning point, because again, you can create your own data types in this language.
	*/

	var aString string = "This is my a first example" // explicit type declaration
	fmt.Printf("aString: %s\n", aString)
	fmt.Println("aString Type : ", reflect.TypeOf(aString))

	var anInteger int = 45
	fmt.Printf("anInteger: %d\n", anInteger)

	var defaultInt int
	fmt.Printf("defaultInt: %d\n", defaultInt)

	var anotherString = "This is another string" // implicit type declaration
	fmt.Printf("anotherString: %s\n", anotherString)
	fmt.Println("anotherString Type : ", reflect.TypeOf(anotherString))

	var anotherInt = 154
	fmt.Println("anotherInt Type : ", reflect.TypeOf(anotherInt))

	myInt := 4569
	fmt.Println("MyInt Type : ", reflect.TypeOf(myInt))

	fmt.Println("Contstant String : ", aConstString)

	numbers := []int{5, 9, 6, 7, 89}
	sumOfListResult := sumOfList(numbers)
	fmt.Println("Sum of List Value : ", sumOfListResult)
}

func sumOfList(numbers []int) int {
	var sum int = 0
	fmt.Println(sum)
	for _, number := range numbers {
		sum += number
	}
	return sum
}
