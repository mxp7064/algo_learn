package main

import (
	"fmt"
)

/*
Binary search uses divide and conquer strategy - in each step we eliminate half of possible choices
*/

// Binary search function
func binarySearch(list []int, item int) *int {
	// low and high keep track of which part of the list you'll search in
	low := 0
	high := len(list) - 1

	// While you haven't narrowed it down to one element ...
	for low <= high {
		// ... check the middle element
		mid := (low + high) / 2
		guess := list[mid]

		// Found the item
		if guess == item {
			return &mid
		}

		// The guess was too high
		if guess > item {
			high = mid - 1
		} else {
			// The guess was too low
			low = mid + 1
		}
	}

	// Item doesn't exist
	return nil
}

func main() {
	myList := []int{1, 3, 5, 7, 9}

	// Prints 1
	if result := binarySearch(myList, 3); result != nil {
		fmt.Println(*result)
	} else {
		fmt.Println("nil")
	}

	// 'nil' means the item wasn't found
	if result := binarySearch(myList, -1); result != nil {
		fmt.Println(*result)
	} else {
		fmt.Println("nil")
	}
}
