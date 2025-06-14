package main

import (
	"fmt"
)

func main() {
	// Create a new map
	book := make(map[string]float64)

	// an apple costs 67 cents
	book["apple"] = 0.67
	// milk costs $1.49
	book["milk"] = 1.49
	book["avocado"] = 1.49

	fmt.Println(book)
}
