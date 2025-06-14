package main

import "fmt"

/*
Selection Sort — Two Versions (non-in-place and in-place)

Problem:
Sort an array of integers using Selection Sort.

Selection Sort is a simple, comparison-based sorting algorithm.
The idea:
- Repeatedly find the smallest element in the unsorted portion of the array
- Move it to the beginning (either by appending it to a result slice or swapping in place)

Time Complexity:
- O(n^2) — for each of the n elements, you scan the rest (n-i)
- Not efficient for large arrays

Space Complexity:
- O(n) for non-in-place version (copies data)
- O(1) for in-place version

When to use:
- Educational purposes
- When simplicity is more important than performance
*/

func main() {
	// Demonstration
	fmt.Println("Non-in-place:")
	fmt.Println(selectionSort([]int{3, 2, 5}))
	fmt.Println(selectionSort([]int{3}))
	fmt.Println(selectionSort([]int{3, 3, 3}))
	fmt.Println(selectionSort([]int{}))

	fmt.Println("\nIn-place:")
	fmt.Println(selectionSortInPlace([]int{5, 3, 1, 4}))
}

// selectionSort (non-in-place version)
// Returns a new sorted slice without modifying the input
func selectionSort(arr []int) []int {
	// Copy the input array so we don't mutate it
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)

	result := make([]int, 0, len(arr))

	// Repeatedly find and remove the smallest element
	for len(copyArr) > 0 {
		minIdx := findMin(copyArr)
		// Append the smallest to the result
		result = append(result, copyArr[minIdx])
		// Remove it from the working array
		copyArr = append(copyArr[:minIdx], copyArr[minIdx+1:]...)
	}

	return result
}

// selectionSortInPlace
// Sorts the array in-place using swapping
func selectionSortInPlace(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		// Find the index of the smallest element in arr[i:] (unsorted part of the array)
		// index returned by findMin is relative to arr[i:], not to the original slice so we need to fix it by adding i
		minIdx := i + findMin(arr[i:])
		// Swap the current element with the smallest found in the unsorted part of the array
		arr[i], arr[minIdx] = arr[minIdx], arr[i] // now element i is sorted
	}
	return arr
}

// findMin returns the index of the smallest value in the array
func findMin(arr []int) int {
	if len(arr) == 0 {
		return -1 // edge case: empty slice
	}
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}
