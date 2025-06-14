package main

import (
	"fmt"
	"math"
	"strings"
)

// This is rearely directly asked in DSA interview, but it's good to know in general.

// Naive Bayes Spam Filter (simplified bag-of-words implementation)

// Problem: Given labeled training data (spam / not_spam), classify a new message.
// Uses word frequency and Bayes theorem to estimate P(spam|message).

// Time complexity: O(n * m), n = number of samples, m = number of words
// Space complexity: O(v), v = vocabulary size

type Sample struct {
	Text  string
	Label string
}

func main() {
	// Training data
	data := []Sample{
		{"win money now", "spam"},
		{"lowest price for best deal", "spam"},
		{"earn dollars fast", "spam"},
		{"hi, are we still on for dinner?", "not_spam"},
		{"meeting scheduled at 10am", "not_spam"},
		{"project deadline tomorrow", "not_spam"},
	}

	classifier := trainNaiveBayes(data)

	tests := []string{
		"win a prize now",
		"meeting about project",
		"low price dollars",
		"are we meeting now?",
	}

	for _, msg := range tests {
		pred := classifier.predict(msg)
		fmt.Printf("Message: %-30q → Prediction: %s\n", msg, pred)
	}
}

// NaiveBayesClassifier stores frequencies and priors
type NaiveBayesClassifier struct {
	wordCount     map[string]map[string]int
	totalWords    map[string]int
	totalDocs     map[string]int
	vocabulary    map[string]bool
	totalDocCount int
}

func trainNaiveBayes(data []Sample) *NaiveBayesClassifier {
	nb := &NaiveBayesClassifier{
		wordCount:     make(map[string]map[string]int),
		totalWords:    make(map[string]int),
		totalDocs:     make(map[string]int),
		vocabulary:    make(map[string]bool),
		totalDocCount: len(data),
	}

	for _, sample := range data {
		label := sample.Label
		nb.totalDocs[label]++

		if nb.wordCount[label] == nil {
			nb.wordCount[label] = make(map[string]int)
		}

		words := tokenize(sample.Text)
		for _, word := range words {
			nb.wordCount[label][word]++
			nb.totalWords[label]++
			nb.vocabulary[word] = true
		}
	}

	return nb
}

func (nb *NaiveBayesClassifier) predict(text string) string {
	words := tokenize(text)
	bestLabel := ""
	maxLogProb := math.Inf(-1)

	for label := range nb.totalDocs {
		logProb := math.Log(float64(nb.totalDocs[label]) / float64(nb.totalDocCount))

		for _, word := range words {
			count := nb.wordCount[label][word]
			smoothed := float64(count+1) / float64(nb.totalWords[label]+len(nb.vocabulary))
			logProb += math.Log(smoothed)
		}

		if logProb > maxLogProb {
			maxLogProb = logProb
			bestLabel = label
		}
	}

	return bestLabel
}

func tokenize(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
