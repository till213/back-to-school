package main

import (
	"fmt"

	"github.com/till213/back-to-school/data/tree"
)

func isValidBST(n *tree.Node) bool {
	var leftValid, rightValid bool
	if n == nil || n.Left == nil && n.Right == nil {
		return true
	} else if n.Left != nil && n.Left.Value > n.Value || n.Right != nil && n.Right.Value < n.Value {
		return false
	}

	leftValid = isValidBST(n.Left)
	rightValid = isValidBST(n.Right)
	return leftValid && rightValid
}

func main() {
	sorted := []int{1, 2, 3, 4}
	root := tree.CreateBinarySearchTree(sorted, nil)
	valid := isValidBST(root)
	fmt.Println("Tree (sorted) is a valid Binary Search Tree:", valid)

	unsorted := []int{3, 2, 1}
	root = tree.CreateBinarySearchTree(unsorted, nil)
	valid = isValidBST(root)
	fmt.Println("Tree (unsorted) is a valid Binary Search Tree:", valid)
}
