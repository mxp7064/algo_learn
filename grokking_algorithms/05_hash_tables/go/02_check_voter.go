package main

import (
	"fmt"
)

// Declare a global map to track who has voted
var voted = make(map[string]bool)

func checkVoter(name string) {
	if voted[name] {
		fmt.Println("kick them out!")
	} else {
		voted[name] = true
		fmt.Println("let them vote!")
	}
}

func main() {
	checkVoter("tom")
	checkVoter("mike")
	checkVoter("mike")
}
