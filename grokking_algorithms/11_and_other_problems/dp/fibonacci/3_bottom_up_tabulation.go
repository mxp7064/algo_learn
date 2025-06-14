package main

import "fmt"

/*
===================================================================================================
🧠 Fibonacci — Bottom-Up Dynamic Programming (Tabulation, O(n) Time and Space)
===================================================================================================

Approach:
- This is the classic bottom-up dynamic programming (DP) approach.
- Instead of using recursion, we use iteration to build the solution from base cases (fib(0), fib(1)) up.
- We store intermediate results (results of subproblems) in a "table" (a slice) — hence the term tabulation.

📌 Why is it called "Tabulation"?
- We create a table (the dp array) to store solutions to all subproblems.
- We fill the table from smallest to largest, from fib(0) to fib(n).
- This is the opposite of "top-down" (which starts with fib(n) and breaks down recursively).

📌 Why use dp array of size n+1?
- We need to store results for all values from 0 up to n (inclusive), so the array must have n+1 entries:
  dp[0], dp[1], ..., dp[n]

📊 Comparison with other approaches:

🔴 Brute-force recursion:
- Exponential time due to repeated computation
- Time: O(2^n)
- Space: O(n) (call stack)

🟡 Top-down with memoization:
- Recursive with caching
- Time: O(n)
- Space: O(n) (memo + call stack)

🟢 Bottom-up tabulation (this version):
- Iterative, no recursion
- Time: O(n)
- Space: O(n) (dp array)

We’ll later optimize space to O(1) by using just two variables in next example: 4_bottom_up_space_optimized.go
*/

func main() {
	// Fibonacci sequence: 0, 1, 1, 2, 3, 5, 8, 13, ...
	fmt.Println(fib(0)) // 0
	fmt.Println(fib(1)) // 1
	fmt.Println(fib(2)) // 1
	fmt.Println(fib(3)) // 2
	fmt.Println(fib(4)) // 3
	fmt.Println(fib(5)) // 5
	fmt.Println(fib(6)) // 8
	fmt.Println(fib(7)) // 13
}

// fib returns the n-th Fibonacci number using bottom-up tabulation
func fib(n int) int {
	// Handle base cases directly
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// Create a dp array of size n+1 to store fib(0) through fib(n)
	dp := make([]int, n+1)

	// Initialize base cases
	dp[0] = 0
	dp[1] = 1

	// Fill dp[i] = dp[i-1] + dp[i-2] for all i from 2 to n
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	// Return the result for fib(n)
	return dp[n]
}
