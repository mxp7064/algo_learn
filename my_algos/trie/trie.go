package trie

import "panca.com/algo/stack"

/*
Trie – Prefix Tree

A Trie (pronounced "try") is a tree-like data structure used to store a dynamic set of strings.
It excels at tasks involving prefix matching, autocomplete, and dictionary-like lookups.

Why not just use an array or hash map?
- Arrays/lists store entire words separately. So if you have ["car", "cart", "carbon"], they each store full copies of "car", duplicating the "car" part.
- Hash maps support full-word lookup in O(1) time, but prefix-based operations like autocomplete or prefix existence check become O(n * k) where n is the number of words and k is the average word length.

Trie eliminates duplication of prefixes and enables:
- Fast word insert/search: O(k)
- Fast prefix search: O(k)
- Fast autocomplete: O(k + output size)
k = length of the word/prefix

Visual Example for: ["cat", "car", "cart"]
(root)
 └── c
      └── a
           ├── t (✓)
           └── r (✓)
                └── t (✓)
Each path from root to ✓ forms a complete word.

Note: methods in this package can appear in DSA interview so make sure to understand them carefully:
- LeetCode 208 – Implement Trie
	- Implement Insert, Search, and StartsWith
*/

// Node represents a node in trie - it doesn't store its own character value - instead it stores a children map and isEnd flag
// How can we know then which character maps to which node? Children map - so information about the character/node mapping is always stored in parent (not in the node itself)
type Node struct {
	Children map[rune]*Node // Child nodes for each character
	IsEnd    bool           // Marks if this node completes a valid word
	Count    int
}

// Trie is the main structure wrapping the root node
// All operations start from the root
type Trie struct {
	Root *Node
}

// NewNode creates an empty node with initialized children map
func NewNode() *Node {
	return &Node{Children: make(map[rune]*Node)}
}

// NewTrie creates a new Trie with an empty root node
func NewTrie() *Trie {
	return &Trie{
		Root: NewNode(),
	}
}

func NewTrieFromArray(words ...string) *Trie {
	tr := NewTrie()
	for _, root := range words {
		tr.Insert(root)
	}
	return tr
}

// Insert adds a word to the trie character by character
// Shared prefixes are reused, only new branches are added
func (t *Trie) Insert(word string) {
	node := t.Root
	for _, ch := range word {
		if node.Children[ch] == nil { // If child doesn't exist, create it
			node.Children[ch] = NewNode()
		}
		node = node.Children[ch]

		// As we traverse each node during Insert, increment the Count which represents the number of descendants
		// of this node - i.e. number of words that start with the prefix [root...node]
		// It's like saying: "One more word passes through here". Count is then the number of words that begin with the prefix up to that node
		node.Count++
	}
	node.IsEnd = true // Mark the last node as end of word
}

// Search returns true if the full word exists in the trie
func (t *Trie) Search(word string) bool {
	node := t.Root
	for _, ch := range word {
		if node.Children[ch] == nil {
			return false // word doesn't exist (there is no such character path)
		}
		node = node.Children[ch]
	}
	return node.IsEnd
}

// StartsWith returns true if any word in the trie starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t.Root
	for _, ch := range prefix {
		if node.Children[ch] == nil {
			return false // prefix doesn't exist (there is no such character path)
		}
		node = node.Children[ch]
	}
	return true
}

/*
Update – Replace one word with another
Returns true if the update was successful (oldWord existed and was replaced), otherwise it returns false.

This is a compound operation:
1. Attempt to delete `oldWord`
2. If successful, insert `newWord`

Time Complexity: O(k), where k = max(len(oldWord), len(newWord))
*/
func Update(t *Trie, oldWord, newWord string) bool {
	if !Delete(t, oldWord) {
		return false
	}
	t.Insert(newWord)
	return true
}

/*
GetAllWords – Return all words currently stored in the trie

Uses DFS to traverse every path from root to word endings
Time Complexity: O(n * k), where:
- n = number of words in the trie
- k = average length of the words
*/
func GetAllWords(t *Trie) []string {
	var result []string
	collect(t.Root, "", &result)
	return result
}

func collect(node *Node, path string, result *[]string) {
	// if word is found, append it to result
	if node.IsEnd {
		*result = append(*result, path)
	}

	// recursively traverse children
	for ch, child := range node.Children {
		collect(child, path+string(ch), result)
	}
}

/*
Delete – Safely remove a word from Trie. Returns true if node exists and was deleted, otherwise, returns false.

💡 Goal:
Implement deletion of a word while:
- Safely preserving prefixes used by other words
- Decrementing count fields
- Cleaning up unnecessary nodes if no longer needed

🔍 Key Requirements:
- If word doesn’t exist → return false
- If word exists:
  - Unset `isEnd`
  - Decrement `.Count` along the path
  - Remove child nodes if they are no longer needed (count == 0 and not end of any other word)

Approach:
- Step 1: collect all nodes along the word path (from root till the last character in the word)
  - follow the path for the word (for example "cat"): "c" → "a" → "t" (like in Search)
  - while doing this, we push the nodes (and their runes) onto a stack so we remember the path
  - stack entry stores each character and its parent node reference as we traverse through the word characters
  - we need parent node reference because later we will need to delete the child node from its parent’s Children map
  - example -> for deleting "cat" our path will be: root → 'c' → 'a' → 't', so stack will contain (top to bottom)
  - ('t', node for 'a')
  - ('a', node for 'c')
  - ('c', root)

- Step 2: delete the word by unmarking isEnd on the last node (e.g., on "t" in "cat")

- Step 3: clean up in reverse (by popping from stack)
  - decrement count for each node
  - remove any child node if it's no longer needed (it has no children and it’s not the end of any other word)

- note: we use a stack because we want to clean up the trie in reverse (direction last char -> root)

📈 Time Complexity: O(k), where k = length of word

📊 Example:
Inserted words: ["app", "apple", "ape"]

Before Delete("apple"):
(root)

	└── a
	     └── p
	          ├── p (isEnd)
	          │    └── l
	          │         └── e (isEnd)  ← "apple"
	          └── e (isEnd)            ← "ape"

After Delete("apple"):
(root)

	└── a
	     └── p
	          ├── p (isEnd)
	          └── e (isEnd)            ← "ape"

✅ Cleaned up unused nodes ("l" and its child "e")
*/
type stackEntry struct {
	parent *Node // parent from char
	char   rune
}

func Delete(t *Trie, word string) bool {
	st := stack.NewStack[stackEntry]()
	node := t.Root

	// Step 1: Traverse the word, collect path in stack
	for _, ch := range word {
		if node.Children[ch] == nil {
			return false // word not found
		}
		st.Push(stackEntry{node, ch})
		node = node.Children[ch]
	}

	// node is now at the last character in word
	if !node.IsEnd {
		return false // word exists as prefix only
	}

	// Step 2: delete the word (unmark isEnd)
	node.IsEnd = false

	// Step 3: backtrack and clean up unused nodes
	cleanup(st)

	return true // return true if word exists
}

func cleanup(st stack.Stack[stackEntry]) {
	for !st.IsEmpty() {
		entry := st.Pop()
		node := entry.parent.Children[entry.char]

		node.Count-- // decrement count as this word is removed - one less word "passes through this node now"

		// If node has no more children and is not end of another word
		if node.Count == 0 && !node.IsEnd {
			delete(entry.parent.Children, entry.char) // prune unused branch
		} else {
			break // Stop cleanup once a shared path is encountered - some other word(s) depends on this character
		}
	}
}
