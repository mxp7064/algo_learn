/*
Fractional Knapsack Problem – Greedy Algorithm (Optimal Solution)
Not found in the book

Problem statement:
You are given a set of items, each with a weight and a value.
You are a thief trying to steal those items in a house.
You have a knapsack with a maximum weight capacity.
You can take fractions of items (e.g. 30% of a gold bar) - in contrast to the 0/1 Knapsack Problem
You cannot take more than one of any item (same as in 0/1 Knapsack).

Goal: Maximize the total value of items in the knapsack.

Key idea: Use a greedy strategy by always taking the item with the highest
value-to-weight ratio first. If it doesn’t fit, take as much of the item as possible (fraction of item).

Time Complexity:
- Sorting by value-to-weight ratio: O(n log n)
- Single pass through items: O(n)
→ Total: O(n log n)

Space Complexity:
- O(1) -> no significant additional space used beyond inputs, we just need a couple of variables (remainingCapacity, totalValue...)

This problem is not NP-complete — it can be solved optimally in polynomial time.

Why does greedy work for fractional knapsack but not for 0/1 knapsack to find the optimal solution?

In the fractional knapsack problem, we are allowed to take fractions of items.
This makes the greedy approach optimal: we always pick the item with the highest
value-to-weight ratio, and if it doesn’t fully fit, we take the exact fraction that does.
This way, we never "waste" capacity — every remaining space is filled with the best
possible value. The greedy choice at each step leads to a globally optimal solution.

In contrast, the 0/1 knapsack problem does not allow fractions. We must take
either the whole item or skip it entirely. In this case, greedy strategies fail because
some combinations of lower-ratio items can yield a higher total value than a single
high-ratio item. So the local best choice doesn’t guarantee a global best result.

→ Fractional knapsack: greedy works perfectly → O(n log n)
→ 0/1 knapsack: greedy is not reliable → needs brute-force or DP → O(n × capacity)
*/
package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Name   string
	Weight float64
	Value  float64
}

func main() {
	capacity := 5.0

	items := []Item{
		{"flour", 1.5, 23.56},
		{"soda", 3.4, 45.6},
		{"cake", 0.54, 56},
	}

	remainingCapacity := capacity
	totalValue := 0.0

	// Step 1: Sort items by descending value-to-weight ratio
	sort.Slice(items, func(i, j int) bool {
		return items[i].Value/items[i].Weight > items[j].Value/items[j].Weight
	})

	// Step 2: Greedily select items while there's remaining capacity
	// Items at beginning have largest value-to-weight ratios so we take those first - greedy strategy
	for _, item := range items {
		if item.Weight <= remainingCapacity { // item can fit in the knapsack
			// Take the entire item
			remainingCapacity -= item.Weight
			totalValue += item.Value
			fmt.Printf("%s 100%%\n", item.Name)
		} else { // item can't fit into the knapsack
			// Take a fraction of the item
			fraction := remainingCapacity / item.Weight // "spread" remainingCapacity over item.Weight
			totalValue += item.Value * fraction
			fmt.Printf("%s %.2f%%\n", item.Name, fraction*100)
			remainingCapacity -= item.Weight * fraction // will be 0, but we do it for clarity
			break                                       // knapsack is full
		}
	}

	// Step 3: Print the total value obtained
	fmt.Printf("total value stolen: %.2f\n", totalValue)
	fmt.Printf("remaining capacity (should be 0): %f", remainingCapacity)
}
