/*
COUNT CONNECTED COMPONENTS – UNION-FIND

🧠 PROBLEM:
Given `n` nodes labeled 0..n-1 and a list of undirected edges, return the number of connected components in the graph.

💡 INTUITION:
- Think of each node as a separate island (a component).
- Initially, none of the islands are connected → n separate components.
- Each edge is a bridge that connects two islands.
- When we build a bridge between 2 nodes/islands, we create a bigger connected landmass.
- If the two islands are already connected (already part of same landmass/component), ignore the bridge.
- If not, the bridge merges two landmasses → total component count drops by 1.
	- You went from having 2 separate components (compA and compB) → to one larger merged component (compAB)
	- That's why the total number of components decreases by 1.
- Final count = number of remaining disconnected landmasses.

✅ APPROACH:
Set initial componentCount to n (all components are separated - there are n distinct components)
For each edge [u, v]:
- use Find(u) and Find(v) to check if they’re already connected
- if Find(u) != Find(v):
	- they’re in different components → this edge truly connects them
	- call Union(u, v) to merge their sets
	- decrement count: componentCount--
- if they’re already connected (Find(u) == Find(v)), do nothing — this edge is redundant - number of components stays the same

⏱ TIME COMPLEXITY: O(n + E * α(n))
- O(n) to initialize parent and rank arrays (if solved using array based UF)
- Each edge takes O(α(n)) time due to path compression
	- α(n) is the inverse Ackermann function (almost constant time)
- Total = O(n + E * α(n)), where E is number of edges

📦 SPACE COMPLEXITY: O(n)
- Parent and rank arrays of size n (if solved using array based UF) or for maps for generic UF

🌍 REAL-WORLD USES:
- Dynamic clustering, friend circles, social network analysis
- Network segmentation, connectivity diagnostics, merging partitions

🧪 LeetCode problems solved by this logic:
- 323. Number of Connected Components in an Undirected Graph
- 547. Number of Provinces (4_number_of_provinces_test.go)
- 1319. Number of Operations to Make Network Connected
*/

package union_find_algos

import (
	"fmt"
	"panca.com/algo/union_find"
	"testing"
)

func countConnectedComponents(n int, edges [][]int) int {
	uf := union_find.NewUnionFind[int]()
	componentCount := n
	for _, edge := range edges { // O(E * α(n))
		u := edge[0]
		v := edge[1]

		// If u and v are in different components, merge them
		// Union returns true only if merge happened (u, v were in different sets)
		// Union has time complexity of O(α(n))
		if uf.Union(u, v) { // same as Find(u) != Find(v) ...
			componentCount-- // two components merged → total count decreases
		}
	}

	return componentCount
}

func Test_countConnectedComponents(t *testing.T) {
	cases := []struct {
		n      int
		edges  [][]int
		expect int
	}{
		{
			n: 5,
			edges: [][]int{
				{0, 1}, {1, 2}, {3, 4},
			},
			expect: 2, // [0-1-2], [3-4]
		},
		{
			n: 5,
			edges: [][]int{
				{0, 1}, {1, 2}, {2, 3}, {3, 4},
			},
			expect: 1, // fully connected (all islands part of same landmass)
		},
		{
			n:      4,
			edges:  [][]int{},
			expect: 4, // no edges → 4 isolated components
		},
		{
			n: 3,
			edges: [][]int{
				{0, 1}, {1, 2}, {0, 2}, // redundant edge {0,2}
			},
			expect: 1,
		},
	}

	for i, c := range cases {
		res := countConnectedComponents(c.n, c.edges)
		if res != c.expect {
			t.Errorf("Case %d failed: got %d, want %d", i, res, c.expect)
		} else {
			fmt.Printf("Case %d passed → %d components\n", i, res)
		}
	}
}
