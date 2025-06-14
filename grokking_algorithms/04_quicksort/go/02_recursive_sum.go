package main

import (
	"fmt"
)

// Recursively sums all elements in a slice
func recSum(list []int) int {
	// Base case: empty list
	if len(list) == 0 {
		return 0
	}

	// Recursive case: first element + sum of the rest
	return list[0] + recSum(list[1:])
}

func main() {
	fmt.Println(recSum([]int{1, 2, 3, 4})) // Output: 10
}
