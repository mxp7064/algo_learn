/*
Problem: Given an array of integers, find the k largest/smallest values.

Approach:
We use a min/max heap of size k to keep track of the top k largest/smallest elements.
The heap is always ordered so that the smallest/largest element is on top.

Why not just sort and slice first k elements?
- Sorting takes O(n log n) time.
- This solution takes O(n log k) time, which is faster when k is much smaller than n

Time Complexity:
- Building the heap: O(k)
- Iterating and updating heap: O((n-k) * log k)
- Total: O(n log k)

Space Complexity: O(k)
- for the heap of size k

TOP-K ELEMENT SELECTION PATTERN:

1. K SMALLEST ELEMENTS:
  - Use a max-heap of size k
  - Iterate through array:
  - If heap has fewer than k elements, insert
  - Else if current element is smaller than root, pop and insert it
  - Heap root always holds the worst (largest) among the current k smallest

2. K LARGEST ELEMENTS:
  - Use a min-heap of size k
  - Iterate through array:
  - If heap has fewer than k elements, insert
  - Else if current element is larger than root, pop and insert it
  - Heap root always holds the worst (smallest) among the current k largest

Here we are using our Go heap implementation (in panca.com/algo/myheap package).
You can find an example which uses Go's heap package below at the end of this file.
*/

package heap_algos

import (
	"container/heap"
	"fmt"
	"panca.com/algo/myheap"
	"testing"
)

func Test_kSmallestElements(t *testing.T) {
	t.Log("kSmallestDumb: ", kSmallestDumb(3, []int{5, 2, 7, 3, 9})) // [2 3 5]
	t.Log("kSmallest", kSmallest(3, []int{5, 2, 7, 3, 9}))           // [5 2 3]
	t.Log("kLargest", kLargest(3, []int{5, 2, 7, 3, 9}))             // [5 7 9]
}

// kSmallestDumb is the simple approach using a full Min-Heap
// Step 1: Convert entire array into a Min-Heap (heapify)
// Step 2: Pop() the smallest element k times
// Time complexity: O(n + k log n)
// Space complexity: O(n) to store heap
func kSmallestDumb(k int, arr []int) (result []int) {
	if k <= 0 || len(arr) == 0 {
		return result
	}

	// Create min-heap from the array
	heap := myheap.CreateHeapFromArray(arr, func(a, b int) bool {
		return a < b
	})

	// Pop k smallest values
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop())
	}

	return result
}

// kSmallest is the optimized approach using a fixed-size Max-Heap
// We maintain a max-heap of size k that always contains the k smallest elements seen so far
// Step 1: Fill the heap with first k elements
// Step 2: For each remaining element:
//   - If it is smaller than the max (root), replace the root
//
// Step 3: Return the heap contents (these are the k smallest)
//
// Why a max-heap?
//   - Because we want to efficiently track and potentially replace the current largest among the k smallest so far
//   - Root gives us quick access to the largest → easy to decide whether to keep or replace
//
// Intuition:
//   - With each new element after the first k, we compare it with the largest of the k smallest so far (i.e. the root)
//   - If it is smaller, we replace the root
//   - By the end, we have eliminated all non-candidates, and only the best k remain
//
// Time complexity: O(n log k)
// Space complexity: O(k)
func kSmallest(k int, arr []int) []int {
	if k <= 0 || len(arr) == 0 {
		return nil
	}

	// Create max heap of capacity k
	heap := myheap.CreateHeapWithCapacity(k, func(a, b int) bool {
		return a > b
	})

	for i := 0; i < len(arr); i++ {
		if heap.Len() < k {
			// Fill the heap initially
			heap.Insert(arr[i])
		} else if arr[i] < heap.Peek() {
			// new value is smaller than current largest among k smallest
			// - root is not among smallest k for sure so we can pop it
			// - new value could be among the k smallest so we insert it
			// in each iteration we are finding better candidates and in the end we will have k smallest elements in the heap
			heap.Pop()          // remove root and reheapify
			heap.Insert(arr[i]) // insert and reheapify
		} // else: arr[i] >= heap.Peek() -> ignore, element is worst than the current worst -> not among k smallest
	}

	return heap.GetHeapArray()
}

// kLargest is just a reverse of kSmallest
func kLargest(k int, arr []int) []int {
	if k <= 0 || len(arr) == 0 {
		return nil
	}

	// Create min heap of capacity k
	heap := myheap.CreateHeapWithCapacity(k, func(a, b int) bool {
		return a < b
	})

	for i := 0; i < len(arr); i++ {
		if i < k { // or heap.Len() < k - it's the same
			heap.Insert(arr[i])
		} else if arr[i] > heap.Peek() {
			heap.Pop()
			heap.Insert(arr[i])
		}
	}

	return heap.GetHeapArray()
}

// ==========================================================================================

/*
Find k largest elements using Go's heap package
We need to create the heap by implementing the heap interface methods: Len, Less, Swap, Push and Pop
And then we can use the heap package functions such as heap.Init, heap.Push and heap.Pop on our heap instance.
This approach is too verbose - compared to other languages as Python where it's much more simple and elegant (no need for implementing an interface).
It's better to just use our Go heap implementation which can be found in panca.com/algo/myheap package
*/
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // for min-heap, for max-heap it's opposite
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1] // get the last value
	*h = old[:n-1]  // shrink the slice (remove last)
	return val
}

func findKLargestGoHeap(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	h := &MinHeap{}
	heap.Init(h)

	// Build initial heap with first k elements
	for _, num := range nums[:k] {
		heap.Push(h, num)
	}

	// Iterate through remaining elements
	for _, num := range nums[k:] {
		if num > (*h)[0] {
			// If current number is larger than min in heap, replace it
			// Pop and push do some swapping and wierd things under the hood
			// and then call our Pop and Push methods
			heap.Pop(h)
			heap.Push(h, num)
		}
	}

	return *h
}

func Test_findKLargestGoHeap(t *testing.T) {
	arr := []int{7, 2, 9, 4, 1, 11, 8, 3, 10}
	k := 4
	res := findKLargestGoHeap(arr, k)
	fmt.Printf("Top %d largest values: %v\n", k, res) // 11, 10, 9, 8
}
