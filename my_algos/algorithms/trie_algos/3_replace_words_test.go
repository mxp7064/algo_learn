package trie_algos

import (
	"fmt"
	"panca.com/algo/trie"
	"strings"
	"testing"
)

/*
LeetCode 648 – Replace Words

💡 Goal:
Given a dictionary of root words (like ["cat", "bat", "rat"]) and a sentence ("the cattle was rattled by the battery"),
replace all words in the sentence that start with a root with the shortest such root.

✅ For example:
Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
Output: "the cat was rat by the bat"

🔍 Why Trie?
- You want to efficiently find the shortest prefix of each word that matches a root.
- Tries allow walking each word character-by-character and stopping as soon as we find a match.

📈 Time Complexity:
- Building the trie = O(N * K), where N = number of roots, K = avg root length
- Replacing words = O(M * L), where M = number of words in the sentence, L = avg word length
- Total = O(N*K + M*L)
*/

// ReplaceWords replaces words in the sentence using the shortest matching root from the dictionary
func ReplaceWords(dictionary []string, sentence string) string {
	tr := trie.NewTrieFromArray(dictionary...)

	words := strings.Split(sentence, " ")
	for i, word := range words { // for each word in the sentence
		prefix := ""
		node := tr.Root           // start from root
		for _, ch := range word { // for each character in the word
			if node.Children[ch] == nil {
				break // no further prefix match
			}
			prefix += string(ch)     // update prefix
			node = node.Children[ch] // go to the next character
			if node.IsEnd {
				words[i] = prefix // replace word with shortest root
				break             // early exist
			}
		}
	}

	return strings.Join(words, " ")
}

func Test_ReplaceWords(t *testing.T) {
	dictionary := []string{"cat", "bat", "ba", "battle", "rat"}
	sentence := "the cattle was rattled by the battery" // battery should be replaced with "ba", not "bat" or "battle"
	result := ReplaceWords(dictionary, sentence)
	fmt.Println(result) // Expected: "the cat was rat by the ba"
}
