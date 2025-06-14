/*

This is rarely asked directly as a problem in DSA interview, but it's good to know, especially for system design discussion (part of full search topic).

Inverted Index in Go

Problem Statement:
------------------
Given a set of web pages (documents), build a search index that maps each word
to the list of documents where that word appears. This is called an inverted index.

Then, when a user types a search query (e.g., two words), return all documents
that contain ALL the queried words (i.e., perform set intersection).

Time Complexity:
----------------
- Building the index: O(n * m), where n is the number of docs, and m is avg number of words per doc
- Querying: O(k * d), where k is the number of search words and d is average number of docs per word

Why we use map[string]map[string]bool instead of map[string][]string:
----------------------------------------------------------------------
- To avoid duplicate document IDs in the list (map keys are unique).
- To allow fast O(1) membership checks when doing intersection logic.
- Set-like operations (intersection, union, etc.) are easier and faster with maps.

Alternative (map[string][]string) is less efficient because you'd have to manually
check for duplicates and do slow linear search for intersections.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Sample documents: documentID → content
	documents := map[string]string{
		"A": "hi there how are you",
		"B": "hi there you are again",
		"C": "there is something else",
	}

	// Build the inverted index
	index := buildInvertedIndex(documents)

	// Example query: find docs that contain both "hi" and "there"
	query := []string{"hi", "there"}
	result := search(index, query)

	fmt.Printf("Query: %v\n", query)
	fmt.Println("Matching documents:", result)
}

// buildInvertedIndex creates a map of word → set of document IDs
func buildInvertedIndex(docs map[string]string) map[string]map[string]bool {
	index := make(map[string]map[string]bool)

	for docID, text := range docs {
		words := strings.Fields(text) // split by spaces
		for _, word := range words {
			// Initialize inner map (set of docIDs) if needed
			if index[word] == nil {
				index[word] = make(map[string]bool)
			}
			// Add docID to set of documents containing this word
			index[word][docID] = true
		}
	}

	return index
}

// search finds the intersection of document sets for all query words
func search(index map[string]map[string]bool, queryWords []string) []string {
	if len(queryWords) == 0 {
		return nil
	}

	// Start with docs that contain the first word
	result := make(map[string]bool) // this is a set which contains documents (document IDs) which contain all queryWords
	firstWord := queryWords[0]
	for doc := range index[firstWord] {
		result[doc] = true
	}

	// Intersect with the sets for remaining words
	remainingWords := queryWords[1:]
	for _, word := range remainingWords {
		if docsForWord, ok := index[word]; ok { // if word exists in index
			for doc := range result { // loop set which contains documents that contain the first word
				if !docsForWord[doc] {
					// doc exists in set of documents which contain the first word, but it doesn't exist in the set of
					// documents which contain this word - this doc is not part of intersection, so we remove it
					delete(result, doc)
				}
			}
		} else {
			// If the word doesn't exist in index (there are no documents containing this word) -> return no results
			return nil
		}
	}

	// Convert final set to slice
	var finalResult []string
	for doc := range result {
		finalResult = append(finalResult, doc)
	}
	return finalResult
}
