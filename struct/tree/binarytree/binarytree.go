package main

import "fmt"

type node struct {
	value       int
	left, right *node
}

// Creates a balanced (unordered) tree in a recursive manner, returning
// its root node
func createTree(values []int) *node {
	n := len(values)
	if n == 0 {
		return nil
	}
	nleft := n >> 1
	node := new(node)
	node.value = values[0]
	if n > 1 {
		node.left = createTree(values[1 : nleft+1])
		node.right = createTree(values[nleft+1:])
	}
	return node
}

func (n *node) copyTree() *node {
	var cp, q, p *node

	p = nil
	q = n
	for q != nil {
		cp = new(node)
		cp.value = q.value
		// Right points to copied parent
		cp.right = p
		// Left points to original right node (if it has one),
		// still to be copied
		if q.right != nil {
			cp.left = q.right
		}
		// Advance left
		p, q = cp, q.left
	}
	// post:
	// * we have reached the bottom left node of the subtree: p
	// * all right pointers point to parent node
	// * all left pointers point to right node of original subtree (if exists)
	q, p = p, nil
	for q != nil {
		// Set left pointer to child
		l := q.left
		q.left = p

		// Step up
		q, p = q.right, q
		p.right = l.copyTree()
	}

	return p
}

// Prints the binary tree given with root node.
// Root is on the left, right nodes in the upper half,
// tree is printed left (root) to right (leaves)
// ("rotated 90 degrees anti-clockwise")
func printTree(node *node, h int) {
	if node != nil {
		printTree(node.right, h+1)
		for i := 0; i < h; i++ {
			fmt.Print("\t")
		}
		fmt.Printf("%6d\n", node.value)
		printTree(node.left, h+1)
	}
}

func main() {
	values := []int{8, 9, 11, 15, 19, 20, 21, 7, 3, 2, 1, 5, 6, 4, 13, 14, 10, 12, 17, 16, 18}
	root := createTree(values)
	fmt.Println("--- Original ---")
	printTree(root, 0)
	cp := root.copyTree()
	fmt.Println("--- Copy ---")
	printTree(cp, 0)
}
