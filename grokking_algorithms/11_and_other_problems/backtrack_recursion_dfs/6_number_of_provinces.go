package main

import "fmt"

/*
======================================================================================
Number of Provinces — using DFS (Count Connected Components in Undirected Graph)
======================================================================================

This is LC 547. We also solved this same problem using Union-Find: algorithms/union_find_algos/4_number_of_provinces_test.go
Comparison between the two is here.

🔷 PROBLEM:

You are given an n x n adjacency matrix `isConnected`, where:
- isConnected[i][j] = 1 means there is a direct connection between city i and city j
- isConnected[i][j] = 0 means no direct connection

Cities are considered connected if they are directly or indirectly connected.
A province is a group of connected cities. Goal: return the number of provinces (connected components in the graph).

🔷 TYPE OF GRAPH:

- Undirected (matrix is symmetric)
- Unweighted
- Possibly disconnected
- No negative weights or cycles (but cycles are not a problem for DFS here)

🔷 WHY THIS WORKS:

This is a classic example of counting connected components in an undirected graph.
Each time we start DFS from an unvisited node (city), we explore all cities that belong to the same group.
We mark all reachable cities as visited.
Each time we start a new DFS from a new unvisited city, we have discovered a new connected component → a new province.

This works because DFS recursively visits all cities reachable from the starting city via edges.
Once a DFS finishes, all cities in that province have been visited.
We repeat this process until all cities have been visited.

🔷 IS THIS A STANDARD APPROACH?

Yes. DFS is one of the three most standard and accepted solutions for this problem.
The others ones are BFS and Union-Find (Disjoint Set Union).
All achieve the same result and are commonly accepted in interviews.

🔷 WHEN IS DFS EXPLICIT?

- In problems like this, DFS is applied directly to graph structure (adjacency list/matrix)
- This is different from problems like Fibonacci, subset generation, permutations etc., where the "graph" is
implicit (it emerges from recursion/decision tree, rather than being explicitly given in the problem)

🔷 ALGORITHM STEPS:

1. Convert matrix to adjacency list, we need to explore only lower/upper triangle of the matrix
2. Create a visited[] array/map
3. For each city i from 0 to n-1:
   - If not visited, start DFS and mark all reachable cities
   - Increment province count
4. Return total province count

🔷 DFS VS UNION-FIND COMPARISON:

- DFS: explore via recursion or stack, visit all neighbors
- Union-Find: treat each city as a set, merge sets on connection
- Both detect and count connected components
- Both run in O(n²) time due to matrix

🔷 DFS VS BFS COMPARISON:

Here we also solved this problem using BFS - the logic and time complexity is the same, we just use a different kind of traversal.

Goal:
- Traverse each connected component of the graph exactly once
- Count how many disconnected components exist → number of provinces

Why both work:
- Each province = connected group of cities
- DFS and BFS both traverse all reachable nodes from a given start node
- Every time you start a new traversal from an unvisited node, you’ve found a new province

Algorithm behaviour:
- DFS: go deep along one path before backtracking (uses recursive call stack)
- BFS: visit all neighbors by expanding level-by-level (uses queue)

When to use which:
- DFS is typically easier to implement recursively (less boilerplate)
- BFS is helpful when you care about shortest paths or levels (not needed here)
- In this problem, both are equally valid and efficient

🔷 TIME COMPLEXITY:
- Building adjacency list: O(n²)
- DFS traverses every city and connection once: O(n + e)
- Overall: O(n²) since we may check every cell in the matrix

🔷 SPACE COMPLEXITY:
- O(n + e) for adjacency list
- O(n) for visited array
*/

func main() {
	// we have 3 nodes: 0, 1 and 2
	// each node connected to itself (diagonal is always 1) - we can ignore it
	isConnected := [][]int{
		{1, 1, 0}, // 0 -> 1
		{1, 1, 0}, // 1 -> 0
		{0, 0, 1},
	}

	// 0 and 1 are connected, while 2 is alone - we have 2 provinces
	fmt.Println(getNumOfProvincesDfs(isConnected)) // Output: 2

	fmt.Println(getNumOfProvincesBfs(isConnected)) // Output: 2
}

func getNumOfProvincesDfs(isConnected [][]int) int {
	n := len(isConnected)

	// Build the adjacency list
	adjList := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				// add both directions, since we are only lopping the upper triangle
				adjList[i] = append(adjList[i], j)
				adjList[j] = append(adjList[j], i)
			}
		}
	}

	// Track visited cities
	visited := make([]bool, n)

	// Recursive DFS function
	var dfs func(node int)
	dfs = func(node int) {
		if visited[node] {
			return
		}

		visited[node] = true

		for _, neighbor := range adjList[node] {
			dfs(neighbor)
		}
	}

	// Traverse all components using DFS and count provinces
	// DFS must be started for each disconnected component, imagine a worst case scenario - all nodes disconnected - we don't want to miss any of them
	provinces := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)      // Start new DFS from unvisited node i
			provinces++ // increment after a full DFS traversal of one component is complete
		}
	}

	return provinces
}

func getNumOfProvincesBfs(isConnected [][]int) int {
	n := len(isConnected)

	// Step 1: Build the adjacency list
	adjList := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				adjList[i] = append(adjList[i], j)
				adjList[j] = append(adjList[j], i)
			}
		}
	}

	// track visited cities
	visited := make([]bool, n)
	provinces := 0 // counter

	var bfs func(i int)
	bfs = func(i int) {
		queue := []int{i}
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			if visited[node] {
				continue
			}

			visited[node] = true

			// explore all immediate unvisited neighbours
			for _, neighbor := range adjList[node] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	// Traverse all components using BFS and count provinces
	// BFS must be started for each disconnected component, imagine a worst case scenario - all nodes disconnected - we don't want to miss any of them
	for i := 0; i < n; i++ {
		if !visited[i] {
			bfs(i)      // Start new BFS from unvisited node i
			provinces++ // increment after a full BFS traversal of one component is complete
		}
	}

	return provinces
}
