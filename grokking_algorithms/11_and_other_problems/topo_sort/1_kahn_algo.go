package main

import (
	"errors"
	"fmt"
)

/*
============================================================================================
🧠 Topological Sort using Kahn's Algorithm - Breadth-First Search (BFS) based
============================================================================================

// To recap the story check chapter 6 and BFS, topo sort is just briefly mentioned.
// Topo sort can be done in 2 ways - Kahn algorithm (BFS based) and DFS based sort

🔷 PROBLEM STATEMENT:
Given a set of tasks (nodes) and a list of dependencies (edges),
find a valid order to perform the tasks such that for every dependency, represented as edge [A → B],
task A is done before task B.

This is a classic Topological Sort problem.

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

🔷 TYPE OF GRAPH:
- Directed
- Acyclic (must have no cycles)
- Can be disconnected

🔷 TERMINOLOGY:
- totalTasks (V): total number of tasks (nodes in the graph)
- edges: list of [from, to] pairs representing directed dependencies
- indegree[i]: number of prerequisites/dependencies for task i
- adjacency/neighbours list: for each source task/node, stores which tasks depend on it (neighbour nodes it points to)

🔷 CORE IDEA (Kahn's Algorithm):
1. Create an adjacency list and compute the indegree of each node.
2. Add all nodes with indegree 0 to a queue (tasks ready to start).
3. While the queue is not empty:
   a. Remove a node from the queue and add it to the result.
   b. Decrease the indegree of its neighbors by 1.
   c. If any neighbor’s indegree becomes 0 (all prerequisites for this task are done), add it to the queue (task is ready to "execute" now)
4. If the result includes all tasks, return it. Otherwise, a cycle exists - some node(s) didn't reach indegree = 0.

🧠 Why Kahn is considered "BFS-style":
- Nodes are explored by visiting their immediate neighbors, similar to BFS
- A queue is used to track nodes that are ready to be processed (in-degree = 0)
- This mirrors the level-by-level exploration found in classic BFS

But unlike classic BFS:
- A node is not immediately added to the result when first discovered.
- It is only processed once all its dependencies are resolved (i.e., when in-degree becomes 0)
- A node may be "discovered" via one incoming edge/path but added to the queue and processed later - when in-degree
becomes 0, i.e. when we reach it later via some other incoming edge/path

Key Insight:
- BFS in Kahn's Algorithm refers to how we explore (via immediate neighbors), not when we process or include a node in the result.
- Processing is delayed until the node has no remaining unresolved dependencies.

This makes Kahn’s Algorithm ideal for task scheduling problems, dependency resolution (for example in software build tools), cycle detection and so on.

LeetCode issues:
- LC 207 — Can finish courses? (Cycle detection)
- LC 210 — Return valid course order
- LC 269 — Alien dictionary (topo on characters)
- LC 1203 — Topo sort with groups and subgroups

🔷 TIME COMPLEXITY:

Let:
- n = number of nodes
- e = number of edges

Total Time Complexity: O(n + e)

Breakdown:

1. Building the graph:
   - For each edge, update in-degree and adjacency list → O(e)

2. Initializing the queue:
   - Scan all nodes to find those with in-degree = 0 → O(n)

3. Processing the queue:
   - Each node is enqueued/dequeued at most once → O(n)
   - Each edge is traversed exactly once (when its source node is processed) → O(e)

So overall:
   - O(e) + O(n) + O(n) + O(e) -> O(n + e)

🔷 SPACE COMPLEXITY:

Total Space Complexity: O(n + e)

Breakdown:

1. In-degree map → O(n)
   - Tracks number of incoming edges for each node

2. Adjacency list → O(e)
   - Stores all outgoing edges for each node (total edges = e)

3. Queue → O(n)
   - Holds nodes with in-degree = 0 (in worst case, all nodes)

4. Topo order result → O(n)
   - Stores the final topological ordering

So total space used is:
   - O(n) + O(e) + O(n) + O(n) → O(n + e)

🧠 WHY IS IT CALLED "TOPOLOGICAL SORT"?

- “Topology” in graph theory refers to the structure of how nodes are connected — not the geometric layout, but the direction
and dependency relationships between nodes
- "Topological" refers to the structure of a directed graph — specifically the dependency relationships between nodes
- A topological sort produces a linear ordering of nodes such that for every directed edge u → v, node u comes before node v in the ordering
- This is not a traditional "sort" of values, but a sorting based on graph structure
- It only applies to DAGs (Directed Acyclic Graphs), where no circular dependencies exist
- We solve topological sort problem either with Kahn’s Algorithm (BFS based topo sort) or DFS based topo sort

*/

// totalTasks is the number of vertices (nodes) in the graph — in our real-world example,
// it represents the total number of tasks
// We need to initialize the adjacency list and indegree array using the size of the graph - It ensures we loop through
// all possible nodes, including those that may not appear as sources in the edges
// Even if a task has no outgoing edges, we still need to include it in the graph, because:
// - It might have incoming edges (dependencies),
// - It might be standalone (yet still valid to execute),
// - Topo sort must return all nodes in a valid order.
// topologicalSort returns false if topo sort is not possible, i.e. when cycle is detected
func topologicalSort(totalTasks int, edges [][]int) ([]int, error) {
	// Step 1: Create adjacency list
	adj := make([][]int, totalTasks) // adj[i] = list of tasks that depend on task i (neighbours)
	// tasks are chosen/marked 0...7 on purpose so task can also be an index, after we process the edges, we will have:
	// [
	// 	 0 → [1]
	// 	 1 → [2 4 6]
	// 	 2 → [3]
	// 	 3 → [5]
	// 	 4 → [7]
	// 	 5 → [7]
	// 	 6 → [7]
	// 	 7 → []
	// ]

	indegree := make([]int, totalTasks) // indegree[i] = number of dependencies for task i
	// at first, indegree slice will be all zeros, and after we process the edges it will be:
	// 0 -> 0
	// 1 -> 1
	// 2 -> 1
	// 3 -> 1
	// 4 -> 1
	// 5 -> 1
	// 6 -> 1
	// 7 -> 3

	// process edges - build adjacency list and count indegree
	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		adj[from] = append(adj[from], to)
		indegree[to]++ // task 'to' has one more dependency
	}

	// Step 2: Initialize queue with tasks that have no dependencies (indegree == 0)
	queue := []int{}
	for i := 0; i < totalTasks; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i) // in our example -> task 0 is "ready to start" because it has no dependencies
		}
	}
	// queue = [0]

	// Step 3: Process the queue
	order := []int{}
	for len(queue) > 0 {
		// Remove task from the front of the queue (pop from queue)
		current := queue[0] // get the first task
		queue = queue[1:]   // remove first task from queue

		// Add it to the final execution order, aka "task executed" or "task can be executed" or "add task to execution order"
		order = append(order, current) // task can be executed when it has no remaining prerequisites/dependencies (indegree == 0)

		// Task current is now done!” - we added it to the execution order, now let’s unlock any tasks that were waiting for it
		// For each task that depends on current, reduce its indegree
		for _, neighbor := range adj[current] { // neighbours are tasks which depend on current task
			indegree[neighbor]--         // 1 dependency is resolved (current task)
			if indegree[neighbor] == 0 { // if all dependencies are resolved, task is now ready to start
				queue = append(queue, neighbor) // add it to queue
			}
		}
	}

	// Step 4: Check for cycles
	// If there’s a cycle, some nodes will never reach indegree = 0
	// - they stay stuck waiting for each other forever
	// - they always have at least one incoming edge from inside the cycle (they will always have +1 in indegree)
	// - so they are never added to queue and order arrays so len(order) < totalTasks
	if len(order) != totalTasks {
		return nil, errors.New("cycle detected — topological sort not possible")
	}

	return order, nil
}

func main() {
	// Total number of tasks
	totalTasks := 8

	// Dependencies: each pair [A, B] means A must be done before B
	edges := [][]int{
		{0, 1}, // Setup repo → Install dependencies
		{1, 2}, // Install dependencies → Design DB schema
		{2, 3}, // DB schema → Implement APIs
		{1, 4}, // Install dependencies → Setup CI/CD
		{3, 5}, // APIs → Unit tests
		{4, 7}, // CI/CD → Deploy
		{5, 7}, // Unit tests → Deploy
		{6, 7}, // Frontend integration → Deploy
		{1, 6}, // Install dependencies → Frontend integration
	}

	order, err := topologicalSort(totalTasks, edges)
	if err == nil {
		fmt.Println(order)
	} else {
		fmt.Println(err)
	}
	// Task 0 → Task 1 → Task 2 → Task 4 → Task 6 → Task 3 → Task 5 → Task 7 → DONE

	// Example with cycle
	edges2 := [][]int{
		{0, 1}, // Setup repo → Install dependencies
		{1, 2}, // Install dependencies → Design DB schema
		{2, 3}, // DB schema → Implement APIs
		{1, 4}, // Install dependencies → Setup CI/CD
		{3, 5}, // APIs → Unit tests
		{4, 7}, // CI/CD → Deploy
		{5, 7}, // Unit tests → Deploy
		{6, 7}, // Frontend integration → Deploy
		{1, 6}, // Install dependencies → Frontend integration
		{7, 5}, // we have a cycle here
	}

	order, err = topologicalSort(totalTasks, edges2)
	if err == nil {
		fmt.Println(order)
	} else {
		fmt.Println(err)
	}

	// Simplest example with cycle - on start all nodes have indegree = 1 - we can't even start the queue
	edges3 := [][]int{
		{0, 1},
		{1, 0},
	}

	order, err = topologicalSort(2, edges3)
	if err == nil {
		fmt.Println(order)
	} else {
		fmt.Println(err)
	}
}
