package main

import (
	"fmt"

	"github.com/till213/back-to-school/data/tree"
)

func firstCommonAncestor(n, c1, c2 *tree.Node) (fca *tree.Node, c1Exists, c2Exists bool) {

	var c1ExistsLeft, c2ExistsLeft, c1ExistsRight, c2ExistsRight bool
	if n == nil || c1 == nil || c2 == nil {
		return nil, false, false
	}

	if n.Left != nil {
		fca, c1ExistsLeft, c2ExistsLeft = firstCommonAncestor(n.Left, c1, c2)
	}
	if fca == nil && n.Right != nil {
		fca, c1ExistsRight, c2ExistsRight = firstCommonAncestor(n.Right, c1, c2)
	}
	c1Exists = c1ExistsLeft || c1ExistsRight
	c2Exists = c2ExistsLeft || c2ExistsRight

	// At this point:
	// - First common ancestor exists in either left or right subtree (we are done)
	// - No first common ancestor found in any subtree yet, in which case we check
	//   existence of children 1 and 2
	// - ...which are either in different or even the same subtree (in case c1 is parent
	//   of c2 or vice versa)
	// - If both children exists, then this n is their first common ancestor
	// - Otherwise (zero or just on child exists) check if the current n is
	//   one of c1, c2 and "report it up"
	//
	// So in post order manner:
	if fca == nil {

		// check if children c1 and c2 exist
		if c1Exists && c2Exists {
			// both children exist in any of the left, right subtrees
			// -> this n is our first common ancestor
			fca = n
		} else {
			// check if n is any of the children (special case: or even both)
			if n == c1 {
				c1Exists = true
			}
			if n == c2 {
				c2Exists = true
			}
		}
	}
	return fca, c1Exists, c2Exists
}

func main() {
	// For this solution the tree does not need to be a binary *search*
	// tree - we use one anyway
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16}
	root := tree.CreateBinarySearchTree(sorted, nil)

	c1 := root.Left.Left
	c2 := root.Left.Right
	fca, _, _ := firstCommonAncestor(root, c1, c2)

	if fca != nil {
		fmt.Printf("First common ancestor of %v and %v: %v\n", c1.Value, c2.Value, fca.Value)
	} else {
		fmt.Printf("No common ancestor found for: %v, %v\n", c1, c2)
	}

	c1 = root.Left
	c2 = root.Left
	fca, _, _ = firstCommonAncestor(root, c1, c2)

	if fca != nil {
		fmt.Printf("First common ancestor of %v and %v: %v\n", c1.Value, c2.Value, fca.Value)
	} else {
		fmt.Printf("No common ancestor found for: %v, %v\n", c1, c2)
	}
	c1 = root.Left
	c2 = root.Right.Left
	fca, _, _ = firstCommonAncestor(root, c1, c2)

	if fca != nil {
		fmt.Printf("First common ancestor of %v and %v: %v\n", c1.Value, c2.Value, fca.Value)
	} else {
		fmt.Printf("No common ancestor found for: %v, %v\n", c1, c2)
	}

	c1 = root
	c2 = root.Right
	fca, _, _ = firstCommonAncestor(root, c1, c2)

	if fca != nil {
		fmt.Printf("First common ancestor of %v and %v: %v\n", c1.Value, c2.Value, fca.Value)
	} else {
		fmt.Printf("No common ancestor found for: %v, %v\n", c1, c2)
	}
}
