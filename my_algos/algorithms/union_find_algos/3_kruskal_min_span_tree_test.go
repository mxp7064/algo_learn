/*
KRUSKAL'S ALGORITHM – MINIMUM SPANNING TREE (MST)

🧠 PROBLEM:
Given a connected, undirected, weighted graph (nodes labeled 0..n-1 and edges [u, v, w] where w is weight),
return the minimum total weight of a set of edges that connects all nodes without forming cycles.

💡 INTUITION:
- Think of each node as a city, and each weighted edge as a possible road.
- You want to build roads to connect all cities with minimal cost (minimize total road length).
- But you must avoid cycles (no wasted roads).
- This means your result must be a tree (connected + acyclic).

📐 WHY IS IT CALLED A "TREE"?
In graph theory, a tree is:
- A connected, undirected, acyclic graph
- Has exactly `n - 1` edges for `n` nodes
- No root is required (not like binary trees)
- Kruskal guarantees this by:
  - Connecting different components (with Union-Find)
  - Skipping edges that form cycles

🔄 HOW KRUSKAL WORKS:
1. Sort all edges by weight
2. Initialize Union-Find structure
3. For each edge [u, v, w]:
   - If u and v are not already connected:
     - Add the edge to MST
     - Merge the two components
   - Else skip it (would form a cycle)
4. When you've added (n - 1) edges or when you process all edges without forming cycles, MST is complete

🔢 TIME COMPLEXITY:
- Sorting: O(E log E), where E is number of edges
- Union-Find: O(E * α(N)) → almost constant
- Total: O(E log E) because sorting dominates

🛋 SPACE COMPLEXITY:
- O(n) for parent and rank arrays of size n (if solved using array based UF) or for maps for generic UF

📆 REAL-WORLD USES:
- Network design (fiber optics, electrical grids)
- Cluster analysis
- Approximation algorithms (e.g., traveling salesman)

🧪 LeetCode examples:
- 1135. Connecting Cities With Minimum Cost
- 1584. Min Cost to Connect All Points
*/

package union_find_algos

import (
	"fmt"
	"panca.com/algo/union_find"
	"sort"
	"testing"
)

func getMstCost(edges [][]int, n int) int {
	totalCost := 0

	// Sort edges by increasing weight, O(E log E)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	uf := union_find.NewUnionFind[int]()

	// Process edges greedily by weight
	for _, edge := range edges { // O(E * α(N))
		u, v, w := edge[0], edge[1], edge[2]
		if uf.Union(u, v) { // Union is O(α(N))
			totalCost += w
		}
	}

	return totalCost
}

func Test_getMstCost(t *testing.T) {
	edges := [][]int{
		{0, 1, 4},
		{0, 2, 4},
		{1, 2, 2},
		{1, 0, 4}, // duplicate (undirected)
		{2, 3, 3},
		{2, 5, 2},
		{2, 4, 4},
		{3, 4, 3},
		{5, 4, 3},
	}
	n := 6
	fmt.Println(getMstCost(edges, n)) // Expected: 14
}
