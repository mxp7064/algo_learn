package main

/*
Check out algorithms/dijkstra_interview_ready_test.go for a more efficient and interview ready version using min heap
*/

import (
	"fmt"
	"math"
	"strings"
)

// Find the lowest-cost node that hasn't been processed yet
func findLowestCostNode(costs map[string]float64, processed map[string]bool) string {
	lowestCost := math.Inf(1)
	lowestCostNode := ""

	// Go through each node
	for node, cost := range costs {
		// If it's the lowest cost so far and hasn't been processed yet...
		if cost < lowestCost && !processed[node] {
			// ... set it as the new lowest-cost node
			lowestCost = cost
			lowestCostNode = node
		}
	}
	if lowestCostNode == "" {
		return "" // No unprocessed nodes found
	}
	return lowestCostNode // b | a | fin
}

func main() {
	// The graph
	graph := map[string]map[string]float64{
		"start": {"a": 6, "b": 2},
		"a":     {"fin": 1},
		"b":     {"a": 3, "fin": 5},
		"fin":   {},
	}

	// The costs table
	infinity := math.Inf(1)
	costs := map[string]float64{
		"a":   6,
		"b":   2,
		"fin": infinity,
	}

	// The parents table
	parents := map[string]string{
		"a":   "start",
		"b":   "start",
		"fin": "",
	}

	// Keep track of processed nodes
	processed := make(map[string]bool)

	// Find the lowest-cost node that hasn't been processed yet
	node := findLowestCostNode(costs, processed) // b

	// If you've processed all the nodes, this while loop is done
	for node != "" {
		cost := costs[node] // 2
		// Go through all the neighbors of this node
		neighbors := graph[node]             // A, FIN
		for n, edgeCost := range neighbors { // edgeCost : 3 za A, 5 za FIN
			newCost := cost + edgeCost // za A 5, za FIN 7
			// If it's cheaper to get to this neighbor by going through this node...
			if newCost < costs[n] { // 5 < 6
				// ... update the cost for this node
				costs[n] = newCost // A -> 5, FIN 7
				// This node becomes the new parent for this neighbor
				parents[n] = node // A -> B
			}
		}
		// Mark the node as processed
		processed[node] = true
		// Find the next node to process, and loop
		node = findLowestCostNode(costs, processed) // a | fin | ""
	}

	fmt.Println("Cost from the start to each node:")
	fmt.Println(costs)
	path := []string{"fin"}
	parent := parents["fin"]
	for parent != "" {
		path = append([]string{parent}, path...)
		parent = parents[parent]
	}
	fmt.Println(strings.Join(path, " -> "))

}
