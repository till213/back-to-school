package main

import (
	"fmt"

	"github.com/till213/back-to-school/data/tree"
)

func isBalanced(n *tree.Node) bool {
	var leftBalanced, rightBalanced bool
	if n == nil || n.Left == nil && n.Right == nil {
		return true
	} else if n.Left != nil && n.Right == nil || n.Left == nil && n.Right != nil {
		return false
	}

	leftBalanced = isBalanced(n.Left)
	rightBalanced = isBalanced(n.Right)
	return leftBalanced && rightBalanced
}


func main() {
	sorted := []int{1, 2, 3, 4}
	root := tree.CreateBinarySearchTree(sorted, nil)
	balanced := isBalanced(root)
	fmt.Println("Tree with 4 nodes is balanced:", balanced)

	sorted = []int{1, 2, 3}
	root = tree.CreateBinarySearchTree(sorted, nil)
	balanced = isBalanced(root)
	fmt.Println("Tree with 3 nodes is balanced:", balanced)
}
