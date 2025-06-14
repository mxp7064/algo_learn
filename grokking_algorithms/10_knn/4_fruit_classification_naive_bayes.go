package main

import (
	"fmt"
	"math"
)

// This is rearely directly asked in DSA interview, but it's good to know in general.

// Naive Bayes Classifier - Fruit Classification Example
// -----------------------------------------------------
// Problem Statement:
// We are given a simple dataset of fruits labeled as either "grapefruit" or "orange".
// Each fruit has two features:
//  - color (redness on a scale from 0 to 1)
//  - size (on a scale from 0 to 1)
// Given a new fruit with observed features, predict whether it's a grapefruit or an orange
// using the Naive Bayes classification algorithm.

// Naive Bayes Theorem Summary:
// ------------------------------
// Bayes' theorem lets us compute (∝ means proportional):
//     P(class | features) ∝ P(features | class) * P(class)
// In our case:
//     P(grapefruit | red, big) ∝ P(red | grapefruit) * P(big | grapefruit) * P(grapefruit)
// In words: The probability that a fruit is a grapefruit, given that it's red and big, is proportional to the
// probability that a fruit is red given it's a grapefruit, multiplied by the probability that it's big given
// it's a grapefruit, multiplied by the overall probability of a fruit being a grapefruit.
// Same goes for orange:
//     P(orange    | red, big) ∝ P(red | orange)    * P(big | orange)    * P(orange)
// We assume that features are conditionally independent given the class:
//     P(red and big | grapefruit) = P(red | grapefruit) * P(big | grapefruit)
// Independent means the presence of one feature (e.g., being red) does not affect the probability of another feature
// (e.g., being big), given the class. This is the "naive" assumption.
// In reality, features often aren’t truly independent (e.g., big fruits are often redder, etc.), but this
// simplification still works well in practice and allows for fast and effective classification

// Time complexity:
// - Training: O(n), where n is the number of training samples
// - Prediction: O(k), where k is the number of features
// (since we compute P(feature | class) for each feature per class)

// Notes:
// - This simplified implementation assumes discrete features (rounded to nearest 0.0/0.5/1.0 etc)
// - For real-world continuous features, you would use probability distributions (e.g., Gaussian)

type Fruit struct {
	Color float64 // 0.0 = orange, 1.0 = red
	Size  float64 // 0.0 = small,  1.0 = big
	Label string  // "orange" or "grapefruit"
}

func roundToHalf(x float64) float64 {
	return math.Round(x*2) / 2.0 // round to nearest 0.5 for simplicity
}

func main() {
	// Our training dataset (features + labels)
	trainingData := []Fruit{
		{0.1, 0.2, "orange"},
		{0.2, 0.3, "orange"},
		{0.1, 0.4, "orange"},
		{0.9, 0.8, "grapefruit"},
		{0.8, 0.9, "grapefruit"},
		{0.7, 0.9, "grapefruit"},
	}

	// Test fruit we want to classify
	newFruit := Fruit{Color: 0.75, Size: 0.85}

	// Step 1: Count occurrences and build conditional probabilities
	labelCounts := make(map[string]int)
	colorGivenLabel := make(map[string]map[float64]int)
	sizeGivenLabel := make(map[string]map[float64]int)

	for _, fruit := range trainingData {
		labelCounts[fruit.Label]++

		color := roundToHalf(fruit.Color)
		size := roundToHalf(fruit.Size)

		if colorGivenLabel[fruit.Label] == nil {
			colorGivenLabel[fruit.Label] = make(map[float64]int)
		}
		if sizeGivenLabel[fruit.Label] == nil {
			sizeGivenLabel[fruit.Label] = make(map[float64]int)
		}

		colorGivenLabel[fruit.Label][color]++
		sizeGivenLabel[fruit.Label][size]++
	}

	totalExamples := len(trainingData)

	// Step 2: Apply Naive Bayes formula to calculate scores for each class
	scores := make(map[string]float64)
	observedColor := roundToHalf(newFruit.Color)
	observedSize := roundToHalf(newFruit.Size)

	for label := range labelCounts {
		prior := float64(labelCounts[label]) / float64(totalExamples)
		colorProb := float64(colorGivenLabel[label][observedColor]+1) / float64(labelCounts[label]+2) // Laplace smoothing
		sizeProb := float64(sizeGivenLabel[label][observedSize]+1) / float64(labelCounts[label]+2)
		scores[label] = prior * colorProb * sizeProb
		fmt.Printf("Score for %s: %.4f\n", label, scores[label])
	}

	// Step 3: Pick the label with the highest score
	bestLabel := ""
	bestScore := 0.0
	for label, score := range scores {
		if score > bestScore {
			bestScore = score
			bestLabel = label
		}
	}

	fmt.Printf("\nPredicted label for test fruit: %s\n", bestLabel)
}
