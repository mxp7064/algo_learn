package main

import "fmt"

/*
LONGEST COMMON SUBSTRING – DYNAMIC PROGRAMMING

🧠 PROBLEM:
Given two strings, find the longest substring (continuous block of characters) that appears in both strings.

⚠️ NOTE:
- This is NOT the same as Longest Common Subsequence (which allows skipping characters).
- We need a contiguous match — once characters don't match, we reset the count.

📌 EXAMPLE:
wordA = "Mario"
wordB = "Marko"
-> Longest common substring: "Mar" (length 3)
-> Longest common subsequence: "Maro" (length 4)

🧠 DP INSIGHT:
Dynamic Programming is a strategy to solve problems by breaking them down into smaller subproblems,
solving each one only once, and building up the solution from the bottom up.

In this case, we use a grid (2D table) to record the length of the longest common substring ending at
each pair of positions in the two words. We build up partial results (substrings) and use them
to compute larger results — reusing work instead of recalculating it from scratch.

🧩 APPROACH:
Use dynamic programming.
- Let dp[i][j] = length of the longest common substring ending at wordA[i-1], wordB[j-1]
- Transition:
    if wordA[i-1] == wordB[j-1]:
        dp[i][j] = dp[i-1][j-1] + 1
    else:
        dp[i][j] = 0
- Track the maximum value and its position to reconstruct the result.

🧮 TIME COMPLEXITY: O(m * n)
🧮 SPACE COMPLEXITY: O(m * n)

💡 OPTIONAL:
We can optimize space to O(n) using two rows (current and previous because that's all that we need), but then substring
reconstruction requires more logic or backtracking the actual input strings.
*/

// createMatrix creates a slice of int slices (a grid) initialized to 0 (all values in the matrix/grid are 0)
// Useful for dynamic programming problems like this one.
func createMatrix(width, height int) [][]int {
	matrix := make([][]int, height)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}
	return matrix
}

func main() {
	wordA := "Mario"
	wordB := "Marko"

	// Classic DP way using endIndex tracking
	lcss := longestCommonSubstring(wordA, wordB)
	fmt.Printf("Longest common substring between '%s' and '%s' is '%s' (length %d)\n", wordA, wordB, lcss, len(lcss))

	// Alternative way: reconstructing diagonally from max position
	lcssAlt := longestCommonSubstringOtherWay(wordA, wordB)
	fmt.Printf("Alternative reconstruction – longest common substring is '%s' (length %d)\n", lcssAlt, len(lcssAlt))
}

// longestCommonSubstring returns the longest common substring using endIndex tracking.
// This version is simple and efficient for both finding length and extracting the result.
func longestCommonSubstring(wordA, wordB string) string {
	m := len(wordA)
	n := len(wordB)
	dp := createMatrix(n+1, m+1)

	maxLength := 0
	endIndex := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if wordA[i-1] == wordB[j-1] { // adjust the index, dp is 1 ahead of string index
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
					endIndex = i // or we can track the end index j of wordB and later extract the substring from wordB
				}
			} else {
				dp[i][j] = 0 // optional, as matrix is already zeroed, but illustrates intent
			}
		}
	}

	return wordA[endIndex-maxLength : endIndex]
}

// longestCommonSubstringOtherWay demonstrates an alternative way to reconstruct the substring
// by tracing back diagonally from the maximum value in the DP matrix.
func longestCommonSubstringOtherWay(wordA, wordB string) string {
	m := len(wordA)
	n := len(wordB)
	dp := createMatrix(n+1, m+1)

	maxLength := 0
	maxI, maxJ := 0, 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if wordA[i-1] == wordB[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
					maxI, maxJ = i, j
				}
			}
		}
	}

	// Backtrack diagonally up-left to reconstruct substring
	substr := ""
	for dp[maxI][maxJ] != 0 {
		substr = string(wordA[maxI-1]) + substr
		maxI--
		maxJ--
	}

	return substr
}
