/*
Binary Search Tree (BST)

🧠 DEFINITION:
A Binary Search Tree (BST) is a binary tree where for every node:
- Each node can have 0, 1 or 2 children - Left and Right (binary tree)
- All values in the left node < node.Value
- All values in the right node > node.Value
This holds true recursively for every subtree, i.e. for every node:
- All values in the left subtree must be strictly less than the node’s value.
- All values in the right subtree must be strictly greater than the node’s value.

✅ OPERATIONS (in this example):
- Insert: O(log n) average, O(n) worst case
- Search: same as insert
- Delete: handles 3 cases (0, 1, or 2 children)
- InOrderTraversal: returns sorted values, we traverse all nodes -> O(n)
- BuildBSTFromSlice: helper to build tree from unsorted slice

📦 USE CASES:
- BST is useful for ordered data where you want reasonably fast insert, search, and delete (average O(log n)).
- database index example:
	- Suppose you're indexing usernames in a database (like Facebook).
	- B-Tree (a self-balancing generalization of BST) might be used to index users by username.
	- In that case, alphabetically earlier usernames go to the left, and later ones to the right.

⚖️ COMPARISON TO SORTED ARRAY:
- We use BST instead of sorted array when we need faster insertion and deletion.
- Remember that for sorted arrays, search is fast - binary search with O(log n) time, but insert and delete are
slow because we need to shift elements - O(n) time.
- Note that array will always be faster for access - O(1) so we use BST instead of sorted arrays when we don't need
fast access

- Search: Sorted array = O(log n), BST = O(log n) avg, O(n) worst
- Insert/Delete: Sorted array = O(n) (due to shifting), BST = O(log n) avg
- Access: Sorted array = O(1), BST = O(log n)

🧠 TIME COMPLEXITY NOTE:
- In BST all operations take on average O(log n) time, in the worst case it's O(n) if the tree is unbalanced
- imagine a sorted array that we converted to BST - all nodes will go to the right of the previous one so we will effectively
get a linked list, ex: [1] -> [2] -> [3] -> ...
- to avoid this, AVL trees or Red-Black trees (self-balancing BSTs) are used
- AVL trees maintain balance by performing rotations during insertion/deletion to keep the tree height log(n) but this discussion
is out of scope for our needs...
- just remember that balanced trees (AVL, Red-Black) guarantee O(log n) for BST operations, but standard BST without balancing does not

Note: BST is basis for many interview questions (e.g., kth smallest, range sum)
*/

package bst

import (
	"fmt"
	"strings"
)

type Ordered interface {
	~int | ~int64 | ~float64 | ~string // these support <, <=, >, >= and == comparator operators
}

// Node is a node in BST
type Node[T Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// BST wraps the root node of a binary search tree
type BST[T Ordered] struct {
	Root *Node[T]
}

func (t *BST[T]) String() string {
	return formatTree(t.Root, 0)
}

func formatTree[T Ordered](node *Node[T], depth int) string {
	if node == nil {
		return ""
	}
	indent := strings.Repeat("  ", depth)
	return fmt.Sprintf("%s%v\n%s%s", indent, node.Value, formatTree(node.Left, depth+1), formatTree(node.Right, depth+1))
}

func NewBST[T Ordered]() *BST[T] {
	return &BST[T]{}
}

func NewBSTFromRootNode[T Ordered](rootNode *Node[T]) *BST[T] {
	return &BST[T]{Root: rootNode}
}

// Insert inserts a value into the BST
func (t *BST[T]) Insert(value T) {
	t.Root = insert(t.Root, value)
}

func insert[T Ordered](node *Node[T], value T) *Node[T] {
	if node == nil {
		return &Node[T]{Value: value}
	}
	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	} // duplicate values are ignored
	return node
}

// Search returns true if the value exists in the tree
func (t *BST[T]) Search(value T) bool {
	return search(t.Root, value)
}

func search[T Ordered](node *Node[T], value T) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	} else if value < node.Value {
		return search(node.Left, value)
	} else {
		return search(node.Right, value)
	}
}

// Delete removes a value from the BST and restructures the tree to maintain BST property
func (t *BST[T]) Delete(value T) {
	t.Root = Delete(t.Root, value)
}

func Delete[T Ordered](node *Node[T], value T) *Node[T] {
	if node == nil {
		return nil
	}

	if value < node.Value {
		// Continue searching in the left subtree
		node.Left = Delete(node.Left, value)
	} else if value > node.Value {
		// Continue searching in the right subtree
		node.Right = Delete(node.Right, value)
	} else { // value == node.Value
		// Found the node to delete

		// Case 1: Node has no children (leaf node)
		if node.Left == nil && node.Right == nil {
			return nil // set the parent's Right/Left reference to nil
		}

		// Case 2: Node has one child
		if node.Left == nil { // node.Right != nil
			// set parent's Right/Left reference to this child node effectively deleting this node - it is replaced with its only child
			return node.Right
		}
		if node.Right == nil { // node.Left != nil
			// same as above
			return node.Left
		}

		// Case 3: Node has two children - we can just delete it because there are nodes which depend on it
		// Replace this node with the in-order successor (smallest node in the right subtree)
		// Transitivity (A < B & B < C -> A < C) ensures correctness: it's safe to replace a node with its in-order successor, BST property remains valid
		minRight := FindMin(node.Right)
		node.Value = minRight.Value

		// Now delete the in-order successor (which is now duplicated)
		node.Right = Delete(node.Right, minRight.Value)
	}

	return node
}

// FindMin returns the node with the smallest value in the subtree
func FindMin[T Ordered](node *Node[T]) *Node[T] {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// InOrderTraversal returns values in sorted order, it always starts from root
func (t *BST[T]) InOrderTraversal() []T {
	var result []T
	inOrder(t.Root, &result)
	return result
}

func inOrder[T Ordered](node *Node[T], result *[]T) {
	if node == nil {
		return
	}
	inOrder(node.Left, result)
	*result = append(*result, node.Value)
	inOrder(node.Right, result)
}

// BuildBSTFromSlice builds a BST from a given unsorted string slice
// input: "David", "Alice", "Maggie", "Zack", "John", "Manning"
// output:
//
//	      David
//	    /     \
//	Alice     Maggie
//	         /     \
//	     John      Zack
//	              /
//	         Manning
func BuildBSTFromSlice[T Ordered](values []T) *BST[T] {
	tree := NewBST[T]()
	for _, val := range values {
		tree.Insert(val)
		// nil, David -> root = David
		// David, Alice -> root = David (David sad ima Left child = Alice cause Alice < David)
		// David, Maggie -> root = David (David sad ima Right child = Maggie cause Maggie > David)
		// David, Zack -> root = David (Maggie sad ima Right child = Zack cause Zack > Maggie)
		// David, John -> root = David (Maggie sad ima Left child = John cause John > David, < Maggie)
		// David, Manning -> root = David (Zack sad ima Left child = Manning cause Manning > David, Manning > Maggie, Manning < Zack)
	}
	return tree
}
