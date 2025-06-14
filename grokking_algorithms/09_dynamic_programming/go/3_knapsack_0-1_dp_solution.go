/*
0/1 Knapsack Problem solved with dynamic programming - finds best solution much faster than brute force
Intuitive explanation can be found in chapter 9 - The knapsack problem, page 161

Problem: Select items with maximum value without exceeding knapsack capacity.
Each item can either be taken fully (1) or not taken (0). No fractions allowed.

Time Complexity: O(n * capacity), where n is the number of items and capacity is the capacity of the knapsack.
This is much better than the brute force solution (check out 08_greedy_algorithms/go/2_knapsack_0-1_brute_force.go)
which is O(2^n)

WHY IS THIS NP-COMPLETE IF IT'S DP?
- The DP solution runs in O(n × capacity), which is pseudo-polynomial.
- capacity = W is just a number, but it only takes log₂(W) bits to represent it.
- So O(n × W) is exponential in input size (number of bits), not truly polynomial in the size of the input.

In complexity theory:
- Input size = number of bits to encode the input
- True polynomial time must be polynomial in the number of bits, not just numeric values
- Therefore, 0/1 Knapsack is still NP-complete
*/

package main

import (
	"fmt"
)

type Item struct {
	Name   string
	Weight int
	Value  int
}

// maxInt is helper method which returns the bigger of the two int numbers
func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// createMatrix creates a slice of int slices which represent a matrix
// all values in the matrix/grid are 0
func createMatrix(width, height int) [][]int {
	matrix := make([][]int, height)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}
	return matrix
}

func main() {
	items := []Item{
		{"A", 3, 4},
		{"B", 2, 3},
		{"C", 4, 5},
	}
	capacity := 6

	n := len(items)
	// we add 1 so we don't have to handle cases when index < 0 in the matrix
	// and also to represent the case when we don't have items or when we have 0 capacity
	// dp[0][0] = 0 and everything will work out
	dp := createMatrix(capacity+1, n+1)
	for i := 1; i <= n; i++ { // O(n)
		item := items[i-1]               // <- careful, dp is 1 ahead of the items index so use i-1 because dp has extra row (we loop 1...n)
		for j := 1; j <= capacity; j++ { // O(capacity)
			previousBest := dp[i-1][j]
			if item.Weight <= j { // item fits
				// If item fits, we can either take it or not → we pick the better
				newBest := item.Value + dp[i-1][j-item.Weight] // dp[i-1][j-item.Weight] represents the max value that we can steal if we make space for this item in the knapsack
				dp[i][j] = maxInt(previousBest, newBest)
			} else { // item too heavy
				// If it doesn’t fit, we just inherit the previous best (without this item)
				dp[i][j] = previousBest
			}
		}
	}

	// Reconstruct which items we need to steal by reversing the logic
	i := n
	j := capacity
	var itemsToSteal []Item
	for i > 0 {
		if dp[i][j] > dp[i-1][j] { // we took the item
			itemsToSteal = append(itemsToSteal, items[i-1])
			j -= items[i-1].Weight // because we made space for it
		}
		i-- // move to previous item in any case
	}
	fmt.Println(itemsToSteal)

	// time complexity: O(n * capacity)
	fmt.Println("Result is:", dp[n][capacity]) // Result is: 8 (we steal items B and C)
}

// Space-optimized 0/1 knapsack using 1D dp array
// To compute row i, you only need row i-1.
// So instead of a 2D dp, we use a 1D slice and update it in place.
func knapsack1D(items []Item, capacity int) int {
	dp := make([]int, capacity+1) // dp[j] = max value for capacity j

	for _, item := range items {
		// Traverse right-to-left to avoid overwriting values from the same row
		for j := capacity; j >= item.Weight; j-- {
			dp[j] = maxInt(dp[j], dp[j-item.Weight]+item.Value)
		}
	}

	return dp[capacity]
}
