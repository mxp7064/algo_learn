package main

import (
	"fmt"
	"math"
	"sort"
)

// K-Nearest Neighbors (KNN) Regression
//
// Problem Statement:
// Suppose you have a dataset of fruits (or objects) described by their features (size and color) similar
// as in the previous example for classification
// Each fruit has a size and color (in simplified numeric form), and you also know its weight.
// Given a new fruit, predict its weight based on the weights of the k most similar (nearest) fruits.
//
// This example demonstrates KNN used for regression, not classification.
//
// NOTE:
// In our previous example (orange vs. grapefruit classification), we used majority voting
// to decide whether a fruit is either an orange or a grapefruit - our final prediction/result was just a label -
// either "orange" or "grapefruit"
// In this case, we are predicting a numerical value (weight), so we use the average
// of the target values of the k nearest neighbors.
//
// In summary:
// - KNN Classification → Choose the most common label among k nearest neighbors.
// - KNN Regression → Take the average of the numeric values (e.g. weight) of the k nearest neighbors.
//
// Time Complexity: O(n log n) for sorting the distances (with n being the number of training points)
// Space Complexity: O(n) for storing distances

type Fruit struct {
	Size   float64 // Feature 1
	Color  float64 // Feature 2 (e.g. 0 = yellow, 1 = red)
	Weight float64 // This is the value we want to predict
}

// Euclidean distance between two fruits
func distance(a, b Fruit) float64 {
	dx := a.Size - b.Size
	dy := a.Color - b.Color
	return math.Sqrt(dx*dx + dy*dy)
}

// Predict the weight of a fruit using KNN regression
func knnPredictWeight(data []Fruit, input Fruit, k int) float64 {
	// Calculate distances between input and all data points
	type neighbor struct {
		fruit    Fruit
		distance float64
	}
	var distances []neighbor
	for _, fruit := range data {
		d := distance(input, fruit)
		distances = append(distances, neighbor{fruit, d})
	}

	// Sort neighbors by distance (ascending)
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	// Average the weights of the k nearest neighbors
	var totalWeight float64
	for i := 0; i < k && i < len(distances); i++ {
		totalWeight += distances[i].fruit.Weight
	}
	return totalWeight / float64(k) // result/prediction is the average of weights of k nearest neighbors
}

func main() {
	// Training data: each fruit has size, color and known weight
	training := []Fruit{
		{Size: 5.0, Color: 0.0, Weight: 120}, // small yellow fruit
		{Size: 6.0, Color: 1.0, Weight: 150}, // small red fruit
		{Size: 7.5, Color: 1.0, Weight: 180}, // large red fruit
		{Size: 8.0, Color: 1.0, Weight: 200}, // very large red fruit
		{Size: 6.5, Color: 0.0, Weight: 160}, // medium yellow fruit
	}

	// New fruit for which we want to predict weight
	input := Fruit{Size: 7.0, Color: 0.8}

	k := 3
	predictedWeight := knnPredictWeight(training, input, k)
	fmt.Printf("Predicted weight of the new fruit is %.2f grams\n", predictedWeight)
}
