package main

import (
	"fmt"
	"container/list"

	"github.com/till213/back-to-school/data/tree"
)

func depthFirst(n *tree.Node, depth int, lod *[]*list.List) {
	if n == nil {
		return
	}

	// Add level of depth, if necessary
	if len(*lod) < depth {
		*lod = append(*lod, list.New())
	}

	(*lod)[depth - 1].PushBack(n)
	if n.Left != nil {
		depthFirst(n.Left, depth + 1, lod)
	}
	if n.Right != nil {
		depthFirst(n.Right, depth + 1, lod)
	}
}

func listOfDepth(root *tree.Node) []*list.List {
	lod := make([]*list.List, 0, 8)
	if root != nil {
		depthFirst(root, 1, &lod)
	}
	return lod
}

func main() {
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16}
	root := tree.CreateBinarySearchTree(sorted)
	
	lod := listOfDepth(root)

	for i, v := range lod {
		fmt.Println("Depth", i+1)
		for e := v.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v -> ", e.Value.(*tree.Node).Value)
		}
		fmt.Println()
	}
	fmt.Println()
}