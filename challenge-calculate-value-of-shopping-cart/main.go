package main

import "fmt"

func main() {
	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	var cucumber = cartItem{"cucumber", 5.49, 12}
	cart = append(cart, apples, oranges, bananas, cucumber)
	result := calculateTotal(cart)
	fmt.Println(result)

}

type cartItem struct {
	name     string
	price    float64
	quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {
	var totalPrice float64 = 0

	for k := range cart {
		totalPrice += cart[k].price * float64(cart[k].quantity)

	}
	return totalPrice
}
