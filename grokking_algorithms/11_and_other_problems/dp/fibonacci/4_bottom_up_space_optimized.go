package main

import "fmt"

/*
============================================================================================
🧠 Fibonacci — Space-Optimized Bottom-Up Dynamic Programming (O(n) Time, O(1) Space)
============================================================================================

Approach:
- This version is still bottom-up because we start from base cases (fib(0), fib(1))
  and build up to fib(n), just like in tabulation (previous example)
- The key difference is that we only store the last two results (fib(n-1), fib(n-2)),
  instead of an entire array of all subproblem results.
- This gives us O(1) space while preserving O(n) time.

🪜 Comparison with Previous Versions:

1. ❌ Brute-force Recursion:
   - Recomputes subproblems many times
   - Time: O(2^n), Space: O(n) call stack
   - Not DP

2. ✅ Top-Down DP (Memoization):
   - Uses recursion + map to cache results
   - Time: O(n), Space: O(n) for memo + call stack

3. ✅ Bottom-Up DP (Tabulation):
   - Uses array to build from base up
   - Time: O(n), Space: O(n)

4. ✅ Bottom-Up Space-Optimized (this):
   - Stores only last two values
   - Time: O(n), Space: O(1)

*/

func main() {
	fmt.Println(fib(0)) // 0
	fmt.Println(fib(1)) // 1
	fmt.Println(fib(2)) // 1
	fmt.Println(fib(3)) // 2
	fmt.Println(fib(4)) // 3
	fmt.Println(fib(5)) // 5
	fmt.Println(fib(6)) // 8
	fmt.Println(fib(7)) // 13
}

// fib returns the n-th Fibonacci number using bottom-up DP with O(1) space
func fib(n int) int {
	// Handle base cases explicitly
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// Initialize the two previous Fibonacci values
	prev2 := 0 // fib(n-2), starts as fib(0) = 0
	prev1 := 1 // fib(n-1), starts as fib(1) = 1
	var current int

	// Build from fib(2) to fib(n) using only the last two values
	for i := 2; i <= n; i++ {
		current = prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return current
}
