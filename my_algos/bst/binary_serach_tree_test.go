package bst

import (
	"fmt"
	"testing"
)

func Test_BST(t *testing.T) {
	// Unsorted username
	usernames := []string{"David", "Alice", "Maggie", "Zack", "John", "Manning"}
	bst := BuildBSTFromSlice(usernames)

	fmt.Print("In-order traversal: ")
	fmt.Println(bst.InOrderTraversal()) // [Alice David John Maggie Manning Zack]

	fmt.Println("Search for Maggie:", bst.Search("Maggie")) // true
	fmt.Println("Search for Bob:", bst.Search("Bob"))       // false

	bst.Delete("Maggie")
	fmt.Print("After deleting Maggie: ")
	fmt.Println(bst.InOrderTraversal()) // [Alice David John Manning Zack]
}
