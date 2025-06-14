/*
Combination Sum using backtracking

🧠 Problem:

Given a list of positive integers `candidates` and a target number,
find all unique combinations of candidates where the chosen numbers sum to the target.
You can use each number in `candidates` an unlimited number of times.

Example:

Input:  candidates = [2,3,6,7], target = 7
Output: [[2,2,3], [7]]

📌 Key Backtracking Concepts:

- At each step, you decide whether to include a candidate (possibly multiple times)
- Recurse to explore further additions
- Backtrack by removing the last added number
- You are allowed to reuse the same number (so we don't increment the index when recursing)

🧠 INTUITION: Combination Sum is a subset generation problem with constraints

This is an extension of the standard subset generation via backtracking:
- Instead of collecting all subsets (or subsets of certain length), we only collect those where the sum equals `target`
- We build all valid combinations by exploring numbers recursively, like in subset generation
- We also use the `start` index, but we are allowed to reuse elements (unlimited times)
- So we don't increment `start` after each pick — this allows `[2,2,3]`, etc.

The only major additions:
- A `sum` parameter that tracks the current path sum
- A base case that checks `sum == target`
- A pruning case if `sum > target` to stop invalid paths early

So it’s essentially:
  - Subset generation using DFS
  - But with an accumulator (`sum`) that decides if a path is valid
  - And recursion continues on same index to allow reuse

🔁 KEY LOGIC STRUCTURE:

	func dfs(start, sum):
	    if sum == target:
	        save path
	        return
	    if sum > target:
	        return
	    for i := start to end:
	        choose candidates[i]
	        dfs(i, sum + candidates[i])   // reuse allowed
	        unchoose candidates[i]

🧠 Note on similarity to Coin Change problem (09_dynamic_programming/go/4_coin_change_unbounded_knapsack.go):
This is conceptually the same as the Coin Change problem:
- Candidates = coin denominations
- Target = amount
- We want to find all unique combinations that sum up to the target
- Reuse of elements is allowed
- Here we collect all combinations that sum up to the target, while in Coin Change problem we just need the length of smallest combination/set

Here we are just using recursion/backtracking/DFS instead of bottom-up DP.

🧠 Why Use Backtracking (DFS) vs. DP?

Backtracking (DFS):
- Explores all possible choices recursively
- Great when you need to list/generate all valid combinations, permutations, or paths
- Useful when decisions are sequential and constraints are complex
- Easier to implement when building the actual solutions (e.g., return all combinations)

Dynamic Programming (DP):
- Optimizes overlapping subproblems using memoization or tabulation
- Great for computing a count, max/min value, or yes/no answer
- Efficient for problems with optimal substructure (e.g., longest subsequence, max profit)
- Often not used when you need to return all paths/combinations (though it's possible, it's messy)

Use backtracking when:
- You need to generate all valid results
- Problem involves combinatorial exploration
- You can prune invalid paths early

Use DP when:
- You need to count, maximize, or decide
- There’s overlap in subproblems
- You don’t need to generate the full result list

⏱ Time Complexity:

- Each number can be picked multiple times
- Each recursive call can try each candidate again
- Time: O(2^target × k), where k = avg combination length
	- O(2^T) → exponential number of recursive paths
	- × k → because each valid combination has length up to k (and copying costs O(k))
- So it's exponential in worst case, depending on number and size of candidates

Example, let's define:
- T = target value
- k = average length of one valid combination

Worst-case arises when:
- We have small candidates like [1, 2, 3]
- And a large target like 20
- Since we can reuse candidates, the algorithm explores many combinations
  (e.g., [1,1,1,...], [2,2,1,...], etc.)

For example:
candidates = [1]
target = 10

- Valid solution: [1,1,1,1,1,1,1,1,1,1]
- But recursion will still explore every possible combination of 1s summing to ≤ 10
- Recursion depth: up to 10
- Each level: always branches back into adding more 1s → like a linear path
- So roughly we have O(2^10)

Bigger branching example:
candidates = [1, 2]
target = 4

Valid combinations:
[1,1,1,1]
[1,1,2]
[1,2,1]
[2,1,1]
[2,2]

- Recursion tree becomes exponential due to multiple candidates and reuse
- Each branch leads to multiple sub-branches → branching factor grows

Why O(2^T × k)?
- 2^T: Exponential number of recursive calls (combinations to reach target)
- × k: Each valid result can contain up to k elements → copying into final result costs O(k)

Summary:
- This complexity is acceptable for small targets
- But when target is large and small candidates are allowed → exponential growth happens
- Pruning (early return if sum > target) helps avoid exploring invalid branches, but doesn't change worst-case

📦 Space Complexity:
- Stack depth up to target / min(candidate), plus result storage

🔑 LeetCode issues (combination sum and variations):

- LC 39. Combination Sum
	- Unbounded pick of candidates (this one)
- LC 40. Combination Sum II
	- Single-use elements, requires skipping duplicates
- LC 377. Combination Sum IV
	- Count combinations → not backtracking, use DP instead
- LC 216. Combination Sum III
	- Pick k numbers that sum to n from 1 to 9

*/

package main

import "fmt"

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7

	fmt.Printf("Combinations that sum to %d:\n", target)
	result := combinationSum(candidates, target)
	for _, combo := range result {
		fmt.Println(combo)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var path []int

	// Recursive backtracking function
	var backtrack func(start int, sum int)
	backtrack = func(start int, sum int) {
		// Base case: if sum equals target, store a copy of the current path
		if sum == target {
			// Make a copy before appending
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		// If sum exceeds target, stop exploring this path (prune)
		if sum > target {
			return
		}

		// Try all candidates starting from current index
		for i := start; i < len(candidates); i++ {
			num := candidates[i]
			path = append(path, num)  // Choose the number
			backtrack(i, sum+num)     // Recurse with same index (reuse allowed)
			path = path[:len(path)-1] // Backtrack: remove last number
		}
	}

	backtrack(0, 0)
	return result
}
