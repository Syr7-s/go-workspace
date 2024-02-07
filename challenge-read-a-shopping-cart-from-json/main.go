package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonString :=
		`[
			{"name":"apple","price":2.99,"quantity":3},
  			{"name":"orange","price":1.50,"quantity":8},
  			{"name":"banana","price":0.49,"quantity":12}
		]`

	result := getCartFromJson(jsonString)
	for _, cart := range result {
		fmt.Println(cart.Name, " ", cart.Price, " ", cart.Quantity)
	}
}

type cartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []cartItem {
	var cart []cartItem
	// Your code goes here.
	//decoder := json.NewDecoder(strings.NewReader(jsonString))
	//_, err := decoder.Token()
	//if err != nil {
	//	panic(err)
	//}
	//var cartItem cartItem
	//for decoder.More() {
	//	err := decoder.Decode(&cartItem)
	//	if err != nil {
	//		panic(err)
	//	}
	//	cart = append(cart, cartItem)
	//}
	//
	err := json.Unmarshal([]byte(jsonString), &cart)
	if err != nil {
		panic("Error reading json string")
	}
	return cart
}
