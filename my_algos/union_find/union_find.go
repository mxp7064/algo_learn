package union_find

/*
===========================================================
🧠 Union-Find (Disjoint Set Union) — Naive Implementation
===========================================================

🔷 PURPOSE:
This data structure efficiently tracks which elements are in the same connected component.

💡 Real-World Analogy: Islands and Bridges
Imagine you have 10 islands, and people start building bridges between them.
Each time a bridge is built, two islands become part of the same landmass (connected component).
You want to answer questions like:
- Are island 3 and island 7 connected?
- How many landmasses are there now?

This is where Union-Find shines.

💡 Why is it called Disjoint set union?
- Disjoint = no overlaps (no element appears in more than one set)
- Set = groups of elements
- Union = operation that merges sets/groups
	- main operation is to merge (union) two of these disjoint sets into one.
	- ex: if a is in set A and c is in set B, then Union(a, c) will merge sets A and B.
- But more commonly we just call it Union-Find
	- Find is operation which returns the set/group representative element (effectively set/group ID)

🔷 Useful for:
- Grouping elements into disjoint sets
- Detecting cycles in undirected graphs
- Counting connected components
- Merging users/accounts/networks/etc.

🔧 Basic Idea
We model each set (component) as a tree, where:
- Each node has a parent (we have some kind of 'parent' function/array which maps nodes to their parents)
	- parent[x] = parent of node x
- The root node is the representative/leader of the set
- Initially each element is its own parent (ex. parent[0] = 0, parent[1] = 1, etc.)
Then we can do the following operations:
- find(x): returns the "representative" (root) of the set containing x
	- root is effectively a group/set ID
- union(x, y): merges the sets containing x and y creating a new bigger set
	- now find(x) is same as find(y), i.e. they have the same root now

🔷 Naive vs optimized version
NAIVE:
- ✅ Simple and correct
- ❌ Not optimized (no path compression, no union by rank/size)
- Trees can become deep
- Slower for large data sets

OPTIMIZED:
- ✅ Path Compression (in `find`)
- ✅ Union by Rank (in `union`)
- These optimizations dramatically reduce tree depth and make operations nearly constant time - O(α(n))
	- α(n) is the inverse Ackermann function (no need to understand it in mathematical sense)
  	- Grows extremely slowly, almost constant for all practical n (α(n) ≤ 4 for n < 2^65536)
  	- It’s treated as constant in practice
- This version is optimal and used in problems like:
	- Detecting cycles in undirected graphs
	- Kruskal’s algorithm
	- Connected components
- This version is ideal for large datasets or many union/find operations.

🔷 IMPORTANT NODE:
This solution only supports ints - if we want to use other data types (such as strings), we need to handle mappings between
int IDs and element strings (or whatever other data type). We also need to initialize parent and rank arrays and pass them
around to Union/Find calls.
In union_find_generic.go you can find improved generic version which uses maps instead of arrays for storing parent and rank
structures.

*/

// Naive find: follow parent pointers until reaching the root of the set
// This just walks up the tree until it finds the root (where parent[x] == x) - root is defined as a node which is its own parent
// Each node must have a parent (initially each node is its own parent, i.e. each node is its own set)
// This implementation is intentionally simple and does not include path compression which makes find faster by flattening the tree
func findNaive(x int, parent []int) int {
	for parent[x] != x {
		x = parent[x]
	}
	return x // return node for which parent[x] == x
}

// union joins two sets if they’re currently separate (merges the sets of x and y)
// After a union, both elements share the same root (leader) — i.e., they are in the same connected group/set
// Finds the root of both x and y, if they are different → attaches one to the other
// We are effectively joining two trees — so that x and y now share the same ultimate root (it can be root of x or root of y)
// This is "naive" union - connects one root to another without tree balancing. This is fine if we are lucky and
// we attach a smaller tree under a larger one (tree height will remain the same), but it is problematic when
// we attach a larger tree under a smaller one - the larger tree gains a new parent - increasing its height by 1
// In the optimized version (union by rank) we will ensure that we always attach smaller tree under the larger one
// so we don't increase the height.
// Note that even in the optimized version we can't avoid increasing the height of the tree by 1 when we do a union of
// trees with equal height but at least we won't have an increase when tree height differs
func unionNaive(x, y int, parent []int) {
	rootX := findNaive(x, parent)
	rootY := findNaive(y, parent)

	// if rootX == rootY that means that x and y are already in the same set so we don't have to do anything
	if rootX != rootY {
		// merge sets by attaching one root under the other - we arbitrarily pick which root becomes the parent
		parent[rootX] = rootY // we can also do parent[rootY] = rootX
	}
}

// Find with path compression - optimized version of findNaive
// Can flatten the whole tree if it's called on the deepest element or otherwise it will flatten at least a part of the tree.
// Path compression affects internal structure of the tree to speed up future find() calls.
// After multiple calls to find(), all nodes in the same set will point directly to the root (flattens the tree)
// This means future find(x) calls will return in O(1) time (we have only one level of traversal now - root is the
// immediate parent of each node) for practical purposes.
func Find(x int, parent []int) int {
	if parent[x] != x {
		parent[x] = Find(parent[x], parent) // flatten the path during recursion
	}
	return parent[x] // base case when parent[x] == x
}

// Union by rank - optimized version of unionNaive
// This optimized version helps keep trees shallow.
// Returns false if x and y are already in the same set (i.e. cycle detected).
// Rank is a heuristic estimate/approximation of tree height (not the exact height).
// Rank helps guide union decisions to keep the trees shallow, but:
// - It’s not updated every time a tree grows
// - It’s only incremented when two trees of equal rank are merged
// When merging two sets, always attach the tree with lower rank under the one with higher rank, i.e. attach the
// smaller tree under the larger one to prevent deep trees.
// If both have equal rank, choose one arbitrarily as new root and increment its rank by 1.
// Remember that when merging trees of equal height we must increase the height by 1 in any case but still note that
// rank is a heuristic measure - we are not tracking the actual tree height - instead we are tracking the "number of times"
// the tree height increased for some tree/set so we can decide which tree is "bigger" or "smaller".
// This helps guarantee that the tree height stays low — even before path compression of find kicks in.
// Combined with path compression, guarantees nearly flat trees.
func Union(x, y int, parent, rank []int) bool {
	rootX := Find(x, parent)
	rootY := Find(y, parent)

	if rootX == rootY {
		// A cycle is found: x and y are already connected
		return false
	}

	// attach smaller rank tree under larger rank tree
	if rank[rootX] < rank[rootY] {
		// smaller tree (rootX) becomes a child of the larger tree (rootY)
		// we don't update the rank because the larger tree's height didn’t increase
		parent[rootX] = rootY
	} else if rank[rootX] > rank[rootY] {
		// same logic as above
		parent[rootY] = rootX
	} else {
		// same rank → choose one as root arbitrarily (rootX or rootY)
		parent[rootY] = rootX
		rank[rootX]++ // The "receiving" tree is now larger/deeper so we increment the rank by 1
	}

	return true
}
