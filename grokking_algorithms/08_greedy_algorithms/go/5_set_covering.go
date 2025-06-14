/*
SET COVERING – GREEDY APPROXIMATION (INTERVIEW-READY)

🧠 PROBLEM:
You are given:
- A set of states that need coverage like ["ny", "alaska", ...]
- A list of broadcasting stations, each of which covers a subset of those states.

📌 GOAL:
Find the minimum number of stations such that all required states are covered.

⚠️ NOTE:
This is an approximation algorithm. It does not guarantee the optimal solution,
but it runs efficiently and gives good results in practice.

🔍 WHY NOT BRUTE FORCE?
- Brute force (optimal solution) would check all combinations of stations (2^n subsets) to find the smallest one
that covers all states.
- Time complexity: O(2^n * m), where m is the number of states.
	- In the book this is wrongly written as O(n!)
- Completely infeasible for even moderate inputs.

✅ GREEDY STRATEGY:
- At each step, select the station that covers the largest number of uncovered states.
- Remove those states from the list of required states.
- Repeat until all states are covered.

🧮 TIME COMPLEXITY:
- Let n = number of stations, m = number of states
- Outer loop runs up to m times (until all states are covered)
- Inner loop compares coverage for each station → O(n)
- So total: O(m * n) = usually written as O(n^2) for clarity and simplicity

🛠 SET OPERATIONS USED:
- intersection(a, b): returns states covered by a station that are still uncovered
- difference(a, b): removes newly covered states from uncovered set
- both operations are O(k), where k is number of states in the station’s set — typically small, so we treat it as constant (ignore it) for simplicity
*/

package main

import "fmt"

// StationMap represents a mapping from station name to the set of states it covers
type StationMap map[string]Set[string]

func main() {
	// Set of all states we want to cover
	requiredStates := NewSet("ny", "alaska", "florida", "california", "virginia")

	// Map of stations and which states they cover
	stations := StationMap{
		"kvm": NewSet("ny", "florida"),
		"abc": NewSet("california", "virginia", "ny"),
		"hhh": NewSet("ny", "virginia", "alaska", "florida"),
		"xyz": NewSet("alaska"),
	}

	selectedStations := FindMinimalStationCoverage(stations, requiredStates)
	fmt.Println("Selected stations:", selectedStations.GetElements()) // [abc hhh]
}

// FindMinimalStationCoverage applies a greedy strategy to select the smallest set of stations
// that collectively cover all required states.
func FindMinimalStationCoverage(stations StationMap, requiredStates Set[string]) Set[string] {
	// Copy the station map so the original remains unchanged
	stationsLeft := make(StationMap)
	for station, stateSet := range stations {
		stationsLeft[station] = stateSet
	}

	result := NewSet[string]() // result set of station names

	// if m is number of states, then we have O(m) for this loop cause in worst case we will cover only one state in each iteration
	for requiredStates.Len() > 0 { // while we have states to cover...
		var bestStation string   // station that covers the most uncovered states
		var maxStatesCovered int // number of states that it covers

		// Find the station that covers the most uncovered states
		for station, states := range stationsLeft { // O(n) if n is number of stations
			coveredStates := intersection(states, requiredStates) // uncovered states that this station covers
			if coveredStates.Len() > maxStatesCovered {
				// chose the station which covers highest number of uncovered states
				maxStatesCovered = coveredStates.Len()
				bestStation = station
			}
		}

		// Add best station to result
		result.Add(bestStation)

		requiredStates = difference(requiredStates, stationsLeft[bestStation]) // remove states that we covered
		delete(stationsLeft, bestStation)                                      // remove station that we picked
	}

	return result
}

/* ------------------ SET IMPLEMENTATION ------------------ */

// Set is a generic set implemented as a map[T]bool
type Set[T comparable] map[T]bool

// NewSet constructs a set from a list of elements
func NewSet[T comparable](elements ...T) Set[T] {
	s := make(Set[T])
	for _, el := range elements {
		s[el] = true
	}
	return s
}

func (s Set[T]) Add(el T) {
	s[el] = true
}

func (s Set[T]) Delete(el T) {
	delete(s, el)
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) GetElements() []T {
	var result []T
	for el := range s {
		result = append(result, el)
	}
	return result
}

/* ------------------ SET OPERATIONS ------------------ */

// intersection returns a ∩ b: elements present in both sets
func intersection[T comparable](a, b Set[T]) Set[T] {
	result := make(Set[T])
	for el := range a {
		if b[el] {
			result[el] = true
		}
	}
	return result
}

// difference returns a - b: elements in a that are not in b
func difference[T comparable](a, b Set[T]) Set[T] {
	result := make(Set[T])
	for el := range a {
		if !b[el] {
			result[el] = true
		}
	}
	return result
}
