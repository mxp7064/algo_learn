package trie

import (
	"fmt"
	"testing"
)

func Test_Search(t *testing.T) {
	tr := NewTrieFromArray("cat", "car", "cart")

	fmt.Println("tr.Search(\"car\"): ", tr.Search("car"))  // true
	fmt.Println("tr.Search(\"cart\": ", tr.Search("cart")) // true
	fmt.Println("tr.Search(\"ca\"): ", tr.Search("ca"))    // false
	fmt.Println("tr.Search(\"x\"): ", tr.Search("x"))      // false
}

func Test_StartsWith(t *testing.T) {
	tr := NewTrieFromArray("cat", "car", "cart")

	fmt.Println("tr.StartsWith(\"x\"): ", tr.StartsWith("x"))     // false
	fmt.Println("tr.StartsWith(\"ca\"): ", tr.StartsWith("ca"))   // true
	fmt.Println("tr.StartsWith(\"dog\"): ", tr.StartsWith("dog")) // false
}

func Test_Delete(t *testing.T) {
	tr := NewTrieFromArray("app", "apple", "ape", "bat", "batch")

	fmt.Println("Delete(tr, \"apple\"): ", Delete(tr, "apple")) // true
	fmt.Println("Delete(tr, \"app\"): ", Delete(tr, "app"))     // true
	fmt.Println("Delete(tr, \"zoo\"): ", Delete(tr, "zoo"))     // false
	fmt.Println("Delete(tr, \"ape\"): ", Delete(tr, "ape"))     // true
	fmt.Println("Delete(tr, \"bat\"): ", Delete(tr, "bat"))     // true
	fmt.Println("Delete(tr, \"batch\"): ", Delete(tr, "batch")) // true
	fmt.Println("Delete(tr, \"batch\"): ", Delete(tr, "batch")) // false, already deleted
}

func Test_AllWords(t *testing.T) {
	tr := NewTrieFromArray("dog", "dot", "dove")

	fmt.Println(GetAllWords(tr)) // [dog dot dove]
}

func Test_Update(t *testing.T) {
	tr := NewTrieFromArray("code", "coder")

	fmt.Println(Update(tr, "code", "cope")) // true
	fmt.Println(GetAllWords(tr))            // [cope coder]
}
