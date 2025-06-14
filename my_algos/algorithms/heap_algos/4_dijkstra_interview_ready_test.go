/*
DIJKSTRA'S ALGORITHM – HEAP-BASED VERSION (INTERVIEW-READY)

🟩 PROBLEM STATEMENT:
You are given a weighted graph with non-negative edge weights, find the shortest path from a given start node to a finish node.

🧠 GRAPH TYPE:
- weights must be non-negative (the only important thing!)
- can be directed or undirected, in case of undirected we treat each undirected edge as two directed edges with same weight
- can have cycles, it won't cause issues because algorithm avoids revisiting nodes and goes for shortest known distance
- can be disconnected - just some nodes will never be reached (will remain with dist[node] == INF)

📥 INPUT:
- A graph represented as an adjacency list: map[string][]Edge
- A starting node (string)
- A finish node (string)

📤 OUTPUT:
- The length of the shortest path from start to finish
- The actual path from start to finish as a slice of strings (optional)

✅ REAL-WORLD USE CASES:
- GPS navigation (shortest route)
- Network routing (minimum latency)
- Game AI pathfinding

🧠 CORE IDEA (INTUITION):

Dijkstra’s algorithm always expands to the next closest (cheapest) node that hasn’t been processed yet.
As it runs, it keeps improving the shortest known distances to every node. Dijkstra is an example of greedy algorithm.

Instead of scanning all nodes for the next closest (which is O(n)), we use a min-heap to pick the best candidate efficiently.

📊 DATA STRUCTURES:
- graph: map[string][]Edge → adjacency list
- dist: map[string]int → shortest distances from start
- parents: map[string]string → for path reconstruction
- visited: map[string]bool → finalized nodes
- heap: min-heap of Items → to process the next closest node

🔻 WHY USE A MIN-HEAP IN DIJKSTRA'S ALGORITHM:

At each step of Dijkstra’s algorithm, we need to determine which unprocessed node has the smallest known distance from the start.

A min-heap allows us to:
- Efficiently retrieve the node with the smallest distance from start in O(log V) time
- Avoid scanning all nodes linearly (which would be O(V))
- Dynamically update the frontier of discovered but unprocessed nodes

When a node is popped from the heap:
- It is guaranteed to have the shortest path from the start
- We mark it as visited and never process it again
- It becomes part of the shortest path tree

Each time we visit neighbors, we only insert a node into the heap if we found a better path to it.
This keeps the heap efficient and ensures we always expand the closest/cheapest known node next.
The heap acts as a dynamic priority queue of “best current candidates” and we always pick the best candidate with heap.Pop()
in each iteration so that is how we build the best shortest path.

⏱️ TIME COMPLEXITY:

Let:
- V = number of vertices (nodes)
- E = number of edges

Using a min-heap:
- Pop and Insert cost O(log V) because we will have at most V nodes in the heap in the worst case
- We perform at most one Pop per node → O(V log V)
- We may insert into the heap once per edge (when exploring neighbors) → O(E log V)

These operations are independent, so we add their costs:
➡️ Total: O(V log V + E log V) = O((V + E) log V)

🟡 Comparison to Grokking version without heap (07_dijkstras_algorithm/go/01_dijkstras_algorithm.go):
If we used a linear scan to find the closest unprocessed node as in Grokking version that would take O(V) per node, leading to O(V²) total time.

Using a min heap

🔁 WHY EARLY EXIT WORKS:cost
As soon as we pop the `finish` node from the heap, we’re guaranteed to have found the shortest path to it.
So we can return early instead of computing distances to all nodes.

*/

package heap_algos

import (
	"fmt"
	"math"
	"panca.com/algo/myheap"
	"slices"
	"strings"
	"testing"
)

// Edge represents a connection from one node to another with a weight
type Edge struct {
	To     string
	Weight int
}

// Item represents an entry in the heap with node and current shortest known distance
type Item struct {
	Node     string
	Distance int
}

func Test_Dijkstra(t *testing.T) {
	// graph as an adjacency list: key is the node and value is list of edges from that node
	graph := map[string][]Edge{
		"A": {{"B", 5}, {"C", 1}},
		"C": {{"B", 2}, {"D", 4}},
		"B": {{"D", 1}},
		"D": {},
	}
	distance, path := Dijkstra("A", "D", graph)
	fmt.Println("Shortest distance:", distance)      // 4
	fmt.Println("Path:", strings.Join(path, " -> ")) // A -> C -> B -> D
}

// Dijkstra returns the shortest distance and path from start to finish in a weighted graph
// Time Complexity: O((V + E) log V) using a min-heap for node selection
func Dijkstra(start string, finish string, graph map[string][]Edge) (int, []string) {
	dist := make(map[string]int) // key is node, value is shortest known distance from start to that node
	for node := range graph {
		dist[node] = math.MaxInt // initialize all distances to infinity
	}
	dist[start] = 0 // distance from start node to start node is zero

	visited := make(map[string]bool)   // tracks processed nodes to avoid reprocessing
	parents := make(map[string]string) // used to reconstruct shortest path

	// min heap of Items, item with smallest distance is root
	heap := myheap.NewHeap(func(a, b Item) bool {
		return a.Distance < b.Distance
	})
	heap.Insert(Item{Node: start, Distance: 0})

	for heap.Len() > 0 {
		// heap.Pop() returns the closest node (node with smallest distance from start) among the encountered (by
		// neighbour exploration) unprocessed nodes. This node is part of the final shortest path
		current := heap.Pop()
		node := current.Node

		// Skip already processed nodes
		if visited[node] {
			continue
		}
		visited[node] = true

		// Early exit if finish node is reached
		if node == finish {
			path := buildPath(finish, parents)
			return dist[finish], path
		}

		// Visit neighbors and update distances
		for _, edge := range graph[node] {
			neighbour := edge.To
			distanceThroughNode := dist[node] + edge.Weight // distance to neighbour if we go through this node

			if distanceThroughNode < dist[neighbour] { // if going through this node is better, update distance and insert into heap
				dist[neighbour] = distanceThroughNode

				// store path for later reconstruction
				// map should only update if you actually improved the distance - we want to reconstruct the shortest path later
				// and initial distances are infinity, we don't want to overwrite a better path with a worse one
				parents[neighbour] = node

				// We discovered a better path to this neighbor
				// You may push a node into the heap multiple times with different distances
				// But you only process it once, the first time it’s popped (when it has the shortest known distance)
				// visited map will stop node form being processed more than once
				heap.Insert(Item{Node: neighbour, Distance: distanceThroughNode})
			}
		}
	}

	// Finish node not reachable
	return -1, nil
}

// buildPath reconstructs the shortest path using the parents map
func buildPath(finish string, parents map[string]string) []string {
	var path []string
	current := finish
	for current != "" {
		path = append(path, current)
		current = parents[current]
	}
	slices.Reverse(path)
	return path
}
