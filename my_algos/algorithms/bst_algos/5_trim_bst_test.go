/*
🧠 TRIM A BINARY SEARCH TREE

PROBLEM:
Given the root of a BST and a range [min, max], remove all nodes with values
outside this range and return the new root.

Only valid nodes within [min, max] should remain in the tree.
The resulting tree must still be a valid BST.

KEY IDEA:
- Use recursion to traverse the tree.
- For each node:
  - If node.Value < min → skip the node (and its left subtree) and recurse right
  - If node.Value > max → skip the node (and its right subtree) and recurse left
  - If node.Value is within [min, max] → keep the node and recursively trim its children

WHY THIS WORKS:
- We leverage the BST property:
    Left subtree < node < Right subtree
- When pruning, we bubble up valid subtrees from left/right children as needed.
- It’s OK to assign a node's child to a subtree that was originally on the opposite side of a trimmed node
  — this does not violate the BST invariant.

TIME & SPACE:
- Time: O(n) — must visit each node once
- Space: O(h) — recursion stack, where h is height of tree
*/

package bst_algos

import (
	"fmt"
	"panca.com/algo/bst"
	"testing"
)

func Test_TrimBst(t *testing.T) {
	data := []int{8, 3, 10, 1, 6, 14, 13}
	tree := bst.BuildBSTFromSlice(data)
	trimmedRootNode := trimBst(tree.Root, 5, 13)
	/*
			Original BST:
			         8
			       /   \
			      3    10
			     / \     \
			    1   6    14
			              /
			             13

			Trim range: [5, 13]

			Trimmed BST should be:
			        8
			       / \
			      6   10
			           \
		               13
	*/
	// trimBst changes the original tree in place but we want to demonstrate that trimBst returns a new root
	trimmed := bst.NewBSTFromRootNode(trimmedRootNode)
	fmt.Println(trimmed.InOrderTraversal()) // [6 8 10 13]
}

// trimBst recursively removes nodes outside the [min, max] range and returns the new subtree root
func trimBst[T bst.Ordered](node *bst.Node[T], min, max T) *bst.Node[T] {
	if node == nil {
		return nil
	}

	if node.Value < min {
		// Node too small, trim/skip it and return the trimmed right subtree
		return trimBst(node.Right, min, max)
	}

	if node.Value > max {
		// Node too large, trim/skip it and return the trimmed left subtree
		return trimBst(node.Left, min, max)
	}

	// min <= node.Value <= max
	// Node is within bounds → keep it and trim children
	node.Left = trimBst(node.Left, min, max)   // can assign nil (and trim subtree) or keep the same node as Left (with possibly some of its subtrees trimmed)
	node.Right = trimBst(node.Right, min, max) // can assign nil (and trim subtree) or keep the same node as Right (with possibly some of its subtrees trimmed)

	return node
}
