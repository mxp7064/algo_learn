package main

import "fmt"

/*
===========================================================================================
Longest Palindromic Subsequence — Bottom-Up Dynamic Programming (O(n²) Time, O(n²) Space)
===========================================================================================

Problem:
Given a string, find the length of the longest subsequence that is also a palindrome.

Definitions:
- A subsequence is a sequence that appears in the same relative order, but not necessarily contiguous.
- A palindrome reads the same forward and backward.

Approach:
- Define dp[i][j] as the length of the longest palindromic subsequence in the substring s[i..j] (inclusive of both i and j)

Transition rules:
- If s[i] == s[j], then dp[i][j] = dp[i+1][j-1] + 2
- Otherwise, dp[i][j] = max(dp[i+1][j], dp[i][j-1])

Why i goes from end to start:
- dp[i][j] depends on dp[i+1][j-1], dp[i+1][j], and dp[i][j-1]
- To ensure these values are already computed when we fill dp[i][j], we must fill the table
  from the bottom row up, meaning i must go from n-1 down to 0

Why j starts at i + 1:
- We only care about substrings where j > i (length 2 or more)
- The base case dp[i][i] is already initialized to 1

Base case:
- All single characters are palindromes, so dp[i][i] = 1

Time Complexity:  O(n²)
Space Complexity: O(n²)
*/

func main() {
	fmt.Println(lps("bbab"))         // Output: 3 ("bbb")
	fmt.Println(lps("bbbab"))        // Output: 4 ("bbbb")
	fmt.Println(lps("cbbd"))         // Output: 2 ("bb")
	fmt.Println(lps("abcd"))         // Output: 1
	fmt.Println(lps("a"))            // Output: 1
	fmt.Println(lps("abacdfgdcaba")) // Output: 11 ("abacdgdcaba")
}

// maxInt returns the greater of a and b
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// lps returns the length of the longest palindromic subsequence in the input string
func lps(s string) int {
	n := len(s)

	// dp[i][j] represents the length of the longest palindromic subsequence in s[i..j]
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1 // Base case: each single character is a palindrome of length 1
	}

	// Fill the dp table in bottom-up fashion
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				// If characters match, extend the known palindrome within the bounds
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				// Otherwise, take the maximum excluding either the left or the right character
				dp[i][j] = maxInt(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	// Return the length of the LPS in the full string s[0..n-1]
	return dp[0][n-1]
}
