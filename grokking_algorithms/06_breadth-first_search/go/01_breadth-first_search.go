package main

import (
	"fmt"
	"slices"
	"strings"
)

/*
Breadth-First Search (BFS) – Find Mango Seller Example

BFS is part of graph traversal problems. It finds the shortest path (number of steps/hops) from start to some/all nodes.
It works on unweighted graphs. Graph can be (un)directed or disconnected. We handle cycles by skipping already visited nodes.

Problem:
You start with a person and want to find the nearest "mango seller" in your social network.
Each person is connected to other people (represented as a directed graph).
A mango seller is defined as a person whose name ends with 'm' (arbitrary).

Approach (BFS):
- Use a queue to explore the graph level by level (FIFO)
- Use a visited map to avoid processing the same person twice
- Optionally use a cameFrom map to reconstruct the shortest path

Why BFS:
- BFS always finds the shortest path (fewest steps) in an unweighted graph
- It's used in social networks, routing systems, friend suggestion, etc.

Time Complexity:
- Each person is processed at most once → O(V)
- Each edge is explored once → O(E)
- Total time complexity: O(V + E)

Answers the following questions:
- Can we reach node X?
- How many hops to reach X?
- What is the path to reach X?
- Reach all nodes → used for grid fill (flood fill, 01 matrix,...)
*/

func main() {
	graph := map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"bob":    {"anuj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {}, // mango seller
		"jonny":  {},
	}

	// you -> claire -> thom
	// distance from start: 2
	firstMangoSeller, distanceFromStart := search(graph, "you")
	fmt.Printf("Mango seller %s found, distance from start: %d", firstMangoSeller, distanceFromStart)
}

// personIsSeller returns true if the name ends with 'm'
func personIsSeller(name string) bool {
	return strings.HasSuffix(name, "m")
}

type NodeDist struct {
	node string // node label
	dist int    // distance from start (number of hops)
}

// search performs BFS starting from `start` and returns the first mango seller found along with its distance from start
func search(graph map[string][]string, start string) (string, int) {
	queue := NewQueue[NodeDist]()
	visited := make(map[string]bool)
	parent := make(map[string]string)

	queue.Enqueue(NodeDist{
		node: start,
		dist: 0,
	})

	for !queue.IsEmpty() {
		personPtr := queue.Dequeue()
		if personPtr == nil { // just in case
			continue
		}
		person := *personPtr

		// we enqueue each person (vertex/node) at most once (visited check)
		// this gives us O(V) time complexity
		if visited[person.node] {
			// if person has already been processed/visited, we skip it
			// this prevents cycles (infinite loops) and duplicate processing and improves our time complexity
			continue
		}

		// mark this person as processed/visited
		visited[person.node] = true

		// if person is seller - we have our solution, we can optionally print the path
		if personIsSeller(person.node) {
			printPath(person.node, parent)
			return person.node, person.dist
		}

		// if the person is not a seller, enqueue their neighbors (friends)
		// so for each person we traverse its neighbors once → O(E) edge traversal
		// so in total we have O(V + E)
		for _, neighbor := range graph[person.node] {
			if !visited[neighbor] {
				nd := NodeDist{
					node: neighbor,
					dist: person.dist + 1, // neighbour is +1 (compared to current node) away from start
				}
				queue.Enqueue(nd)
				parent[neighbor] = person.node
			}
		}
	}

	return "", 0
}

// printPath reconstructs and prints the path from start to the found person
func printPath(end string, parent map[string]string) {
	var path []string
	for person := end; person != ""; person = parent[person] {
		path = append(path, person)
	}
	slices.Reverse(path)
	fmt.Println("Shortest path:", strings.Join(path, " -> "))
}

// Generic FIFO Queue
type Queue[T any] struct {
	arr []T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) Enqueue(el T) {
	q.arr = append(q.arr, el)
}

func (q *Queue[T]) Dequeue() *T {
	if len(q.arr) == 0 {
		return nil
	}
	first := q.arr[0]
	q.arr = q.arr[1:]
	return &first
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.arr) == 0
}
