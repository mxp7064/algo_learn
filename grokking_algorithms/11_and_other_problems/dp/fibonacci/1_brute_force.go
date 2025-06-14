package main

import "fmt"

/*
========================================================================================
🧩 Fibonacci Numbers — Brute Force Recursive Approach (Exponential Time)
========================================================================================

Problem:
Compute the n-th Fibonacci number where:
  Fib(0) = 0
  Fib(1) = 1
  Fib(n) = Fib(n-1) + Fib(n-2) for n ≥ 2

Approach:
- This implementation directly follows the mathematical recursive definition.
- It uses plain recursion with no memoization/caching.
- It recomputes many values multiple times, leading to exponential growth in calls.

----------------------------------------------------------------------------------------

🧠 Time & Space Complexity

Time Complexity:   O(2ⁿ) — exponential due to overlapping subproblems
Space Complexity:  O(n)  — max depth of the recursive call stack

🔍 Why is Time Complexity O(2ⁿ), not O(n!)?

Each fib(n) call recursively branches into two more calls:
  fib(n) = fib(n-1) + fib(n-2)

This creates a binary recursion tree with roughly 2^n nodes (calls).

🧪 Example: Call tree for fib(5)

                      fib(5)
                    /        \
               fib(4)        fib(3)
              /     \        /     \
         fib(3)   fib(2)  fib(2)  fib(1)
        /    \    /   \   /   \
   fib(2) fib(1) fib(1) fib(0) ...
   ...

Notice:
- fib(2), fib(1), etc. are recomputed many times
- The number of calls grows roughly like a binary tree → O(2ⁿ)

✅ It is not O(n!) — factorial time arises in problems where recursion involves branching into n, n-1, ... subproblems (like permutations).

📌 Why is this solution considered "Top-Down"?
- We start from the original problem fib(n) and recursively break it down into smaller subproblems: fib(n-1), fib(n-2), ..., fib(0)
- We go from the "top" (the final goal) down to the "bottom" (the base cases)

*/

func main() {
	// Fibonacci Sequence: 0, 1, 1, 2, 3, 5, 8, 13, ...
	fmt.Println(fib(0)) // 0
	fmt.Println(fib(1)) // 1
	fmt.Println(fib(2)) // 1
	fmt.Println(fib(3)) // 2
	fmt.Println(fib(4)) // 3
	fmt.Println(fib(5)) // 5
	fmt.Println(fib(6)) // 8
	fmt.Println(fib(7)) // 13
}

// fib recursively computes the n-th Fibonacci number
func fib(n int) int {
	// Base cases
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// Recursive case
	return fib(n-1) + fib(n-2)
}
