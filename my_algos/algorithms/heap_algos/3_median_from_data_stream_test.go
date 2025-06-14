/*
FIND MEDIAN FROM DATA STREAM – HEAP-BASED VERSION (INTERVIEW-READY)

🟩 PROBLEM STATEMENT:
Design a data structure that supports:
- AddNum(num int) – adds a number to the stream
- GetMedian() float64 – returns the current median of all numbers added so far

You should be able to add numbers one by one and get the median at any point efficiently.

📥 INPUT:
- A stream of integers added via AddNum(num)

📤 OUTPUT:
- The median of all elements seen so far via GetMedian()

🧠 CORE IDEA:
To find the median dynamically, we need to keep track of the "middle" value(s) (depending on odd/even length).
We split the numbers into two halves:
- Left half (smaller numbers) → max-heap
- Right half (larger numbers) → min-heap

Why this works:
- The left max-heap keeps track of the largest value in the left half
- The right min-heap keeps track of the smallest value in the right half
- The median is:
  - Just the top of the max-heap (if total is odd)
  - Average of tops of both heaps (if total is even)
	- In Mathematics, median can be any value in between but in DSA we usually take the average

This works due to transitivity:
- All elements in max-heap ≤ max-heap.Peek()
- All elements in min-heap ≥ min-heap.Peek()
→ So we know:
  max(left) ≤ median ≤ min(right)

⚠️ The median (by definition) is the value such that:
- Half the elements are ≤ median
- Half are ≥ median
- If the total count is even, the median is the average of the two middle values
  → note that it does not have to be an actual value from the dataset (e.g., median of [1, 3] is 2)

🆚 NAIVE APPROACH:
- Store all numbers in an array
- Sort on every insertion → O(n log n)
- Median is middle element (or average of two middles)

🟡 Limitation: sorting every time is too expensive for real-time streams

✅ OPTIMAL HEAP APPROACH:
- Use two heaps:
  - Max-heap for left half (lower values)
  - Min-heap for right half (higher values)
- Always maintain:
  - Either same size
  - Or left max-heap has 1 more element than right min-heap

⏱️ TIME COMPLEXITY:
- AddNum(num) → O(log n)
  - Insert + optional rebalance with insert/pop
  - AddNum is called for each number we add -> O(n log n)
- GetMedian() → O(1)
  - Just a peek at one or two heap roots
- Total → O(n log n)

🧠 Why better than sorting:
- Sorting is O(n log n) per median query
- This is O(log n) per insert, and O(1) per median query

*/

package heap_algos

import (
	"panca.com/algo/myheap"
	"testing"
)

type DataStream struct {
	leftSide  *myheap.Heap[int] // Max-heap for lower half
	rightSide *myheap.Heap[int] // Min-heap for upper half
}

// AddNum inserts a new number into the correct heap and rebalances the two sides if needed
// Time: O(log n)
func (d *DataStream) AddNum(num int) {
	left := d.leftSide
	right := d.rightSide

	// Decide where to insert:
	// If left is empty or num ≤ top of max-heap, insert to left (lower half)
	if left.IsEmpty() || num <= left.Peek() {
		left.Insert(num)
	} else {
		right.Insert(num)
	}

	// Rebalance the heaps so that:
	// - max-heap is either equal in size or one element bigger than min-heap
	if left.Len() > right.Len()+1 {
		right.Insert(left.Pop())
	}
	if right.Len() > left.Len() {
		left.Insert(right.Pop())
	}
}

// GetMedian returns the median of the current stream
// Time: O(1)
func (d *DataStream) GetMedian() float64 {
	leftLen := d.leftSide.Len()
	rightLen := d.rightSide.Len()
	totalLen := leftLen + rightLen

	if totalLen%2 == 0 {
		// Even number of elements → average of two middle elements
		return float64(d.leftSide.Peek()+d.rightSide.Peek()) / 2
	} else {
		// Odd → middle is always at top of max-heap
		return float64(d.leftSide.Peek())
	}
}

// CreateDataStream initializes the DataStream with two heaps
func CreateDataStream() *DataStream {
	return &DataStream{
		leftSide: myheap.NewHeap[int](func(a, b int) bool {
			return a > b // Max-heap: higher number has higher priority
		}),
		rightSide: myheap.NewHeap[int](func(a, b int) bool {
			return a < b // Min-heap: lower number has higher priority
		}),
	}
}

// Test_DataStream demonstrates how the data structure behaves step-by-step
func Test_DataStream(t *testing.T) {
	ds := CreateDataStream()

	ds.AddNum(5)
	t.Logf("Median after [5]: %.1f", ds.GetMedian()) // 5.0

	ds.AddNum(1)
	t.Logf("Median after [1, 5]: %.1f", ds.GetMedian()) // 3.0

	ds.AddNum(2)
	t.Logf("Median after [1, 2, 5]: %.1f", ds.GetMedian()) // 2.0

	ds.AddNum(3)
	t.Logf("Median after [1, 2, 3, 5]: %.1f", ds.GetMedian()) // 2.5

	ds.AddNum(11)
	t.Logf("Median after [1, 2, 3, 5, 11]: %.1f", ds.GetMedian()) // 3.0
}
