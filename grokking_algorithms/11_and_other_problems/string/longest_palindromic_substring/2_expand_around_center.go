package main

import "fmt"

/*
===================================================================================
🧠 Longest Palindromic Substring — Expand Around Center (Optimized O(n²) Solution)
===================================================================================

🔷 PROBLEM:
Given a string, return the longest substring that is a palindrome.

🔷 STRATEGY:
Instead of checking all possible substrings (brute force), we treat every index and index pair
as the center of a potential palindrome, and expand outward to find the longest one.

✅ We expand in two ways for each index:
1. Odd-length palindromes:     expandFromCenter(i, i)
   → Example: "aba" → center at 'b'
2. Even-length palindromes:    expandFromCenter(i, i+1)
   → Example: "abba" → center between 'b' and 'b'

Without checking both, we would miss one type.

🔷 WHY THIS IS BETTER THAN BRUTE FORCE:

Brute Force:
- Generate all substrings → O(n²)
- Check if each is a palindrome → O(n)
- Overall time: O(n³)

Expand Around Center:
- For each center, expand outward while characters match
- There are 2n-1 centers (odd and even)
- Each expansion takes up to O(n)
- Overall time: O(n²)

🔷 TIME & SPACE COMPLEXITY:

Time:  O(n²)
Space: O(1) — constant extra space (no DP table or extra structures)

🔷 CONCLUSION:
In brute force approach we generate all substrings, but in optimized approach we generate all possible palindromes
so we reduce the time complexity from O(n³) to O(n²)
*/

func main() {
	fmt.Println(lps("babad")) // Output: "bab" or "aba"
	fmt.Println(lps("cbbd"))  // Output: "bb"
	fmt.Println(lps("a"))     // Output: "a"
	fmt.Println(lps("ac"))    // Output: "a" or "c"
}

// lps finds the longest palindromic substring using center expansion
func lps(str string) string {
	longest := ""

	for i := 0; i < len(str); i++ {
		// Expand around odd-length center (single character center)
		p1 := expandFromCenter(str, i, i)

		// Expand around even-length center (between two characters)
		p2 := expandFromCenter(str, i, i+1)

		// Update longest if necessary
		if len(p1) > len(longest) {
			longest = p1
		}
		if len(p2) > len(longest) {
			longest = p2
		}
	}

	return longest
}

// expandFromCenter returns the longest palindrome found from a given center
func expandFromCenter(str string, left, right int) string {
	for left >= 0 && right < len(str) && str[left] == str[right] {
		left--
		right++
	}
	// After the loop ends, left and right have gone one step too far so the last valid palindrome was from
	// left+1 to right-1 and the second index in slice is not included so we need [left+1:right-1+1]
	return str[left+1 : right]
}
