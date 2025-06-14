package main

import (
	"fmt"
	"slices"
)

/*
===================================================
✅ VALID ANAGRAM — Two Approaches
===================================================

🧠 PROBLEM:
Given two strings s and t, return true if t is an anagram of s.
An anagram is formed by rearranging the letters of another (ex. "listen" and "silent" are anagrams)

===================================================
🔹 APPROACH 1: SORTING
===================================================
- Convert both strings to rune slices.
- Sort both slices.
- Compare sorted strings.

Time Complexity: O(n log n) – for sorting
Space Complexity: O(n) – for rune slices

===================================================
🔹 APPROACH 2: CHARACTER FREQUENCY MAP
===================================================
- Use a map to count character frequencies in s.
- Subtract character counts using t.
- If the map is empty at the end, they are anagrams.

Time Complexity: O(n)
Space Complexity: O(k), where k is number of unique characters

===================================================
*/

func main() {
	// Test case
	fmt.Println(isAnagram("tugaax", "gutaa"))         // false
	fmt.Println(isAnagram("listen", "silent"))        // true
	fmt.Println(isAnagramSorting("listen", "silent")) // true
	fmt.Println(isAnagramSorting("hello", "helloo"))  // false
}

// isAnagramSorting checks if t is an anagram of s using sorting.
func isAnagramSorting(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Convert strings to rune slices to sort them
	r1 := []rune(s)
	r2 := []rune(t)

	slices.Sort(r1)
	slices.Sort(r2)

	// Compare the sorted versions
	return string(r1) == string(r2)
}

// isAnagram checks if t is an anagram of s using a character frequency map.
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freqMap := make(map[rune]int)

	// Count characters in s
	for _, c := range s {
		freqMap[c]++
	}

	// Subtract counts using t
	for _, c := range t {
		freqMap[c]--
		if freqMap[c] == 0 {
			delete(freqMap, c)
		} else if freqMap[c] < 0 {
			return false // early exit if t has extra chars
		}
	}

	// If map is empty, it's a valid anagram
	return len(freqMap) == 0
}
