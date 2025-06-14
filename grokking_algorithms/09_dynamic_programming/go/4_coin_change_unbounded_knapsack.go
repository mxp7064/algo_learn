/*
Coin Change Problem (Dynamic Programming)
Not found in book.

Problem Statement:
Given an array of coin denominations and a target amount,
return the minimum number of coins needed to make that amount.
If it's not possible to make the amount, return -1.

You can use each coin an unlimited number of times (unbounded) so this problem is
like "unbounded knapsack" problem

Brute-force approach: For each amount from 0 to target amount, recursively try every coin denomination,
allowing repeated use of the same coin. This explores all combinations with replacement.
Time complexity: O(k^n), where k is the number of coin types and n is the target amount,
since at each step we branch into k recursive calls until we reach amount 0 or below.

DP approach: Bottom-up DP using 1D dp array.
Real-World Analogy: ATM trying to give you exact money using the smallest number of bills/coins.
Limitations: Works only when coins have positive values and amount is ≥ 0.

Time Complexity: O(n * amount), where n is the number of coin denominations.
Space Complexity: O(amount + 1), for the dp array.
*/

package main

import "fmt"

func minCoins(coins []int, amount int) int {
	// dp[i] will hold the minimum number of coins needed to make amount i
	// We initialize with a large value (amount + 1), which is effectively "infinity/unreachable" here.
	dp := make([]int, amount+1) // 1D slice
	for i := range dp {
		dp[i] = amount + 1
	}

	// Base case: 0 coins are needed to make amount 0, very important so we can start building bottom up
	dp[0] = 0

	// For each coin, and each sub-amount, calculate minimum number of coins to make up that amount
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			// If we use this coin, do we get a better (smaller) result?
			// dp[i-coin]+1 represents the number of coins needed if we use this coin
			// We are trying to get the best result for a case where we don’t take this coin yet (dp[i-coin]), and
			// then calculate the result if we do take it (dp[i-coin]+1) and then compare it with previous best
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	// If we still have our "infinity/unreachable" placeholder value at dp[amount], we couldn’t make the amount
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// Helper function to get the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test Case 1
	coins1 := []int{1, 2, 5}
	amount1 := 11
	fmt.Printf("Minimum coins to make %d from %v = %d\n", amount1, coins1, minCoins(coins1, amount1))
	// Output: 3 (5 + 5 + 1)

	// Test Case 2
	coins2 := []int{2}
	amount2 := 3
	fmt.Printf("Minimum coins to make %d from %v = %d\n", amount2, coins2, minCoins(coins2, amount2))
	// Output: -1 (not possible)

	// Test Case 3
	coins3 := []int{1}
	amount3 := 0
	fmt.Printf("Minimum coins to make %d from %v = %d\n", amount3, coins3, minCoins(coins3, amount3))
	// Output: 0 (no coins needed)

	// Test Case 4
	coins4 := []int{1, 3, 4}
	amount4 := 6
	fmt.Printf("Minimum coins to make %d from %v = %d\n", amount4, coins4, minCoins(coins4, amount4))
	// Output: 2 (3 + 3)
}
