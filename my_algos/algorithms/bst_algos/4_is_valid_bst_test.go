/*
🧠 Validate Binary Search Tree (BST)

PROBLEM:
Given the root of a binary tree, determine if it is a valid BST.

✅ DEFINITION:
A binary search tree (BST) is a binary tree where for every node:
- All values in the left subtree < node.Value
- All values in the right subtree > node.Value
These constraints must hold for the entire subtree, not just direct children.

👀 KEY IDEA:
- We track a range of valid values: `min` and `max`
- For each node:
  - Its value must be strictly greater than `min` and strictly less than `max`
  - When going left, we update `max = node.Value`
  - When going right, we update `min = node.Value`

📈 TIME & SPACE:
- Time: O(n), each node is visited once
- Space: O(h), due to recursion stack (h = height of tree)
*/

package bst_algos

import (
	"fmt"
	"math"
	"panca.com/algo/bst"
	"testing"
)

func Test_IsValidBST(t *testing.T) {
	// ✅ TEST CASE 1: A valid BST
	data := []int{20, 10, 30, 5, 15, 25, 35, 3, 7, 13, 17}
	validTree := bst.BuildBSTFromSlice(data)
	fmt.Println(isValidBST(validTree.Root, -math.MaxInt, math.MaxInt)) // true

	// ❌ TEST CASE 2: Not a valid BST (left child > parent)
	invalid := &bst.Node[int]{
		Value: 10,
		Left: &bst.Node[int]{
			Value: 12, // Invalid: left > parent
		},
		Right: &bst.Node[int]{
			Value: 15,
		},
	}
	fmt.Println(isValidBST(invalid, -math.MaxInt, math.MaxInt)) // false

	// ❌ TEST CASE 3: Not a valid BST (duplicate value on right)
	duplicate := &bst.Node[int]{
		Value: 10,
		Left:  &bst.Node[int]{Value: 5},
		Right: &bst.Node[int]{
			Value: 10, // Invalid: duplicate on right
		},
	}
	fmt.Println(isValidBST(duplicate, -math.MaxInt, math.MaxInt)) // false

	// ✅ TEST CASE 4: Minimal valid BST (single node)
	single := &bst.Node[int]{Value: 42}
	fmt.Println(isValidBST(single, -math.MaxInt, math.MaxInt)) // true
}

// isValidBST recursively validates the BST using min/max range check.
// The node value must be in the range (min, max), exclusive.
func isValidBST[T bst.Ordered](node *bst.Node[T], min, max T) bool {
	if node == nil {
		return true // base case: empty subtree is valid
	}

	// current node must be strictly between min and max
	if node.Value <= min || node.Value >= max {
		return false
	}

	// recursively validate left and right with updated bounds
	// Recurse left: all values < node.Value
	// Recurse right: all values > node.Value
	// isValidBST must be true for both in order for the tree to be valid, we also have an early exit mechanism for the
	// left recursion - if it returns false, right recursion won't be called
	return isValidBST(node.Left, min, node.Value) && isValidBST(node.Right, node.Value, max)
}
