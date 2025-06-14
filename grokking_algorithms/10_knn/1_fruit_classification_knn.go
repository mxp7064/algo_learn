package main

import (
	"fmt"
	"math"
	"sort"
)

// K-Nearest Neighbors (KNN) Classification

// Problem: Classify a fruit as either an "orange" or "grapefruit" based on two features: size and redness.
// This is a typical example of k-nearest neighbors (KNN) algorithm
// Grapefruit is bigger and more red while orange is smaller and more orange
// Training data includes labeled fruits with known size and redness values.
// We classify a new fruit by finding the k nearest labeled fruits and using majority voting.

// Time Complexity: O(n log n) due to sorting distances
// Space Complexity: O(n) to store distances

type Fruit struct {
	Size    float64
	Redness float64
	Label   string
}

type Neighbor struct {
	Distance float64
	Label    string
}

// Euclidean distance in 2D feature space
func distance(a, b Fruit) float64 {
	dx := a.Size - b.Size
	dy := a.Redness - b.Redness
	return math.Sqrt(dx*dx + dy*dy)
}

// KNN classifier
func classify(trainingSet []Fruit, input Fruit, k int) string {
	// calculate distance between the input and all the data points in the training set
	var neighbors []Neighbor
	for _, train := range trainingSet {
		d := distance(input, train)
		neighbors = append(neighbors, Neighbor{Distance: d, Label: train.Label})
	}

	// sort all neighbors based on distance (ascending)
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].Distance < neighbors[j].Distance
	})

	// count number of occurrences for first k nearest neighbours
	voteCount := make(map[string]int)
	for i := 0; i < k; i++ { // take first k neighbours
		voteCount[neighbors[i].Label]++ // for each we count the number of occurrences
	}

	// Check which neighbour is in majority (most occurrences), i.e. are there more oranges or grapefruits
	// among the k nearest neighbours
	maxVotes := 0
	bestLabel := ""
	for label, votes := range voteCount {
		if votes > maxVotes {
			bestLabel = label
			maxVotes = votes
		}
	}
	return bestLabel
}

func main() {
	trainingSet := []Fruit{
		{Size: 3.0, Redness: 1.0, Label: "orange"},
		{Size: 2.5, Redness: 0.8, Label: "orange"},
		{Size: 3.5, Redness: 0.9, Label: "orange"},
		{Size: 4.5, Redness: 1.5, Label: "grapefruit"},
		{Size: 4.0, Redness: 1.8, Label: "grapefruit"},
		{Size: 5.0, Redness: 1.7, Label: "grapefruit"},
	}

	// New fruit to classify
	newFruit := Fruit{Size: 4.0, Redness: 1.4}

	result := classify(trainingSet, newFruit, 3)
	fmt.Printf("The classified fruit is likely: %s\n", result)
}
