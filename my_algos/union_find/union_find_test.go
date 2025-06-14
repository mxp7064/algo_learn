package union_find

import (
	"fmt"
	"testing"
)

func Test_Naive(t *testing.T) {
	n := 6                   // Number of elements (0 through 5) so our elements are: 0, 1, 2, 3, 4, 5
	parent := make([]int, n) // parent[i] = parent of node i

	// Initially, each node is its own parent → each node is a separate set
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	// Merge some elements into common sets
	unionNaive(0, 1, parent)
	unionNaive(1, 2, parent)
	unionNaive(3, 4, parent)

	// At this point:
	// - 0, 1, 2 are connected
	// - 3 and 4 are connected
	// - 5 is still alone

	fmt.Println("Check which root each element belongs to")
	fmt.Println("findNaive(0):", findNaive(0, parent)) // should be same as 1 and 2
	fmt.Println("findNaive(1):", findNaive(1, parent))
	fmt.Println("findNaive(2):", findNaive(2, parent))

	fmt.Println("findNaive(3):", findNaive(3, parent)) // should be same as 4
	fmt.Println("findNaive(4):", findNaive(4, parent))

	// Check if two elements are in the same set
	fmt.Println("findNaive(2) == findNaive(3):", findNaive(2, parent) == findNaive(3, parent)) // false

	// Merge the two sets: now 0–1–2–3–4 should all be connected
	unionNaive(2, 3, parent)

	fmt.Println("After union(2, 3):")
	for i := 0; i <= 4; i++ {
		fmt.Printf("findNaive(%d) = %d\n", i, findNaive(i, parent))
	}
}

func Test_Optimized(t *testing.T) {
	n := 6
	parent := make([]int, n)
	rank := make([]int, n) // rank[i] = approx height of the tree rooted at i

	// Initial setup: each node is its own parent, rank = 0
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	// Perform unions just like before
	Union(0, 1, parent, rank) // 0, 1
	Union(1, 2, parent, rank) // 0, 1, 2
	Union(3, 4, parent, rank) // 3, 4

	// Show representatives (roots) — these should now resolve faster
	fmt.Println("Find(0):", Find(0, parent))
	fmt.Println("Find(1):", Find(1, parent))
	fmt.Println("Find(2):", Find(2, parent))
	fmt.Println("Find(3):", Find(3, parent))
	fmt.Println("Find(4):", Find(4, parent))
	fmt.Println("Find(2) == Find(3):", Find(2, parent) == Find(3, parent)) // false

	// Merge the two groups
	Union(2, 3, parent, rank)

	fmt.Println("After Union(2, 3):")
	for i := 0; i <= 4; i++ {
		fmt.Printf("Find(%d) = %d\n", i, Find(i, parent))
	}
}
