/*
Sorted array (as a data structure)

This is a simple demonstration of basic operations on sorted array. This is just a good to know thing in general so it's clear
what are time complexities of a sorted array when we compare it to other approaches.

Time complexity:
- Access: O(1)
	- when you do arr[i], you’re asking - what is the address of the i-th element, so I can go to that address and look up the element value
	- array is stored as a contiguous block of memory
	- it has a base starting memory address (address of first element)
	- each element has a fixed size S (some constant - for example 8 bytes for int64)
    - so address of arr[i] = base + i × S
	- this is a simple arithmetic operation, which the CPU performs in constant time (no loops or any kind of traversal)
- Search: O(log n)
	- binary search
- Insert, Delete: O(n)
	- O(log n) to find the index + O(n) to shift elements
	- total is: O(log n + n) but O(n) dominates so we simplify it to O(n)
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(searchSorted([]int{1, 4, 5, 7, 8}, 5))     // 2
	fmt.Println(searchSorted([]int{1, 4, 5, 7, 8}, 55))    // -1
	fmt.Println(insertSorted([]int{2, 3, 4, 7, 8, 10}, 5)) // [2 3 4 5 7 8 10]
	fmt.Println(deleteSorted([]int{2, 3, 4, 7, 8, 10}, 7)) // [2 3 4 8 10]
}

// searchSorted searches for a given value in the array and returns its index, if value doesn't exist in array it returns -1
func searchSorted(arr []int, value int) int {
	i := sort.SearchInts(arr, value) // O(log n)
	if i == len(arr) {
		return -1
	}
	return i
}

// Insert value into sorted array
func insertSorted(arr []int, value int) []int {
	i := sort.SearchInts(arr, value) // O(log n)
	// grow array with arbitrary/dummy value
	arr = append(arr, 0)
	// shift elements right starting from end
	for j := len(arr) - 1; j > i; j-- {
		arr[j] = arr[j-1]
	}
	arr[i] = value // set the inserted value
	return arr
}

// Delete value from sorted array
func deleteSorted(arr []int, value int) []int {
	i := sort.SearchInts(arr, value)     // O(log n)
	if i < len(arr) && arr[i] == value { // if the value exists in array, delete it
		// shift left
		for j := i; j < len(arr)-1; j++ {
			arr[j] = arr[j+1]
		}
		arr = arr[:len(arr)-1] // remove last element
	}
	return arr
}
