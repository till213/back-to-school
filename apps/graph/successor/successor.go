package main

import (
	"fmt"

	"github.com/till213/back-to-school/data/tree"
)

func nextInRightSubtree(n *tree.Node) *tree.Node {
	var successor *tree.Node
	if n.Right != nil {
		successor = n.Right
		// keep left, as far down the tree as possible
		for successor.Left != nil {
			successor = successor.Left
		}
	}
	return successor
}

// Returns the next parent which is a "left-child"; nil if
// no such parent exist (= we are done)
func nextParent(n *tree.Node) *tree.Node {
	var parent *tree.Node
	for parent == nil && n.Parent != nil {
		if n.Parent.Right != n {
			// n is a "left-child"
			parent = n.Parent
		} else {
			// advance up
			n = n.Parent
		}
	}
	return parent
}

func next(n *tree.Node) *tree.Node {
	var successor *tree.Node
	if n.Parent == nil {
		// n is root
		successor = nextInRightSubtree(n)
	} else {
		// n has parent
		if n.Right == nil {
			if n != n.Parent.Right {
				// n is the left child of parent
				successor = n.Parent
			} else {
				// n is the right child of parent, traverse all the
				// way up, until we find a left-child (if no such
				// left-child parent exist, we are done)
				np := nextParent(n)
				if np != nil {
					successor = nextInRightSubtree(np)
				} else {
					// no more successor
					successor = nil
				}
			}
		} else {
			// n has right child
			successor = n.Right
		}
	}
	return successor
}

func main() {
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	root := tree.CreateBinarySearchTree(sorted, nil)
	start := root
	for start.Left != nil {
		start = start.Left
	}
	fmt.Println("-- Inorder (left-most) --")
	for n := start; n != nil; n = next(n) {
		fmt.Print(n.Value, ", ")
	}
	fmt.Println()
	fmt.Println("-- Inorder (root) --")
	for n := root; n != nil; n = next(n) {
		fmt.Print(n.Value, ", ")
	}
	fmt.Println()
}
