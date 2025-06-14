package main

import (
	"fmt"
)

// Recursively finds the maximum element in a slice
func maxRec(list []int) int {
	// Base case: if the list has only 2 elements, return the greater one
	if len(list) == 2 {
		if list[0] > list[1] {
			return list[0]
		} else {
			return list[1]
		}
	}

	// Recursive case: find the max in the rest of the list
	subMax := maxRec(list[1:])
	if list[0] > subMax {
		return list[0]
	} else {
		return subMax
	}
}

func main() {
	fmt.Println(maxRec([]int{1, 5, 10, 3, 8})) // Output: 10
}
