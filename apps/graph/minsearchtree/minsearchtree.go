package main

import "fmt"

type node struct {
	value       int
	left, right *node
}

func createBinarySearchTree(sortedValues []int) *node {
	var n *node
	len := len(sortedValues)
	switch len {
	case 0:
		n = nil
	case 1:
		n = new(node)
		n.value = sortedValues[0]
	case 2:
		n = new(node)
		n.value = sortedValues[0]
		n.right = new(node)
		n.right.value = sortedValues[1]
	default:
		n = new(node)
		n.value = sortedValues[len / 2]
		n.left = createBinarySearchTree(sortedValues[0:len / 2])
		n.right = createBinarySearchTree(sortedValues[len / 2 + 1:])
	}
	return n
}

func main() {
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := createBinarySearchTree(sorted)
	fmt.Println("Root", root)
}
