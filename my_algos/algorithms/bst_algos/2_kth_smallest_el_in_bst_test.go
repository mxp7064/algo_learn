/*
🧠 Kth Smallest Element in a Binary Search Tree (BST)

PROBLEM:
Given a BST and an integer k, return the kth smallest element in the tree (1-based index).
This is a very common interview problem and is based on the fact that in-order traversal
of a BST gives sorted elements.

KEY CONCEPT:
- In-order traversal (Left → Node → Right) of a BST yields elements in sorted order.
- So the kth visited node during in-order traversal is the kth smallest element.

TWO APPROACHES:

1) Naive approach:
- Do a full in-order traversal to get all elements
- Return the (k-1)th element from the sorted slice
- Time: O(n), Space: O(n)

2) Optimized Early-Exit Approach:
- Traverse in-order recursively
- Keep a counter `k`, decrement it as we visit nodes
- When counter reaches 0, return the current node's value immediately
- No need to traverse the rest of the tree
- Time: O(k) best case, O(n) worst case
- Space: O(h) for recursion stack (h = height of tree)

NOTE:
The optimized version has better memory efficiency and can exit early without visiting the whole tree.
*/

package bst_algos

import (
	"fmt"
	"panca.com/algo/bst"
	"testing"
)

func Test_kthSmallestElementInBst(t *testing.T) {
	// input
	data := []int{1, 4, 2, 5, 7}
	k := 3 // we want to find third smallest element

	// build the BST
	bst := bst.BuildBSTFromSlice(data)

	// ✅ Naive approach: get sorted array from in-order traversal
	sorted := bst.InOrderTraversal() // [1 2 4 5 7]
	fmt.Println(sorted[k-1])         // Output: 4

	// ✅ Better approach: early exit and more efficient memory footprint
	value, found := getKthSmallestNodeValueEarlyExit(bst.Root, &k)
	if found {
		fmt.Println(value) // Output: 4
	} else {
		fmt.Println("k is larger than the number of nodes in the BST")
	}
}

// getKthSmallestNodeValueEarlyExit performs an in-order traversal with early exit.
// It decrements *k as it visits each node. When *k reaches 0, we found the kth element.
// Returns (value, true) if found, or (-1, false) if not found.
func getKthSmallestNodeValueEarlyExit(node *bst.Node[int], k *int) (int, bool) {
	if node == nil {
		// k smallest not found (for example if k is bigger than tree length)
		return -1, false // base case: empty subtree
	}

	// 🡒 Recurse left
	if value, found := getKthSmallestNodeValueEarlyExit(node.Left, k); found {
		return value, true // early exit if node already found
	}

	// 🡒 Visit current node
	*k--         // decrement counter
	if *k == 0 { // when it reaches 0 - that's our element
		return node.Value, true // // early exit if node is found
	}

	// 🡒 Recurse right
	return getKthSmallestNodeValueEarlyExit(node.Right, k)
}
