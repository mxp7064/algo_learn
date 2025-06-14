package main

import "fmt"

/*
=============================================================
🧠 Longest Palindromic Substring — Brute Force O(n³) Approach
=============================================================

Problem statement:
Given a string s, return the longest substring of s that is a palindrome.
A palindrome is a word that reads the same forwards and backwards (ex. 'ana', 'abba', etc.)

Approach:
- Generate all possible substrings of the input string (O(n²))
- For each substring, check if it's a palindrome (O(n))
- Track and return the longest palindromic substring found

Time Complexity:
- O(n²) substrings × O(n) palindrome check = O(n³)

Note:
- This is the brute-force solution just for demonstration and learning.
- We'll later optimize it using the expand-around-center approach (2_expand_around_center.go)
*/

func main() {
	fmt.Println(lps("babad")) // Output: "bab" or "aba"
	fmt.Println(lps("cbbd"))  // Output: "bb"
	fmt.Println(lps("a"))     // Output: "a"
	fmt.Println(lps("ac"))    // Output: "a" or "c"
}

func checkPalindrome(str string) bool {
	left := 0
	right := len(str) - 1
	for left <= right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// lps returns the longest palindromic substring using brute-force
func lps(str string) string {
	longest := ""

	// if str = panca, generated substrings would be:
	// p, pa, pan, panc, panca
	// a, an, anc, anca
	// n, nc, nca
	// c, ca
	// a

	// generate all substrings starting from index 0, 1,...n-1
	for i := 0; i < len(str); i++ {
		// for each index, generate 1 len, 2 len,...n len substrings
		for j := i; j < len(str); j++ {
			substr := str[i : j+1]

			// Check if substring is a palindrome and longer than current longest
			if checkPalindrome(substr) && len(substr) > len(longest) {
				longest = substr
			}
		}
	}

	return longest
}
