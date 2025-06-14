/*
GENERIC MAP-BASED UNION-FIND – WHY IT’S BETTER

🧠 This implementation uses a map-based Union-Find with generics (instead of array based solution which only supports
ints - union_find.go). Supports any comparable type (e.g. strings, ints, etc).

✅ Key advantages:
- Works directly with strings and other comparable keys
- No need to preallocate parent and rank slices
- Cleaner memory usage — only adds elements on demand
- Supports dynamic, non-contiguous inputs (real-world data)
- Find/Union initialize nodes on the fly (Find calls Add if element is not in the UF structure)
- We use struct and methods instead of passing parent and rank slices around
- Much more natural and readable in problems like Accounts Merge (LeetCode 721)
	- we avoid handling email→ID and ID->email mapping

This approach is more flexible, safer, and more readable — especially in interview scenarios and production code.

⏱ TIME COMPLEXITY NOTE:

Compared to array-based Union-Find, this generic map-based version removes the initial O(n) setup cost for `parent[]` and `rank[]`.

Why?
- Nodes are added on-demand in O(1) using the Add method, which simply inserts entries into the parent and rank maps.
- No need to initialize all possible elements upfront.
- Especially useful when the universe of elements is unknown or sparse.

Let m = number of Union-Find operations (Union/Find calls), and n = max possible elements:

- Array-based version:    O(n + m * α(n))      // upfront + operational cost
- Generic map version:    O(m * α(n))          // avoids initialization

✅ This yields better performance when input is sparse or dynamically discovered.

*/

package union_find

type UnionFind[T comparable] struct {
	parent map[T]T
	rank   map[T]int
}

func NewUnionFind[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
	}
}

func (uf *UnionFind[T]) Find(x T) T {
	if _, ok := uf.parent[x]; !ok {
		uf.Add(x)
	}
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind[T]) Union(x, y T) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootX] = rootY
		uf.rank[rootY]++
	}

	return true
}

func (uf *UnionFind[T]) Add(x T) {
	if _, ok := uf.parent[x]; !ok {
		uf.parent[x] = x
		uf.rank[x] = 0
	}
}
