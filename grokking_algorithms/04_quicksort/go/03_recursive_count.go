package main

import (
	"fmt"
)

// Recursively counts the number of elements in a slice
func count(list []int) int {
	// Base case: empty list
	if len(list) == 0 {
		return 0
	}

	// Recursive case: 1 (current element) + count of the rest
	return 1 + count(list[1:])
}

func main() {
	fmt.Println(count([]int{1, 2, 3, 4})) // Output: 4
}
