package main

import (
	"fmt"
	"math"
	"sort"
)

// ==============================================
// KNN-based OCR Simulation in Go
//
// 🧠 Problem Statement:
// Simulate a basic version of Optical Character Recognition (OCR)
// using the k-nearest neighbors (KNN) algorithm.
// Given a new character represented by numerical features (not image pixels),
// determine which digit (0-9) it most likely represents based on known samples.
//
// 🎯 Goal:
// Use Euclidean distance to compare the new character to existing
// labeled examples and classify it based on the majority vote of k nearest neighbors.
// So we are doing classification by similarity
//
// 🔍 Real-world Context:
// In real OCR, images are scanned and processed into feature vectors
// (e.g. stroke counts, curves, aspect ratio).
// This example skips the image part and uses hand-crafted features for simplicity.
//
// 💡 Example:
// Digit 8 may be represented as [2 loops, 0 lines] = [2, 0]
// Digit 1 may be [0 loops, 2 lines] = [0, 2]
//
// ==============================================

// Digit represents a labeled example with extracted features
type Digit struct {
	Label    int       // e.g., 0–9
	Features []float64 // simulated features like [loops, lines]
}

// EuclideanDistance computes the distance between two vectors
func EuclideanDistance(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

// predictDigit performs the KNN classification
func predictDigit(trainingSet []Digit, inputFeatures []float64, k int) int {
	// Step 1: Calculate distances between input and all vectors in training set
	type neighbor struct {
		label    int
		distance float64
	}
	var distances []neighbor
	for _, digit := range trainingSet {
		dist := EuclideanDistance(digit.Features, inputFeatures)
		distances = append(distances, neighbor{label: digit.Label, distance: dist})
	}

	// Step 2: Sort neighbors by distance (ascending)
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	// Step 3: Majority voting on k nearest neighbors
	voteCount := make(map[int]int) // key is digit label, value is number of occurrences
	for i := 0; i < k; i++ {       // we take first k elements and count number of occurrences
		voteCount[distances[i].label]++
	}

	// Step 4: Return the digit label with the most votes
	bestLabel := -1
	maxVotes := 0
	for label, votes := range voteCount {
		// for simplicity, we don't have a tie breaking strategy
		// results can be random in case of a tie since we are using Go map
		if votes > maxVotes {
			bestLabel = label
			maxVotes = votes
		}
	}
	return bestLabel
}

func main() {
	// Training set: each entry is a digit (label) with its simplified feature vector (loop counts and line counts)
	// values are arbitrary
	trainingSet := []Digit{
		{Label: 2, Features: []float64{2, 1}},
		{Label: 2, Features: []float64{2, 2}},
		{Label: 2, Features: []float64{1, 2}},

		{Label: 3, Features: []float64{2, 1}},
		{Label: 3, Features: []float64{2, 2}},
		{Label: 3, Features: []float64{1, 1}},

		{Label: 8, Features: []float64{2, 2}},
		{Label: 8, Features: []float64{3, 2}},
		{Label: 8, Features: []float64{2, 3}},
		{Label: 8, Features: []float64{2, 2}},

		{Label: 1, Features: []float64{0, 2}},
		{Label: 1, Features: []float64{0, 3}},
	}

	// Simulated input to classify
	newDigitFeatures := []float64{0.1, 2.5} // "close" to digit 1

	k := 3
	predicted := predictDigit(trainingSet, newDigitFeatures, k)
	fmt.Printf("Predicted digit: %d\n", predicted)
}

// ==============================================
// ⏱ Time Complexity:
// - Distance computation: O(n * d)
// - Sorting: O(n log n)
// - Voting: O(k)
// - Total: O(n * d + n log n)
//
// 💾 Space Complexity:
// - O(n) for storing distances
// - O(n * d) for training data
//
// ✅ Summary:
// This is a toy example. In real OCR, feature extraction is complex,
// and KNN is often replaced by neural networks (CNNs) for better performance.
// Still, this illustrates the fundamental principle clearly.
//
// ==============================================
