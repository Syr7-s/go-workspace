package main

import "fmt"

func main() {
	slice := []string{"apple", "banana", "orange", "date"}

	result := convertToMap(slice)
	for k, v := range result {
		fmt.Println(k, ": ", v)
	}
}

func convertToMap(items []string) map[string]float64 {

	// Create a map object
	result := make(map[string]float64)
	for k := range items {
		result[items[k]] = 100 / float64(len(items))
	}
	// Your code goes here
	return result
}
