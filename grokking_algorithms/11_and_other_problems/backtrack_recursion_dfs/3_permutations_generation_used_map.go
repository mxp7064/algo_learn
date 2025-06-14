package main

import "fmt"

/*
SUBSETS vs PERMUTATIONS – LENGTH K VARIANTS

We demonstrate how to generate:
1. Subsets of length k (combinations – order doesn’t matter, no repeats)
2. Permutations of length k (order matters, no repeats)

These two classic DFS/backtracking problems use slightly different techniques:
- Subsets: use `start` index to avoid duplicates and reorderings
- Permutations: use `used` map to prevent selecting the same element twice along one path

Time Complexity:

Subsets:
- We explore combinations of length up to k
- At each level we make O(n) choices, pruning using `start`
- Total time: O(n^k) in worst case
- Space: O(k) recursion + O(#results × k) for output

Permutations:
- We generate P(n, k) = n! / (n - k)! permutations
- Each permutation takes O(k) time to build/copy
- Total time: O(k × P(n, k)) = O(k × n! / (n - k)!)
- Space: O(k) recursion stack + O(P(n, k) × k) result

🔑 LeetCode issues:
- LC 46. Permutations
	- All permutations of distinct numbers (used map OR in-place swap)
- LC 47. Permutations II
	- Like LC46 but with duplicates → sort + used + skip logic
*/

func main() {
	input := []string{"A", "B", "C"}
	k := 2

	fmt.Println("Subsets of length", k)
	for _, combo := range getSubsetsOfLengthK(input, k) {
		fmt.Println(combo)
	}

	fmt.Println("\nPermutations of length", k)
	for _, perm := range getPermutationsOfLengthK(input, k) {
		fmt.Println(perm)
	}
}

// This is our standard generalized and simplified way to generate subsets (as in 1_choosing_dish_combinations.go and 2_subset_generation.go)
// We rely on the `start` index to prevent duplicates and permutations (e.g. avoid both [A,B] and [B,A]) along the recursive path.
func getSubsetsOfLengthK(input []string, k int) [][]string {
	var result [][]string

	var dfs func(start int, path []string)
	dfs = func(start int, path []string) {
		if len(path) == k {
			result = append(result, append([]string{}, path...))
			return
		}
		for i := start; i < len(input); i++ {
			path = append(path, input[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, nil)
	return result
}

// This is a typical way to generate permutations using a `used` map.
// The `used` map prevents reusing the same element along the current recursive path,
// but it allows reusing previously used elements on other paths because we reset the flag after the recursive call.
// Time complexity: O(n! / (n-k)!) because we are generating permutations of k out of n.
func getPermutationsOfLengthK(input []string, k int) [][]string {
	var result [][]string

	var dfs func(used map[string]bool, path []string)
	dfs = func(used map[string]bool, path []string) {
		if len(path) == k {
			result = append(result, append([]string{}, path...))
			return
		}
		for i := 0; i < len(input); i++ {
			if used[input[i]] {
				continue
			}
			used[input[i]] = true
			path = append(path, input[i])

			dfs(used, path)

			path = path[:len(path)-1]
			used[input[i]] = false
		}
	}

	dfs(make(map[string]bool), nil)
	return result
}
