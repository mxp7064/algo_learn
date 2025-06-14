# 📊 Core Graph Properties (for DSA & Algorithm Selection)

Understanding graph properties is key to choosing the right algorithm. These are the most relevant traits to track in **interview-level DSA problems**, along with their implications:

---

## 1. Directed vs Undirected

- **Directed**: Edges have direction (u → v).  
  - Traversal must respect direction.
  - Use cases: course scheduling, build systems, web links.
  - Algorithms: Topological Sort, Kahn’s Algorithm, DFS (with direction), Dijkstra.

- **Undirected**: Edges go both ways (u ↔ v).
  - Treat neighbors as mutual.
  - Use cases: social networks, road networks.
  - Algorithms: Union-Find, DFS/BFS for connected components, MST (Kruskal, Prim).

---

## 2. Connected vs Disconnected

- **Connected**: All nodes reachable from any node.
  - No need to check for isolated components.
  - BFS/DFS from one node visits all others.

- **Disconnected**: Multiple separate components.
  - Must run BFS/DFS/Union-Find across all nodes.
  - Algorithms: Count connected components, Number of Provinces (DFS/Union-Find).

---

## 3. Presence of Cycles

- **Cycles Present**:
  - Some algorithms break or must detect cycles (e.g. Topo Sort, DFS recursion path).
  - Must be avoided in scheduling, dependency resolution.

- **Acyclic (DAG)**:
  - Enables topological ordering.
  - Safe for backtracking, DP on DAGs.

- Detection:
  - **DFS**: via recursion stack
  - **Union-Find**: for undirected graphs
  - **Topo Sort**: if result size < node count → cycle

---

## 4. Weighted vs Unweighted

- **Weighted**: Edges have cost.
  - Use Dijkstra, Bellman-Ford, or Prim/Kruskal.
  - Sorting, priority queues (heap) involved.

- **Unweighted**:
  - Use plain BFS for shortest path (min # of edges).
  - Use DFS for reachability, connected components.

---

## 5. Negative Edge Weights?

- **No**: Safe to use Dijkstra (greedy with min-heap).
- **Yes**:
  - Dijkstra breaks.
  - Use Bellman-Ford (O(VE)) — handles negatives, detects negative cycles.

---

## Traversal Implications Summary

| Property              | Safe with BFS | DFS | Dijkstra | Union-Find | Topo Sort |
|-----------------------|---------------|-----|----------|-------------|------------|
| Directed              | ✅            | ✅  | ✅       | ❌          | ✅         |
| Undirected            | ✅            | ✅  | ✅       | ✅          | ❌         |
| Disconnected          | Loop required | ✅  | ✅       | ✅          | ✅         |
| Cycles Present        | ✅ (loop-safe) | ✅* | ❌*      | ✅ (detect) | ❌         |
| Negative Edge Weights | ✅            | ✅  | ❌       | ✅          | ✅         |

*DFS with recursion stack detects cycles in directed graphs.  
*Dijkstra fails with negative edges — use Bellman-Ford.

---

## Tips

- Always check edge weight sign and graph direction.
- Topological Sort requires DAG (Directed Acyclic Graph).
- BFS is best for unweighted shortest paths.
- Union-Find only works on undirected graphs for cycle/component detection.
