package main

import (
	"fmt"
)

/*
========================================================
🧠 Longest Substring Without Repeating Characters
========================================================

Given a string s, find the length of the longest substring without repeating characters.

Approach:
- Use a sliding window (left and right pointers)
- Use a map to track the last seen index of each character
- If a character repeats inside the current window, move the left pointer to exclude the duplicate
- Always update the max window length

Time: O(n) — each character is visited at most twice
Space: O(k) — where k is number of unique characters in the input (typically O(1) if limited to ASCII)
*/

func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[rune]int) // rune → last seen index
	maxLen := 0
	left := 0

	for right, c := range s {
		if prevIndex, found := lastSeen[c]; found && prevIndex >= left {
			// character was seen in current window → move left pointer
			// this shrinks the window to exclude the duplicate
			left = prevIndex + 1
		}

		// Update last seen index of character
		lastSeen[c] = right

		// Update max length
		windowLen := right - left + 1
		if windowLen > maxLen {
			maxLen = windowLen
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3 ("abc")
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1 ("b")
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3 ("wke")
	fmt.Println(lengthOfLongestSubstring(""))         // Output: 0
	fmt.Println(lengthOfLongestSubstring("dvdf"))     // Output: 3 ("vdf")
}
