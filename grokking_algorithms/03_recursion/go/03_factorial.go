package main

import (
	"fmt"
)

// Calculates factorial recursively
func fact(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * fact(x-1)
	}
}

func main() {
	fmt.Println(fact(5)) // Output: 120
}
