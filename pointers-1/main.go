package main

import "fmt"

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println("Valueof anInt : ", &anInt)
	fmt.Println("Value of p : ", p)
	fmt.Println("Value of p : ", &p)
	fmt.Println("Value of p : ", *p)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Value of value1 : ", &value1)
	fmt.Println("Value of pointer1 : ", &pointer1)
	fmt.Println("Value1 :  ", *pointer1)
	fmt.Println("pointer1 :  ", pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Pointer 1  :", *pointer1)
	fmt.Println("Value1 : 1 ", value1)
	fmt.Println("Value of value1 : ", &value1)
	fmt.Println("Value of pointer1 : ", &pointer1)
	fmt.Println("pointer1 :  ", pointer1)
}
