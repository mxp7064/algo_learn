package main

import (
	"fmt"
)

func greet2(name string) {
	fmt.Println("how are you, " + name + "?")
}

func bye() {
	fmt.Println("ok bye!")
}

func greet(name string) {
	// call stack: greet
	fmt.Println("hello, " + name + "!")
	greet2(name) // call stack: greet, greet2
	// cal stack: greet
	fmt.Println("getting ready to say bye...")
	bye() // call stack: greet, bye
	// call stack: greet
}

func main() {
	greet("adit") // call stack: greet
	// call stack: none (ofcourse not counting main)
}
