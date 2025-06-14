package trie_algos

import (
	"fmt"
	"panca.com/algo/trie"
	"sort"
	"testing"
)

/*
Autocomplete – Prefix-Based Word Retrieval Using Trie

💡 Goal:
Given a prefix (e.g. "car") and a trie filled with words, return all words that start with that prefix.

🔍 Why Trie?
- In a trie, all words that share the same prefix follow the same branch.
- This structure makes prefix-based retrieval very efficient, both in time and memory.

📈 Time Complexity: O(k + output)
- k = length of the prefix
- output = total number of characters in all matching words (not number of words!)
- Step-by-step:
  1. Walk from root to prefix node → O(k)
  2. DFS from that node to collect words → proportional to output

💡 Compared to brute-force:
- Brute-force would check every word: O(n * k) where n is number of words
- Trie avoids unnecessary comparisons — it skips entire branches that don't match the prefix

🔧 How it works:
1. Traverse the trie down to the last character in the prefix
2. From there, recursively collect all valid words using DFS

🧪 Visual Example for prefix "car":
(root)
 └── c
      └── a
           └── r (✓)
                ├── t (✓)
                └── b
                     └── o
                          └── n (✓)

✅ Result: ["car", "cart", "carbon"]
*/

// Autocomplete returns all words in the trie that begin with the given prefix
// Time complexity: O(k + output), where k = prefix length
func Autocomplete(t *trie.Trie, prefix string) []string {
	node := t.Root
	for _, ch := range prefix {
		if node.Children[ch] == nil {
			return nil // Prefix not found
		}
		node = node.Children[ch]
	}

	// Node is at last prefix character, now perform DFS to collect all words
	var results []string
	collectWords(node, prefix, &results)
	return results
}

// collectWords performs DFS from the current node and collects valid words
func collectWords(node *trie.Node, path string, results *[]string) {
	if node == nil {
		return
	}

	// If this node marks the end of a word, add the current path to results
	if node.IsEnd {
		*results = append(*results, path)
	}

	// Recurse into all children, appending their character to the path
	for ch, child := range node.Children {
		collectWords(child, path+string(ch), results)
	}
}

// PrefixCount returns how many words in the trie start with the given prefix.
// This is a prefix count in trie problem - variation of Autocomplete.
func PrefixCount(t *trie.Trie, prefix string) int {
	node := t.Root
	for _, ch := range prefix {
		if node.Children[ch] == nil {
			return 0 // Prefix not found -> 0 words start with that prefix
		}
		node = node.Children[ch]
	}

	// Node is at last prefix character, now perform DFS to collect all words
	count := 0
	countWords(node, &count)
	return count
}

// countWords performs DFS from the current node and counts valid words
func countWords(node *trie.Node, count *int) {
	if node == nil {
		return
	}

	// If this node marks the end of a word, increment the counter
	if node.IsEnd {
		*count++
	}

	// Recurse into all children
	for _, child := range node.Children {
		countWords(child, count)
	}
}

// PrefixCountPrecomputedCount returns how many words in the trie start with the given prefix using the Count field that
// we increment each time we traverse a node in trie.Insert method
// So we can just walk to the node for last prefix character and return node.count
// This is an optimized version of PrefixCount, time complexity is O(k), k = length of prefix
func PrefixCountPrecomputedCount(t *trie.Trie, prefix string) int {
	node := t.Root
	for _, ch := range prefix {
		if node.Children[ch] == nil {
			return 0 // Prefix not found -> 0 words start with that prefix
		}
		node = node.Children[ch]
	}

	return node.Count
}

// AutocompleteWithSuggestions returns suggestions after each character typed (we simulate this by calling Autocomplete for each "typed character")
// Returns top 3 lexicographically sorted matches at each step
// This is LeetCode 1268 (Search Suggestions System)
// If searchWord is "cart" we will have output: [[car carbon cart] [car carbon cart] [car carbon cart] [cart]]
// Explanation:
// Step 1 (prefix "c"): [car carbon cart]
// Step 2 (prefix "ca"): [car carbon cart]
// Step 3 (prefix "car"): [car carbon cart]
// Step 4 (prefix "cart"): [cart]
func AutocompleteWithSuggestions(t *trie.Trie, searchWord string) [][]string {
	var suggestions [][]string
	prefix := ""
	for _, ch := range searchWord {
		prefix += string(ch) // first "c", then "ca", then "car" and then "cart"
		matches := Autocomplete(t, prefix)

		// we want first 3 results sorted
		sort.Strings(matches)
		if len(matches) > 3 {
			matches = matches[:3] // truncate to length 3
		}

		suggestions = append(suggestions, matches)
	}
	return suggestions
}

func Test_Autocomplete(t *testing.T) {
	tr := trie.NewTrie()
	tr.Insert("cat")
	tr.Insert("car")
	tr.Insert("cart")
	tr.Insert("carbon")
	tr.Insert("dog")

	fmt.Println(Autocomplete(tr, "car")) // [car carbon cart]
	fmt.Println(Autocomplete(tr, "ca"))  // [cat car carbon cart]
	fmt.Println(Autocomplete(tr, "do"))  // [dog]
	fmt.Println(Autocomplete(tr, "z"))   // []

	fmt.Println("PrefixCount(tr, \"ca\"):", PrefixCount(tr, "ca"))     // 4
	fmt.Println("PrefixCount(tr, \"z\"):", PrefixCount(tr, "z"))       // 0
	fmt.Println("PrefixCount(tr, \"car\"):", PrefixCount(tr, "car"))   // 3
	fmt.Println("PrefixCount(tr, \"cart\"):", PrefixCount(tr, "cart")) // 1

	fmt.Println("PrefixCountPrecomputedCount(tr, \"ca\"):", PrefixCountPrecomputedCount(tr, "ca"))     // 4
	fmt.Println("PrefixCountPrecomputedCount(tr, \"z\"):", PrefixCountPrecomputedCount(tr, "z"))       // 0
	fmt.Println("PrefixCountPrecomputedCount(tr, \"car\"):", PrefixCountPrecomputedCount(tr, "car"))   // 3
	fmt.Println("PrefixCountPrecomputedCount(tr, \"cart\"):", PrefixCountPrecomputedCount(tr, "cart")) // 1

	fmt.Println("\nAutocompleteWithSuggestions:")
	searchWord := "cart"
	res := AutocompleteWithSuggestions(tr, searchWord)
	for i, step := range res {
		fmt.Printf("Step %d (prefix \"%s\"): %v\n", i+1, searchWord[:i+1], step)
	}
}
