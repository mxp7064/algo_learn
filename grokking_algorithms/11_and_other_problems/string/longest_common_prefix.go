package main

import "fmt"

/*
===========================================================
🧠 Longest Common Prefix — Horizontal Scanning Approach
===========================================================

Approach:
- Start with the first word as the current prefix.
- For each following word, trim the prefix from the end
  until it matches the beginning of the current word.
- If at any point the prefix becomes empty, return "" early.

Why this works:
- We're narrowing down the prefix by comparing it across all strings.
- Efficient for small-to-medium input sizes.

Time Complexity: O(S), where S is the total number of characters in all strings
Space Complexity: O(1) extra space
*/

func main() {
	fmt.Println(lcp([]string{"Marko", "Mario"}))                             // "Mar"
	fmt.Println(lcp([]string{"Baba", "Mama"}))                               // ""
	fmt.Println(lcp([]string{"flower", "flow", "flight"}))                   // "fl"
	fmt.Println(lcp([]string{"dog", "racecar", "car"}))                      // ""
	fmt.Println(lcp([]string{"interspecies", "interstellar", "interstate"})) // "inters"
}

// lcp returns the longest common prefix among all strings in the array.
func lcp(arr []string) string {
	if len(arr) == 0 {
		return ""
	}

	// Start with the first string as the initial prefix
	prefix := arr[0]

	// Compare the prefix with every other string
	for i := 1; i < len(arr); i++ {
		current := arr[i]

		// Trim prefix until it matches the start of current string
		for len(prefix) > 0 && !startsWith(current, prefix) {
			// Remove last character from prefix
			prefix = prefix[:len(prefix)-1]
		}

		// Early exit if there's no common prefix
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

// startsWith returns true if string s starts with the given prefix
func startsWith(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	return s[:len(prefix)] == prefix
}
