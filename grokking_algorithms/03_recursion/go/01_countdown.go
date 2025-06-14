package main

import (
	"fmt"
)

// Countdown function
func countdown(i int) {
	fmt.Println(i)

	// Base case
	if i <= 0 {
		return
	}

	// Recursive case
	countdown(i - 1)
}

func main() {
	countdown(5)
}
