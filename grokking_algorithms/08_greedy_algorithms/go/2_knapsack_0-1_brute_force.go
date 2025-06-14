/*
0/1 KNAPSACK PROBLEM – BRUTE FORCE BITMASKING SOLUTION

This solution finds the best/optimal solution but it uses an inefficient brute force approach.
This solution is not found in book.

This is not fractional Knapsack problem where we can fill fractions of items in the knapsack — we can either
include or not include the whole item, hence 0/1 in the name.

PROBLEM:
A thief is robbing a house. His knapsack has limited weight capacity so he must decide
which items to steal to maximize the total value of the loot without exceeding the knapsack weight capacity.

ABSTRACT FORMULATION:
Given a set of items with weights and values, find the subset of items that fits within the knapsack's weight capacity
and maximizes total value.

SOLUTION:
This code uses bitmasking to generate all possible subsets (the power set) of the items.
It evaluates each subset to check if it fits within the knapsack's weight limit and keeps track of
the subset with the highest total value.

IMPORTANT NOTE:
The 0/1 Knapsack is an NP-complete problem.

- Brute-force (checking all subsets) guarantees the best solution but has exponential time complexity O(2^n).
- Dynamic Programming (DP) can also find the exact best solution in O(n × capacity) time,
which is much faster for reasonable capacities.

Checkout 3_knapsack_0-1_dp_solution.go file in 09_dynamic_programming folder for the DP version.

TIME COMPLEXITY:
- O(2^n), where n is the number of items, because all possible subsets are generated and checked.
- Technically it’s O(n × 2^n) because generating each subset also requires checking up to n items,
but in Big O notation we drop the smaller factor (n) since 2^n dominates.

WHY IS THIS NP-COMPLETE IF DP SOLUTION EXISTS?
- The DP solution runs in O(n × capacity), which is pseudo-polynomial.
- capacity = W is just a number, but it only takes log₂(W) bits to represent it.
- So O(n × W) is exponential in input size (number of bits), not truly polynomial in the size of the input.

In complexity theory:
- Input size = number of bits to encode the input
- True polynomial time must be polynomial in the number of bits, not just numeric values
- Therefore, 0/1 Knapsack is still NP-complete
*/

package main

import "fmt"

type Item struct {
	Name   string
	Weight int
	Value  int
}

func main() {
	items := []Item{
		{"A", 3, 4},
		{"B", 4, 5},
		{"C", 2, 3},
	}
	capacity := 6

	n := len(items)
	bestValue := 0
	var bestSubset []Item

	// generate powerset (all subsets) by generating all numbers from 0 to 2^n - 1
	for subset := 0; subset < 1<<n; subset++ { // O(2^n)
		var currentSubset []Item
		totalWeight := 0
		totalValue := 0

		// check each bit position — if i-th bit is set, include item i in the current subset
		for i := 0; i < n; i++ {
			if subset&(1<<i) != 0 {
				item := items[i]
				currentSubset = append(currentSubset, item)
				totalWeight += item.Weight
				totalValue += item.Value
			}
		}

		// if this subset fits and is more valuable, remember it
		if totalWeight <= capacity && totalValue > bestValue {
			bestSubset = currentSubset
			bestValue = totalValue
		}
	}

	// print best subset and its total value
	for _, item := range bestSubset {
		fmt.Printf("item: %s: value=%d, weight=%d\n", item.Name, item.Value, item.Weight)
	}
	fmt.Println("Total value:", bestValue)
	// item: B: value=5, weight=4
	// item: C: value=3, weight=2
	// Total value: 8
}
