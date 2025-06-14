package main

import (
	"fmt"
)

// Calculates the sum of all elements in an array
func loopSum(arr []int) int {
	total := 0
	for _, x := range arr {
		total += x
	}
	return total
}

func main() {
	fmt.Println(loopSum([]int{1, 2, 3, 4})) // Output: 10
}
