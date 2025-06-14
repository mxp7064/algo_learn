/*
🧠 Lowest Common Ancestor (LCA) in a Binary Search Tree (BST)

PROBLEM:
Given a BST and two values (a and b), return the lowest common ancestor (LCA) node.
The LCA is the lowest node in the tree that has both a and b as descendants.
Lowest means the deepest common node (closest to a and b), NOT the lowest value.
A node can also be a descendant of itself.

Nodes a and b can be in any kind of relationship: siblings, one parent to other, be very distant from each other, related only through root, etc.

BST PROPERTY:
- Left subtree contains only values < node.Value
- Right subtree contains only values > node.Value

STRATEGY (recursive approach):
- If both a and b < node → LCA is in the left subtree
- If both a and b > node → LCA is in the right subtree
- If a and b "split" (one on each side, or one equals the node) → current node is the LCA

TIME COMPLEXITY:
- Average: O(log n) if tree is balanced
- Worst: O(n) if tree is skewed
*/

package bst_algos

import (
	"fmt"
	"panca.com/algo/bst"
	"testing"
)

func Test_LowestCommonAncestor(t *testing.T) {
	// input BST
	data := []int{20, 10, 30, 5, 15, 25, 35, 3, 7, 13, 17}

	/*
	    Constructed BST:
	               20
	            /     \
	         10         30
	       /   \       /  \
	      5    15     25  35
	    /  \   / \
	   3    7 13 17
	*/

	tree := bst.BuildBSTFromSlice(data)

	// TEST CASE 1: Nodes are siblings → LCA is parent
	a, b := 13, 17 // siblings under 15
	expectLCA(t, tree.Root, a, b, 15)

	// TEST CASE 2: One node is ancestor of the other
	a, b = 10, 13 // 10 is ancestor of 13
	expectLCA(t, tree.Root, a, b, 10)

	// TEST CASE 3: Nodes very distant → LCA is higher up (in this case it's root)
	a, b = 7, 25 // left/right split
	expectLCA(t, tree.Root, a, b, 20)
}

// expectLCA is a helper that asserts the LCA result and prints a message
func expectLCA[T bst.Ordered](t *testing.T, root *bst.Node[T], a, b, expected T) {
	lca := lowestCommonAncestor(root, a, b)
	if lca == nil {
		t.Errorf("LCA of %v and %v not found", a, b)
	}
	if lca.Value != expected {
		t.Errorf("LCA of %v and %v should be %v, but got %v", a, b, expected, lca.Value)
	}
	fmt.Printf("✅ LCA(%v, %v) = %v\n", a, b, lca.Value)
}

// lowestCommonAncestor finds the LCA of two values in a BST
func lowestCommonAncestor[T bst.Ordered](node *bst.Node[T], a, b T) *bst.Node[T] {
	// Base case: empty tree
	if node == nil {
		return nil
	}

	// If both values are smaller than current node → LCA must be in the left subtree
	if a < node.Value && b < node.Value {
		return lowestCommonAncestor(node.Left, a, b)
	}

	// If both values are larger than current node → LCA must be in the right subtree
	if a > node.Value && b > node.Value {
		return lowestCommonAncestor(node.Right, a, b)
	}

	// If values split (a < current < b) or match current node (a = current or b = current) → current node is the LCA
	return node
}
