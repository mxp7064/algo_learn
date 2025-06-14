/*
HEAPS (MIN-HEAP AND MAX-HEAP)

WHAT IS A HEAP?
A heap is a specialized binary tree-based data structure that satisfies two properties:

1. HEAP PROPERTY:
   - MIN-HEAP: Every parent node is less than or equal to its children. The smallest element is always at the root.
   - MAX-HEAP: Every parent node is greater than or equal to its children. The largest element is always at the root.

2. COMPLETE BINARY TREE:
   - Every level is completely filled except possibly the last, which is filled from left to right.
   - This structure allows a heap to be stored compactly in a linear array.

WHY USE A HEAP?
The biggest advantage of a heap is that the smallest (in min-heap) or largest (in max-heap) element can be accessed in O(1) time.
This makes heaps ideal for problems where you repeatedly need quick access to min/max value.

COMMONLY ASKED ALGORITHMS AND USE CASES:
- Dijkstra’s Algorithm (Min-Heap for shortest path)
- A* Search Algorithm (Min-Heap using cost + heuristic)
- Prim’s Algorithm for Minimum Spanning Tree (Min-Heap)
- Top-K Largest or Smallest Elements (Min-Heap or Max-Heap)
- Running Median (two heaps: max-heap on left, min-heap on right)
- Huffman Encoding Tree Construction (Min-Heap)
- Task Scheduling or Event Queue (Min-Heap)

ARRAY REPRESENTATION AND FORMULAS:
- Left child of node at index i: 2*i + 1
- Right child of node at index i: 2*i + 2
- Parent of node at index i: (i - 1) / 2

KEY OPERATIONS:
INSERT:
- Add element at the end of the array.
- Restore heap property using bubbleUp (sift up).
- Time Complexity: O(log n)

PEEK:
- Return the root element at index 0.
- Time Complexity: O(1)

POP:
- Replace the root with the last element.
- Remove the last element.
- Restore heap property using bubbleDown (sift down).
- Time Complexity: O(log n)

HEAPIFY (SMART CONSTRUCTION):
- Convert an unordered array into a valid heap in-place.
- Start from the last parent node and call bubbleDown moving upwards.
- Time Complexity: O(n)
*/

/*
Generic Heap implementation in Go using generics

Example usage:

type Item struct {
	Node     string
	Distance int
}

func ItemComp(a, b Item) bool {
	return a.Distance < b.Distance
}

heap := NewHeap(ItemComp)
heap.Insert(Item{"A", 0})

Example min vs max int heap:

Max-heap -> higher number has higher priority:
heap := myheap.NewHeap[int](func(a, b int) bool {
	return a > b
})

Min-heap -> lower number has higher priority
heap := myheap.NewHeap[int](func(a, b int) bool {
	return a < b
})

Why not use GO's heap package?
- Go's `container/heap` package provides heap logic, and we can implement the heap interface but it's verbose and clumsy -
it's not elegant and simple as in some other languages as Python. Instead, it's better to use our implementation
which you can just copy/paste during interview

You can find an example which uses Go's heap package here: algorithms/heap_algos/top_k_elements_Go's_heap_package_test.go
*/

package myheap

// Comparator is a function that returns true if a should come before b
// Used to define min-heap or max-heap behavior
// For example: func(a, b Item) bool { return a.Distance < b.Distance }
type Comparator[T any] func(a, b T) bool

// Heap is a generic binary heap (min-heap or max-heap depending on Comparator)
type Heap[T any] struct {
	data       []T
	comparator Comparator[T]
}

// NewHeap creates an empty heap with the given comparator.
//
// Use this for min heap (lower number has higher priority):
//
//	 myheap.NewHeap[int](func(a, b int) bool {
//		   return a < b
//	 })
//
// Use this for max heap (higher number has higher priority):
//
//	myheap.NewHeap[int](func(a, b int) bool {
//	   return a > b
//	})
func NewHeap[T any](comparator Comparator[T]) *Heap[T] {
	return &Heap[T]{data: []T{}, comparator: comparator}
}

// CreateHeapFromArray builds a heap in-place from the given array using heapify
func CreateHeapFromArray[T any](arr []T, comparator Comparator[T]) *Heap[T] {
	heap := &Heap[T]{data: arr, comparator: comparator}
	// Start heapify from last parent down to root
	// we start from index (n-2)/2 because that is the parent of the last element: ((n-1)-1)/2 = (n-2)/2
	// all nodes before it are guaranteed to have children, i.e. they are parents so we just move i--
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		heap.bubbleDown(i)
	}
	return heap
}

func CreateHeapWithCapacity[T any](c int, comparator Comparator[T]) *Heap[T] {
	// We initialize the slice with length 0 and capacity c.
	// This way we can insert up to c elements without resizing, heap starts empty and behaves correctly
	return &Heap[T]{data: make([]T, 0, c), comparator: comparator}
}

// Insert adds a new element to the heap
func (h *Heap[T]) Insert(val T) {
	h.data = append(h.data, val)
	h.bubbleUp(len(h.data) - 1)
}

// Pop removes and returns the root element (min or max depending on comparator)
func (h *Heap[T]) Pop() T {
	if len(h.data) == 0 {
		var zero T
		return zero
	}
	root := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.bubbleDown(0)
	return root
}

// Peek returns the root element without removing it
func (h *Heap[T]) Peek() T {
	if len(h.data) == 0 {
		var zero T
		return zero
	}
	return h.data[0]
}

// Len returns the number of elements in the heap
func (h *Heap[T]) Len() int {
	return len(h.data)
}

// IsEmpty return true if there are no elements in the heap
func (h *Heap[T]) IsEmpty() bool {
	return h.Len() == 0
}

// Internal method to compare elements using the comparator
func (h *Heap[T]) compare(i, j int) bool {
	return h.comparator(h.data[i], h.data[j])
}

func (h *Heap[T]) bubbleUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if !h.compare(index, parent) {
			break
		}
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		index = parent
	}
}

func (h *Heap[T]) bubbleDown(index int) {
	size := len(h.data)
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < size && h.compare(left, smallest) {
			smallest = left
		}
		if right < size && h.compare(right, smallest) {
			smallest = right
		}
		if smallest == index {
			break
		}
		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		index = smallest
	}
}

// GetHeapArray returns the internal slice\
func (h *Heap[T]) GetHeapArray() []T {
	return h.data
}
