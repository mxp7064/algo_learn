# DSA Pattern Decision Guide

## ◄ Arrays / Strings

- **Is the array sorted or partially sorted?**  
  → Try: Two Pointers, Binary Search, Prefix Sums, Sliding Window

- **Is it an optimization (max/min/longest/smallest subarray)?**  
  → Try: Sliding Window, Greedy, or DP

- **Looking for duplicates, counts, or frequencies?**  
  → Try: HashMap, HashSet, or Counting Array

- **Need all substrings, fixed-size subarrays, or dynamic ranges?**  
  → Try: Sliding Window (fixed or dynamic), Two Pointers

- **Frequent min/max in window?**  
  → Try: Deque (Monotonic Queue), Heap

- **Match/remove characters, balance parentheses, decode strings?**  
  → Try: Stack

- **Generate all combinations, subsets, permutations?**  
  → Try: Backtracking (Recursive DFS)

---

## ◄ Graph / Matrix Problems

- **Find shortest path in unweighted graph or grid?**  
  → Use: Breadth-First Search (BFS)

- **Find shortest path in weighted graph?**  
  → Use: Dijkstra (non-negative weights), Bellman-Ford (negative weights)

- **Count connected components or check connectivity?**  
  → Use: DFS, BFS, or Union-Find (Disjoint Set)

- **Detect or avoid cycles?**  
  → Use: DFS + Recursion Stack, or Union-Find (for undirected)

- **Find topological order of tasks (with dependencies)?**  
  → Use: Kahn’s Algorithm (BFS) or DFS + post-order

- **Explore or search a grid? (e.g. Islands, Water Flow)**  
  → Use: DFS, BFS, or Backtrack DFS

---

## ◄ Trees & BSTs

- **Traverse a tree?**  
  → Use: Inorder, Preorder, Postorder, Level Order (BFS)

- **Balanced tree check or diameter calculation?**  
  → Use: Postorder + height return

- **Find Lowest Common Ancestor?**  
  → Use: Recursive DFS, or Parent Map + Ancestor Set

- **Need recursive logic on left/right subtree with pruning?**  
  → Use: BST + recursion (e.g. range sum, kth smallest)

---

## ◄ Linked Lists

- **Cycle detection?**  
  → Use: Fast and Slow Pointers (Floyd’s Cycle Detection)

- **Find middle or nth node, or merge sorted lists?**  
  → Use: Two Pointers, Dummy Nodes

- **Reverse list or segments?**  
  → Use: Prev / Curr / Next pointers

---

## ◄ Dynamic Programming (DP)

- **Optimal substructure + overlapping subproblems?**  
  → Use: Top-Down (Memoization) or Bottom-Up (Tabulation)

- **Subset, knapsack, coin change?**  
  → Use: 1D/2D DP Arrays

- **String comparisons/edit distance/LCS?**  
  → Use: DP Matrix

- **DP state has multiple variables?**  
  → Use: dp[i][j][k] or map[[...]]int depending on constraints

---

## ◄ Range Queries / Prefix Logic

- **Many sum/range queries, no updates?**  
  → Use: Prefix Sums, Difference Arrays

- **Many sum/range queries + updates?**  
  → Use: Segment Tree, Fenwick Tree (Binary Indexed Tree)

---

## ◄ Bit Manipulation

- **Check parity, toggle, mask bits, XOR trick?**  
  → Use: &, |, ^, >>, <<

- **Enumerate subsets of a set?**  
  → Use: Bitmasking

- **Fast frequency toggling (e.g. Palindromic Anagram Check)?**  
  → Use: Bitmask for character parity

---

## ◄ Heaps, Top K, and Priority

- **Find top K / bottom K / frequent K elements?**  
  → Use: Heap (min/max depending on goal)

- **Need median in stream or dynamic set?**  
  → Use: Two Heaps (Max/Min)

- **Need exact K-th smallest/largest element?**  
  → Use: QuickSelect (O(n) average)

---

## ◄ Tricky Techniques / Others

- **Subarray max/min with push/pop logic?**  
  → Use: Monotonic Stack

- **Need optimal subset of intervals or items?**  
  → Use: Greedy (if local optimal → global optimal), else DP

- **Prefix-based string matching?**  
  → Use: Trie

- **Avoid recursion or hitting stack limit?**  
  → Convert recursion to iterative with explicit Stack