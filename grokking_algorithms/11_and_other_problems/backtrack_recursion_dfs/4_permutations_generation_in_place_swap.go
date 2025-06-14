package main

import "fmt"

/*
PERMUTATION GENERATION (In-Place Swapping Method)

🧠 PROBLEM:
Given an array of distinct integers (e.g. [1, 2, 3]), generate all possible permutations.

We want to generate all arrangements where the order of elements matters (e.g., [1,2,3] ≠ [2,1,3]).

----------------------------------------
🔍 APPROACH: In-Place Swapping
----------------------------------------
We use backtracking by recursively fixing one position at a time (index `from`), and generating permutations
of the remaining suffix by swapping the current index with all possible choices.

Why it works:
- Each recursion level represents a fixed index (`from`) in the output.
- At each level, we try all possible elements that could be placed in that index.
- We swap the candidate into position `from`, recurse, and then swap back ("backtrack").

This mutates the original array during recursion, but restores it to its original form before exiting each level.

----------------------------------------
🌳 VISUAL DECISION TREE for input: [A, B, C]
----------------------------------------

                []
           /      |      \
         [A]     [B]     [C]
        /   \    /  \    /  \
     [A B] [A C] ...     ...
     /         \
 [A B C]     [A C B]     ...

Each path to a leaf node gives one complete permutation.

----------------------------------------
🚨 MUTATION & BACKTRACKING EXPLAINED
----------------------------------------
We swap to simulate a choice (mutating array in-place), then backtrack by undoing the swap.
This ensures we don’t accidentally corrupt the current branch’s state before returning.

Example:
1. Fix position 0:
   - Try A, B, C → swap A/B/C to front
2. For each, fix position 1 by swapping remaining elements
3. Recurse to position 2, when base case hit, copy and store result
4. Swap back to undo

----------------------------------------
📌 COMPARISON: Used Map vs In-Place
----------------------------------------

| Approach     | Space   | Speed | Notes |
|--------------|---------|-------|-------|
| Used Map     | Extra map[string]bool | More overhead | Easy to extend, no mutation |
| In-Place     | No extra structures   | Faster (fewer allocations) | Mutates input, must backtrack correctly |

In-place method is more space efficient and elegant when mutation is safe and acceptable.

----------------------------------------
⏱ TIME COMPLEXITY:
- There are n! permutations of length n → O(n!)
- Each permutation takes O(n) to copy → O(n × n!) total time

📦 SPACE COMPLEXITY:
- O(n) recursion stack (depth of call stack)
- O(n × n!) result storage (all permutations)

*/

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("All permutations (in-place swap method):")

	results := permute(nums)

	for _, p := range results {
		fmt.Println(p)
	}
}

// permute generates all permutations using in-place swap + backtracking
func permute(nums []int) [][]int {
	var results [][]int

	var backtrack func(from int)
	backtrack = func(from int) {
		// Base case: all positions are fixed
		if from == len(nums) {
			// Make a deep copy of current permutation before saving
			perm := make([]int, len(nums))
			copy(perm, nums)
			results = append(results, perm)
			return
		}

		// Try placing each possible number in index `from`
		for i := from; i < len(nums); i++ {
			// Make a choice
			nums[from], nums[i] = nums[i], nums[from]

			// Recurse with this choice fixed
			backtrack(from + 1)

			// Undo the choice (backtrack)
			nums[from], nums[i] = nums[i], nums[from]
		}
	}

	backtrack(0)
	return results
}
