package main

import "fmt"

/*
Subsets (Power Set) Generation Using Backtracking

Note that this is almost the same as previous example (choosing dish combinations), it's just written
in more abstract way and we generate all possible subsets.

Problem:
Given a list of numbers (e.g. [1, 2, 3]), generate all possible subsets (the power set).
This includes:
- The empty subset []
- All combinations of elements without duplicates
- Order of elements in a subset does not matter

Backtracking Approach:
- At each step, make a decision: include the current number or skip it
- Recurse to explore deeper decisions
- Backtrack (undo the last choice) to explore alternate paths

This problem follows the exact same recursive pattern as the dish menu example:
1. Make a decision
2. Recurse
3. Backtrack (undo decision)

Decision Tree Example for input: [1, 2]

                     []
                   /    \
                 [1]     []
                /   \      \
            [1,2]   [1]    [2]

Each level represents a choice to include or exclude the next number.

Time Complexity:
- There are 2^n subsets for a list of length n → O(2^n)
- For each subset we may copy up to n elements → O(n)
	- when we do the array copy
- Total complexity: O(n × 2^n)

Space Complexity:
- O(n) recursion stack (maximum depth)
- O(n × 2^n) result storage

🧠 Why use backtracking for generating subsets when bitmasking exists?

Remember this example where we used bitmasking to generate all binary numbers up to 2^n to generate the power set:
08_greedy_algorithms/go/2_knapsack_0-1_brute_force.go

Bitmasking is a fast way to generate all 2^n subsets — perfect for when:
- You just need all subsets (no constraints)
- You don’t care about order or optimization

But backtracking gives more power and flexibility:
- Can handle constraints (e.g., subset sum, fixed length, no duplicates)
- Can prune branches early (e.g., when a partial subset already violates a rule)
- Supports reuse of elements or skipping based on index/used map
- Easier to extend into problems like permutations, combination sum, etc.

Use bitmasking when:
- You just want the power set
- Performance is critical and constraints are minimal

Use backtracking when:
- You have constraints or need custom logic (like pruning, skipping, reuse)
- This is what most LeetCode subset generation problems revolve around

🔑 LeetCode issues:
- LC 78. Subsets
	- Generate all subsets (power set). (start index DFS or bitmasking)
- LC 90. Subsets II
	- Like LC78 but with duplicate elements → needs sorting + skip duplicates
- LC 131. Palindrome Partitioning
	- Split string into all palindromic partitions (DFS + path building)

*/

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("All subsets (power set):")

	var result [][]int
	var subset []int

	var backtrack func(start int)
	backtrack = func(start int) {
		// Print the current subset at this level
		fmt.Printf("Subset so far: %v\n", subset)

		// Add a copy of the current subset to the result
		temp := make([]int, len(subset))
		copy(temp, subset)
		result = append(result, temp)

		// Try adding each remaining number starting from index 'start'
		for i := start; i < len(nums); i++ {
			fmt.Printf("  Including %d\n", nums[i])
			subset = append(subset, nums[i]) // Make a decision
			backtrack(i + 1)                 // Recurse
			fmt.Printf("  Backtracking, removing %d\n", subset[len(subset)-1])
			subset = subset[:len(subset)-1] // Undo the decision (backtrack)
		}
	}

	backtrack(0)

	// Print all resulting subsets
	for _, s := range result {
		fmt.Println(s)
	}
}
