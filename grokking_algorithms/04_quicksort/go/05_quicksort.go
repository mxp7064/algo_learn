package main

import "fmt"

/*
Quicksort — Recursive Implementation (Functional Style)

Problem:
Sort an array of integers using the Quicksort algorithm.

Core idea (Divide and Conquer):
1. Pick a pivot
2. Partition the array into:
   - Elements less than the pivot
   - Elements greater than the pivot
3. Recursively sort the left and right partitions
4. Combine them: [sorted left] + pivot + [sorted right]

Why pivot choice matters:
- Good pivot (middle, random, median-like) → roughly equal partitions
- Bad pivot (e.g. first/last in sorted input) → highly unbalanced partitions

Intuitive Time Complexity:
- Every time we split the array "well", we cut the problem in half → log n levels of recursion
- At each level, we process n elements → n work per level
- So total work = n (per level) * log n levels → O(n log n)

Worst case:
- Happens when pivot always gives an unbalanced split
- For example we chose first element as pivot and the input array is already sorted
- Then we have n levels of recursion → O(n^2)

Space Complexity:
- O(log n) recursion depth in balanced case
- O(n) total in functional style (due to slice copying)

Sorting Algorithms – Summary and Comparison:

Selection Sort / Bubble Sort:
- Time: O(n^2)
- Very slow
- Used mainly for learning and demonstration
- Not used in practice

Merge Sort:
- Time: O(n log n) guaranteed, even in worst case
- Stable: preserves order of equal elements
- Requires extra memory: O(n)
- Used when stability is important (e.g. sorting objects)
- Good for linked lists and external sorting

Quicksort:
- Time: O(n log n) average, O(n^2) worst case (e.g. sorted input with bad pivot)
- In-place, uses less memory than Merge Sort
- Not stable, but usually not a problem
- Fast in practice and widely used
- With tweaks (e.g. random pivot or introsort), worst-case performance can be avoided
- Used in many standard libraries (Go, Java, C++, etc.)
*/

func main() {
	fmt.Println(qs([]int{5, 3, 7, 8, 1, 2, 4}))         // Output: [1 2 3 4 5 7 8]
	fmt.Println(qs([]int{555, 128, 57, 6, 5, 555, 33})) // Output: [5 6 33 57 128 555 555]
	fmt.Println(qs([]int{3, 3, 3}))                     // Output: [3 3 3]
	fmt.Println(qs([]int{}))                            // Output: []
}

// qs sorts the input slice using recursive Quicksort and returns a new sorted slice
func qs(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Choose a middle pivot for better balance
	midIdx := (len(arr) - 1) / 2
	pivot := arr[midIdx]

	var less, greater []int

	// Partition the array into elements less than and greater than the pivot
	for _, el := range arr {
		if el < pivot {
			less = append(less, el)
		} else if el > pivot { // we skip the pivot because it will be added in next step
			greater = append(greater, el)
		}
	}

	// Recursively sort the partitions and combine them with the pivot
	return append(append(qs(less), pivot), qs(greater)...) // or append(qs(less), append([]int{pivot}, qs(greater)...)...)
}
