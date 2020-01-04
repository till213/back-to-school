package main

import (
	"fmt"

	"github.com/till213/back-to-school/data/tree"
)

func compare(t1, t2 *tree.Node) bool {

	if t1 == nil && t2 == nil {
		// both (sub-)trees have same length and identical
		// structure, we are done
		return true
	}
	if t1 == nil || t2 == nil {
		// either tree is shorter than the other -> not equal
		return false
	}
	if t1.Value != t2.Value {
		// Values don't match
		return false
	}

	return compare(t1.Left, t2.Left) && compare(t1.Right, t2.Right)
}

func subTree(t1, t2 *tree.Node) bool {
	if t1 == nil {
		return false
	}
	if t1.Value == t2.Value && compare(t1, t2) {
		return true
	}
	return subTree(t1.Right, t2) || subTree(t1.Right, t2)
}

// Space: O(log n + log m), where n and m nodes in larger and smaller tree
// Time: O (n + km), where k (expected) occurrence of t2's root in t1
func containsTree(t1, t2 *tree.Node) bool {
	if t2 == nil {
		return true
	}
	return subTree(t1, t2)
}

func main() {
	sorted := []int{1, 2, 3, 4, 5, 6, 7}
	t1 := tree.CreateBinarySearchTree(sorted, nil)

	sorted2 := []int{5, 6, 7}
	t2 := tree.CreateBinarySearchTree(sorted2, nil)
	res := subTree(t1, t2)
	fmt.Println("T2 is subtree of T1", res)

	sorted3 := []int{5, 6, 8}
	t3 := tree.CreateBinarySearchTree(sorted3, nil)
	res = subTree(t1, t3)
	fmt.Println("T3 is subtree of T1", res)
}
