package main

import (
	"errors"
	"fmt"
	"slices"
)

/*
====================================================================================
🧠 Topological Sort using DFS (Depth-First Search)
====================================================================================

🔷 PROBLEM STATEMENT:
Given a set of tasks and their dependencies (as a DAG), find a valid execution order such that every task is performed
only after all its prerequisites are completed. Same as in previous problem (1_kahn_algo.go) - everything about the problem
is the same - we are just using a different approach to solve it.

🔷 REAL-WORLD EXAMPLE: Software Project Task Scheduling

Tasks:
0 - Setup project repository
1 - Install dependencies
2 - Design database schema
3 - Implement backend APIs
4 - Setup CI/CD pipeline
5 - Write unit tests
6 - Frontend integration
7 - Deploy to staging environment

Dependencies (A must be done before B):
[0, 1] // Setup repo → Install dependencies
[1, 2] // Install dependencies → Design DB schema
[2, 3] // DB schema → Implement APIs
[1, 4] // Install dependencies → Setup CI/CD
[3, 5] // APIs → Unit tests
[4, 7] // CI/CD → Deploy
[5, 7] // Unit tests → Deploy
[6, 7] // Frontend integration → Deploy
[1, 6] // Install dependencies → Frontend integration

🔷 TYPE OF GRAPH — DAG (Directed Acyclic Graph):
- Directed: all edges go one way
- Acyclic: no cycles allowed
- Can be disconnected but topo sort must process all nodes

🔷 GRAPH REPRESENTATION:
- Adjacency List: `adj[i]` holds all tasks that depend on task i (neighbours)
- Visited Array: `visited[i]` holds one of 3 states:
  - 0 = unvisited
  - 1 = visiting (currently on the same DFS path)
  - 2 = visited (fully processed)

🔷 CORE IDEA — Post-Order DFS:
- Start from any node (all nodes are unvisited in the beginning)
	- Actually we call the recursive dfs function on all (unvisited) nodes to make sure that we cover all nodes (in case of disconnected graph)
- Traverse to all neighbors as far down as possible
- After vising all neighbors/dependencies mark the current node as complete and add it to result
	- node is added to result after visiting all its dependencies
    - so node is added in "post-order" (in reverse order)
- Reverse the result at the end to obtain correct topological order

Note: we use a stack (implicit via recursion).

🔷 CYCLE DETECTION:
- If we revisit a node that is already in the current recursion path, it means we’ve encountered a cycle
	- current recursion path contains nodes which are neighbours (direct/indirect) to each other
	- if we encountered the same node twice - it's related to itself
- So if we enter a node where `visited[i] == 1` -> node is already being explored on the current DFS path → cycle exists

🔷 ALGORITHM STEPS:
1. Build the adjacency list from the edges
2. Initialize `visited` array and an empty `result` slice
3. For each node 0 to n-1:
   - If `visited[node] == 0`, run DFS on it
4. In DFS:
   - If `visited[node] == 2`, return true (already processed)
   - If `visited[node] == 1`, return false (cycle detected — node already active in recursive call stack)
   - Mark node as visiting
   - Recurse on all neighbors
   - Mark node as visited
   - Append to result (post-order)
5. Reverse result for correct topological order

🔷 COMPARISON — DFS vs KAHN’S ALGORITHM:

Kahn/BFS uses “expand and explore” pattern:
- start with what's reachable now
- gradually move outward (level by level) to reachable neighbors once dependencies are cleared
- uses a queue

DFS uses "deepest path first" pattern:
- explores as far down/deep as possible along one path before backtracking
- uses a stack (implicit via recursion) instead of a queue
- adds nodes to the result in post-order and then reverses the result

| Feature               | DFS-Based Topo Sort           | Kahn’s Algorithm (BFS)         |
|-----------------------|-------------------------------|--------------------------------|
| Traversal Method      | Recursive depth-first         | Iterative breadth-first        |
| Cycle Detection       | Revisiting node in dfs path   | Incomplete result length       |
| Result Construction   | Post-order + reverse          | Append as processed            |
| In-degree Tracking    | ❌ Not needed                 | ✅ Required                    |
| Queue/Stack Used      | Stack via recursion           | Queue of in-degree zero nodes  |
| Use Case Strengths    | Elegant and concise recursion | Task execution readiness model |

🔷 TIME & SPACE COMPLEXITY:

Time complexity: O(n + e)
- Build adjacency list: O(e)
- DFS traversal: O(n) for nodes + O(e) for edges
	- We visit all nodes and edges once
		- we explore all nodes and their neighbours
    	- we don't revisit already explored nodes (visited check)
- Total: O(n + e)

Space complexity: O(n + e)
- Adjacency list: O(e)
- Visited array: O(n)
- Recursion stack: up to O(n) in depth (in worst case)
	- If all nodes are part of one big neighbour stack
	- Ex. we have nodes [A,B,C,D,E] and edges are A->B->C->D->E
	- Then the stack depth will be 5 (same as number of nodes)
- Result slice: O(n)
- Total: O(n + e)

🔷 WHEN TO USE KAHN VS DFS:
- Both have same time/space complexity: O(n + e) and both solve the same exact problem.

Prefer Kahn:
- when tasks must be executed as soon as they’re ready (as soon as all dependencies are ready)
	- task scheduling problems and similar
- if you are already computing in-degrees for another purpose

Prefer DFS:
- when you prefer recursive flow
- when you want more localized and stronger cycle detection -> easy to reconstruct the cycle path
	- ex: track some parent/ancestor map and reconstruct the cycle
*/

func topoSortDFS(n int, edges [][]int) ([]int, error) {
	// Step 1: Build the adjacency list
	adj := make([][]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		adj[from] = append(adj[from], to)
	}

	// 0 = unvisited, 1 = visiting in current dfs path, 2 = visited/processed
	visited := make([]int, n)
	var result []int

	// Step 2: Define recursive DFS function
	var dfs func(node int) bool
	dfs = func(node int) bool {
		if visited[node] == 1 {
			// Node is already on the current DFS path → cycle
			return false
		}
		if visited[node] == 2 {
			// Node already processed in a previous DFS branch
			return true
		}

		visited[node] = 1 // mark this node as visiting

		// Recurse on all neighbors (dependencies)
		for _, neighbor := range adj[node] {
			if !dfs(neighbor) {
				return false // exit as soon as cycle is detected
			}
		}

		visited[node] = 2             // done processing this node
		result = append(result, node) // post-order add (in reverse order)
		return true
	}

	// Step 3: DFS from all unvisited nodes
	// we do this to ensure that we cover all nodes, for example if we have a disconnected graph
	for node := 0; node < n; node++ {
		if visited[node] == 0 {
			if !dfs(node) {
				return nil, errors.New("cycle detected")
			}
		}
	}

	// Step 4: Reverse post-order result to get topological order
	slices.Reverse(result)
	return result, nil
}

func main() {
	// Test case 1: Valid DAG
	order1, err1 := topoSortDFS(4, [][]int{
		{1, 0}, {2, 0}, {3, 1}, {3, 2},
	})
	if err1 != nil {
		fmt.Println("❌", err1)
	} else {
		fmt.Println("✅ Topological Order (1):", order1) // [3 2 1 0]
	}

	// Test case 2: Direct self-loop (cycle)
	_, err2 := topoSortDFS(4, [][]int{
		{1, 0}, {2, 0}, {3, 1}, {3, 2}, {3, 3},
	})
	fmt.Println("Expected cycle (2):", err2)

	// Test case 3: Indirect cycle
	_, err3 := topoSortDFS(4, [][]int{
		{1, 0}, {2, 0}, {3, 1}, {3, 2}, {1, 3},
	})
	fmt.Println("Expected cycle (3):", err3)

	// Test case 4: Indirect cycle involving 0
	_, err4 := topoSortDFS(4, [][]int{
		{1, 0}, {2, 0}, {3, 1}, {3, 2}, {0, 3},
	})
	fmt.Println("Expected cycle (4):", err4)

	// Test case 5: Disconnected graph, valid sort
	order5, err5 := topoSortDFS(6, [][]int{
		{0, 1}, {1, 2}, {3, 4},
	})
	if err5 != nil {
		fmt.Println("❌", err5)
	} else {
		fmt.Println("✅ Topological Order (5):", order5) // [5 3 4 0 1 2]
	}
}
