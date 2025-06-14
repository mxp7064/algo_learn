package main

import "fmt"

/*
=============================================================================
🧩 Fibonacci with Memoization — Top-Down Dynamic Programming (O(n) time)
=============================================================================

Improves the brute-force recursion by storing (memoizing) already-computed values in a map so we never compute the
same value more than once.

🧠 Time Complexity: O(n)
Space Complexity: O(n) for the memoization map and call stack
Thanks to memoization, we're turning an exponential problem into a linear-time one

🔍 Why is memoized recursion considered Dynamic Programming (DP)?
Because it follows the core principle of DP:
✅ Break a problem into overlapping subproblems (same as in brute force)
✅ Store the result of each subproblem (key for being DP)
✅ Reuse those results instead of recomputing them (key for being DP)

This avoids redundant work — the essence of dynamic programming.

📌 Why is this solution considered "Top-Down"?
- We start from the original problem fib(n) and recursively break it down into smaller subproblems: fib(n-1), fib(n-2), ..., fib(0)
- We go from the "top" (the final goal) down to the "bottom" (the base cases)
- This is same as in the brute force approach but since the results of subproblems are cached (memoized) and reused, it is also
considered as DP so we say that this solution if top-down DP!

This contrasts with "bottom-up", where we start from the base cases and build up.

*/

func main() {
	// Fibonacci Sequence: 0, 1, 1, 2, 3, 5, 8, 13, ...
	m := make(map[int]int)
	fmt.Println(fib(0, m)) // 0
	fmt.Println(fib(1, m)) // 1
	fmt.Println(fib(2, m)) // 1
	fmt.Println(fib(3, m)) // 2
	fmt.Println(fib(4, m)) // 3
	fmt.Println(fib(5, m)) // 5
	fmt.Println(fib(6, m)) // 8
	fmt.Println(fib(7, m)) // 13
}

// fib returns the n-th Fibonacci number using top-down memoization
func fib(n int, memo map[int]int) int {
	// If we've already computed it, return cached result
	if v, ok := memo[n]; ok {
		return v
	}

	// Base cases
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// Recursive computation with caching
	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}
