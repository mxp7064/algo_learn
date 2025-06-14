package trie_algos

import (
	"fmt"
	"panca.com/algo/trie"
	"testing"
)

/*

SearchWithWildcard – Flexible Pattern Matching Using Trie

💡 Goal:
Allow search for a word in a trie using '.' as a wildcard character, where each '.' can match any single character.

🔍 Why Trie?
- With standard characters, trie lookup is O(k) – same as exact word search
- With '.', we branch into all children at that node
- This makes it a flexible tool for regex-like matching or masked word queries

📈 Time Complexity:
- Worst-case: O(26^w), where w is the number of wildcards ('.')
- For regular characters, we follow one path: O(k)
- For each '.', we may branch into all children → exponential if wildcards are stacked

🔧 How It Works:
1. Recursively walk through the trie
2. At each character:
   - If it's a normal letter → move to the corresponding child (next character)
   - If it's '.' → try all children recursively
3. Base case: if `index == len(word)`, check if current node ends a word

🧪 Visual Trace example:
Assume your trie contains ["cat", "cut", "cot", "dog"]:

(root)
 ├── c
 │    ├── a
 │    │    └── t (✓)
 │    ├── o
 │    │    └── t (✓)
 │    └── u
 │         └── t (✓)
 └── d
      └── o
           └── g (✓)

Let's trace the call SearchWithWildcard("c.t"):

searchHelper(root, ['c','.', 't'], 0)
│
├── node has child 'c' → go deeper
    searchHelper(node='c', ['c','.', 't'], 1)
    │
    ├── word[1] is '.' → wildcard
    │   Try all children of 'c': ['a', 'o', 'u']
    │
    ├── Try 'a':
    │   searchHelper(node='a', ['c','.', 't'], 2)
    │   └── word[2] == 't' → go to child 't'
    │       searchHelper(node='t', ['c','.', 't'], 3)
    │       └── index == 3 == len(word) → return node.isEnd = true ✅
    │
    ├── Try 'o':
    │   same path: 'c' → 'o' → 't' → isEnd = true ✅
    │
    ├── Try 'u':
    │   same path: 'c' → 'u' → 't' → isEnd = true ✅
    │
    └── So wildcard match succeeds → return true to root ✅

✅ Any success returns true (first one)

If it were "c.x" (no such words):

searchHelper('c', ['c','.', 'x'], 1)
 ├── Try 'a' → no 'x' child → return false
 ├── Try 'o' → no 'x' child → return false
 ├── Try 'u' → no 'x' child → return false
 → return false

📊 Example usage:
- SearchWithWildcard("c.t") → true (matches "cat", "cot", "cut")
- SearchWithWildcard("..t") → matches all 3-letter words ending in 't'

Note: This is almost same as LeetCode 211 (Design Add and Search Words Data Structure)
- AddWord method is in the trie package (Insert)
- Search is SearchWithWildcard
- We are just missing the WordDictionary struct wrapper
*/

func SearchWithWildcard(t *trie.Trie, word string) bool {
	return searchHelper(t.Root, []rune(word), 0)
}

// searchHelper performs DFS with branching logic for '.' wildcards
func searchHelper(node *trie.Node, word []rune, index int) bool {
	// Base case: reached end of pattern
	if index == len(word) {
		return node.IsEnd
	}

	ch := word[index]

	// Wildcard case: try every possible child
	if ch == '.' {
		for _, child := range node.Children {
			if searchHelper(child, word, index+1) {
				return true // early exit upon first successful match
			}
		}
		return false // none matched
	}

	// Normal character case
	if node.Children[ch] == nil {
		return false // character is not in path
	}
	// character matches, recurse on next character
	return searchHelper(node.Children[ch], word, index+1)

}

func Test_SearchWithWildcard(t *testing.T) {
	tr := trie.NewTrie()
	tr.Insert("cat")
	tr.Insert("car")
	tr.Insert("cart")
	tr.Insert("carbon")
	tr.Insert("dog")

	fmt.Println(SearchWithWildcard(tr, "c.t"))    // true
	fmt.Println(SearchWithWildcard(tr, "c.r"))    // true
	fmt.Println(SearchWithWildcard(tr, "c.x"))    // false
	fmt.Println(SearchWithWildcard(tr, "a."))     // false
	fmt.Println(SearchWithWildcard(tr, "c."))     // false
	fmt.Println(SearchWithWildcard(tr, "c.r..n")) // true
}
