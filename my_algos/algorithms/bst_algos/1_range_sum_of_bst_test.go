/*
RANGE SUM OF BST – RECURSIVE OPTIMIZED

🧠 PROBLEM:
Given the root of a Binary Search Tree and two integers `low` and `high`, return
the sum of all node values `x` such that `low <= x <= high`.

The input tree is a valid BST.

🔍 INTUITION:
- We take advantage of BST properties:
  - All values in left subtree < current node
  - All values in right subtree > current node

- This allows us to prune (skip) unnecessary subtrees:
  - If node.Val < low → discard left subtree (everything is too small)
  - If node.Val > high → discard right subtree (everything is too large)

💡 KEY IDEA:
This is a modified DFS where we only explore paths that could contain valid values.

🧩 TIME COMPLEXITY:
- Best case (heavy pruning): O(log n)
- Worst case (no pruning / all values in range - we explored the whole tree): O(n)

🧮 SPACE:
- Recursive stack space: O(h), where h is tree height
*/

package bst_algos

import (
	"fmt"
	"panca.com/algo/bst"
	"testing"
)

func Test_RangeSumBST(t *testing.T) {
	data := []int{1, 4, 2, 5, 3, 7, 12, 45, 23}
	bst := bst.BuildBSTFromSlice(data)
	low := 4
	high := 10
	fmt.Println(sumNodesInRange(bst, low, high)) // 4 + 5 + 7 = 16
}

func sumNodesInRange(tree *bst.BST[int], low, high int) int {
	return sumNodes(tree.Root, low, high)
}

func sumNodes(node *bst.Node[int], low, high int) int {
	if node == nil {
		return 0
	}

	if node.Value < low { // // Skip left subtree, all values are too small
		// If the current node’s value is less than the lower bound, then all values in the left subtree are even smaller (because it's a BST)
		// so we skip the whole left subtree and instead we go to the Right child which may or may not be in the range bounds
		return sumNodes(node.Right, low, high)
	}

	if node.Value > high { // Skip right subtree, all values are too large
		// if node is higher than high bound, skip right subtree and go left
		return sumNodes(node.Left, low, high)
	}

	// Current node is in range: include its value and recurse both sides
	// low <= node.Value <= high
	// this node is withing bounds, we add its value to the sum and recurse both left and right
	return node.Value + sumNodes(node.Left, low, high) + sumNodes(node.Right, low, high)
}
