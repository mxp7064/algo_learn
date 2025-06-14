package main

import (
	"fmt"
	"slices"
)

// groupAnagrams groups strings that are anagrams of each other.
// Anagrams share the same sorted character signature (for example "mug" and "gum" are anagrams)
// Input is a list of words, output is an array of string arrays where each inner array is a group of words
// which are anagrams of each other
// We can easily detect that two words are anagrams of each other by just comparing their sorted characters
// We can then use a map to track and group/store anagrams

// Time Complexity:
// Let n = number of words, and k = average length of each word
//
// For each word (n):
// - Sorting characters: O(k log k)
// Therefore, total time = O(n * k log k)
//
// Space Complexity:
// - Map stores up to n groups with up to k characters per key/value
// - Total space = O(n * k)
func groupAnagrams(words []string) [][]string {
	// key: sorted word
	// value: list of words that are anagrams to that word
	group := make(map[string][]string)

	for _, word := range words {
		// Convert word to rune slice, then sort it
		wordChars := []rune(word)
		slices.Sort(wordChars)
		sorted := string(wordChars)

		// Append original word to its group
		group[sorted] = append(group[sorted], word)
	}

	// Collect all groups into the final result
	var result [][]string
	for _, g := range group {
		result = append(result, g)
	}

	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"mug", "gum", "gm", "tea", "ate"})) // [[mug gum] [gm] [tea ate]]
}
