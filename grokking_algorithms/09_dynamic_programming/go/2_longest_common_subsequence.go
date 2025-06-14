package main

import (
	"fmt"
	"slices"
)

/*
LONGEST COMMON SUBSEQUENCE (LCS) – DYNAMIC PROGRAMMING

Contains variations of the LCS problem which are frequently asked in interviews.

🧠 PROBLEM:
Given two strings, return the length of the longest subsequence present in both of them.
A subsequence is a sequence of characters that appear in the same relative order, but not necessarily contiguous.

📌 EXAMPLE:
wordA = "fosh"
wordB = "fish"
-> Longest common subsequence: "fsh" (length 3)

📌 ALSO SEE:
See 1_longest_common_substring.go for a simpler related problem — that one requires the matching characters to be contiguous,
whereas LCS allows skipping characters as long as order is preserved.

🧠 DP INSIGHT:
We're building a grid where each cell represents the length of the LCS between the first `i` characters of wordA and the first `j` characters of wordB.

If the characters at wordA[i-1] and wordB[j-1] match:
→ That means the LCS can be extended, so we set:
    dp[i][j] = dp[i-1][j-1] + 1

If they don't match:
→ We are preserving the best LCS so far by taking the max of:
	- the value above (dp[i-1][j]) → best LCS when skipping current char in wordA
    - the value to the left (dp[i][j-1]) → best LCS when skipping current char in wordB

This means we’re always keeping the best possible LCS so far in the grid, so when a match happens, we can grow from the right base.

🧮 TIME COMPLEXITY: O(m * n)
🧮 SPACE COMPLEXITY: O(m * n)
*/

// maxInt returns the maximum of two integers
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// lcs returns the length of the longest common subsequence between two strings
func lcs(wordA, wordB string) int {
	m := len(wordA)
	n := len(wordB)

	// Create a 2D slice initialized to 0
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill the table using the bottom-up DP approach
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if wordA[i-1] == wordB[j-1] {
				// Characters match → extend LCS
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// Characters don't match → preserve best LCS so far
				dp[i][j] = maxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n] // Result is in the bottom-right corner
}

// this variation reconstructs the longest common subsequence by reversing the logic
func lcsReconstruct(wordA, wordB string) string {
	// Same code as before
	m := len(wordA)
	n := len(wordB)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if wordA[i-1] == wordB[j-1] {
				// Characters match → extend LCS
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// Characters don't match → preserve best LCS so far
				dp[i][j] = maxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Reconstruct the LCS by going backwards from the end (we are reversing the logic which builds LCS)
	i, j := m, n
	var result []rune
	for i > 0 && j > 0 {
		if wordA[i-1] == wordB[j-1] {
			result = append(result, rune(wordA[i-1]))
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	slices.Reverse(result)
	return string(result)
}

// lcsNumOfTransformations prints the minimum number of insertions and deletions needed to transform wordA into wordB
// Deletions: characters in A not in LCS
// Insertions: characters in B not in LCS
func lcsNumOfTransformations(wordA, wordB string) {
	lcsLen := lcs(wordA, wordB)
	deletions := len(wordA) - lcsLen
	insertions := len(wordB) - lcsLen
	fmt.Printf("To turn %q into %q, we need %d deletions and %d insertions.\n", wordA, wordB, deletions, insertions)
}

// scs returns the shortest common supersequence (SCS) of wordA and wordB
// This is the shortest string that contains both wordA and wordB as subsequences
// We build the SCS by combining characters from both strings, but we only include the LCS characters once - giving us
// the shortest common supersequence
// Uses LCS as a guide to interleave unmatched characters between LCS points
// Length formula: lenA + lenB - len(LCS)
func scs(wordA, wordB string) string {
	lcs := lcsReconstruct(wordA, wordB)
	i, j := 0, 0
	var result []rune
	for _, c := range lcs {
		// take all characters from A up until this LCS character
		for i < len(wordA) && rune(wordA[i]) != c {
			result = append(result, rune(wordA[i]))
			i++
		}

		// take all characters from B up until this LCS character
		for j < len(wordB) && rune(wordB[j]) != c {
			result = append(result, rune(wordB[j]))
			j++
		}

		// always take LCS character (once)
		result = append(result, c)
		i++
		j++
	}

	// we can have more characters in A or B after LCS
	result = append(result, []rune(wordA[i:])...)
	result = append(result, []rune(wordB[j:])...)

	return string(result)
}

func main() {
	wordA := "fosh"
	wordB := "fish"

	fmt.Printf("LCS length of %q and %q: %d\n", wordA, wordB, lcs(wordA, wordB))
	fmt.Printf("LCS of %q and %q: %q\n", wordA, wordB, lcsReconstruct(wordA, wordB))
	lcsNumOfTransformations(wordA, wordB)

	wordA = "abac"
	wordB = "cab"
	fmt.Printf("Shortest common supersequence of %q and %q: %q\n", wordA, wordB, scs(wordA, wordB))
}
