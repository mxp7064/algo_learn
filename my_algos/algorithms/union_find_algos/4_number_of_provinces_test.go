/*
NUMBER OF PROVINCES – LeetCode 547

We also solved this same problem using dfs (comparison can be found there): backtrack_recursion_dfs/6_number_of_provinces.go

🫠 PROBLEM:
Given an n x n matrix `isConnected`, where:
- `isConnected[i][j] == 1` means person i is directly connected to person j
- `isConnected[i][j] == 0` means no direct connection
- Persons are labeled from 0 to n-1
Return the number of provinces, where a province is a group of directly or indirectly connected people.

🔄 MODELING:
This is a classic connected components problem.
- Each person is a node
- A direct connection (1) between `i` and `j` is an undirected edge [i, j]
- This is the same problem as counting the number of connected components in an undirected graph — the only difference is
that the graph is given as an adjacency matrix instead of a list of edges.

🔄 MATRIX PROPERTIES:
- The matrix is always square (n x n)
	- because we have n persons and we need to express all of their relationships
- It is symmetric: isConnected[i][j] = isConnected[j][i]
- The diagonal is always `1` (every person is connected to themselves)

🔍 STRATEGY:
1. Use Union-Find to group connected nodes
2. Traverse only the upper triangle of the matrix
	- avoid duplicate edges ([i, j] = [j, i])
	- avoid edge to the person itself ([i, i]) because that is redundant information
3. Start with `n` components, and decrement on every successful union

📊 TIME COMPLEXITY:
- Matrix traversal: O(n^2)
- Union-Find operations: O(α(n))
- Total: O(n^2 * α(n))

📅 SPACE COMPLEXITY:
- O(n) for parent and rank arrays of size n (if solved using array based UF) or for maps for generic UF

📆 REAL-WORLD ANALOGIES:
- Friend groups
- Social networks
- Cluster detection in social or biological systems
*/

package union_find_algos

import (
	"fmt"
	"panca.com/algo/union_find"
	"testing"
)

func getNumProvinces(isConnected [][]int) int {
	n := len(isConnected)
	numProvinces := n

	uf := union_find.NewUnionFind[int]()

	// Only iterate upper triangle (j > i) to avoid duplicate connections
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 && uf.Union(i, j) {
				numProvinces--
			}
		}
	}

	return numProvinces
}

func Test_getNumProvinces(t *testing.T) {
	isConnected1 := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}                                          // n = 3, we have persons 0, 1, and 2
	fmt.Println(getNumProvinces(isConnected1)) // Expected: 2

	isConnected2 := [][]int{
		{1, 0, 0, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 0, 1, 1},
	}                                          // n = 4, we have persons 0, 1, 2 and 3
	fmt.Println(getNumProvinces(isConnected2)) // Expected: 1
}
