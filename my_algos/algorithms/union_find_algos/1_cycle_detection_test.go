package union_find_algos

import (
	"fmt"
	"panca.com/algo/union_find"
	"testing"
)

/*
=======================================================
🧠 Cycle Detection in an Undirected Graph (Union-Find)
=======================================================

🔷 PROBLEM:
Given an undirected graph with 'n' nodes and a list of edges,
determine if the graph contains a cycle.

🔷 CYCLE DEFINITION:

In an UNDIRECTED graph:
- A cycle is a path that starts and ends at the same node
- It must have at least 3 nodes
- Traversing the same edge back and forth (like A—B—A) is NOT a cycle

In a DIRECTED graph:
- A cycle can have just 2 nodes (e.g., A → B → A)

🔷 UNION-FIND STRATEGY:
For each edge (u, v):
1. Use `find(u)` and `find(v)` to check their set leaders (roots).
2. If `find(u) == find(v)`, they are already connected (part of same group/set) → a cycle exists.
3. Otherwise, perform `union(u, v)` to merge their sets.

This uses:
✅ Path compression in `find()` — flattens trees to improve speed
✅ Union by rank in `union()` — prevents deep trees by attaching smaller trees under larger ones

⏱ TIME COMPLEXITY: O(n + E * α(n))
- O(n) to initialize parent and rank arrays (if solved using array based UF)
- Each edge takes O(α(n)) time due to path compression
	- α(n) is the inverse Ackermann function (almost constant time)
- Total = O(n + E * α(n)), where E is number of edges

🛋 SPACE COMPLEXITY: O(n)
- for parent and rank arrays of size n (if solved using array based UF) or for maps for generic UF

*/

func containsCycle(edges [][]int) bool {
	uf := union_find.NewUnionFind[int]()

	// Process each edge
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if !uf.Union(u, v) {
			return true // Found a cycle
		}
	}

	// No cycles found
	return false
}

// containsCycleOtherWay is a version which would use array based Union-find solution with parent and rank arrays
// which need to be initialized in the beginning and they need to be passed around to Union/Find calls.
// returns true if cycle is found, otherwise false.
func containsCycleOtherWay(n int, edges [][]int) bool {
	parent := make([]int, n)
	rank := make([]int, n)

	// Initialize: each node is its own parent and rank is zero
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	// Process each edge
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if !union_find.Union(u, v, parent, rank) {
			// Found a cycle
			return true
		}
	}

	// No cycles found
	return false
}

func Test_containsCycle(t *testing.T) {
	// Test case 1: No cycle
	n1 := 5
	edges1 := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
	fmt.Println("Contains cycle (should be false):", containsCycle(edges1))                       // false
	fmt.Println("Contains cycle other way (should be false):", containsCycleOtherWay(n1, edges1)) // false

	// Test case 2: Cycle exists (0-1-2-0)
	n2 := 4
	edges2 := [][]int{{0, 1}, {1, 2}, {2, 0}}
	fmt.Println("Contains cycle (should be true):", containsCycle(edges2))                       // true
	fmt.Println("Contains cycle other way (should be true):", containsCycleOtherWay(n2, edges2)) // true
}
