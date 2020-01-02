package main

import "fmt"

type mode int

const (
	preorder mode = iota
	inorder
)

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

// The morrisTreeTraversal traverses the tree with root n
// according to order: either preorder or inorder.
// In O(n) and with O(1) aux space.
func (n *node) morrisTreeTraversal(f func(n *node), order mode) {
	current := n
	for current != nil {
		if current.left == nil {
			f(current)
			current = current.right
		} else {
			// has left child
			predecessor := current.left
			for predecessor.right != current && predecessor.right != nil {
				predecessor = predecessor.right
			}
			if predecessor.right == nil {
				// we have not yet visited that left branch
				// so temporarily set a "backtracking path" to current...
				predecessor.right = current
				if order == preorder {
					f(current)
				}
				// ... and visit left branch
				current = current.left
			} else {
				// the "backtracking path" points to current,
				// which indicates that we have already visited that
				// left branch -> restore pointers...
				predecessor.right = nil
				if order == inorder {
					f(current)
				}
				// ... and proceed with right branch
				current = current.right
			}
		}
	}
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

func visit(n *node) {
	fmt.Printf("%v, ", n.value)
}

func main() {
	values := []int{8, 9, 11, 15, 19, 20, 21, 7, 3, 2, 1, 5, 6, 4, 13, 14, 10, 12, 17, 16, 18}
	root := createTree(values)
	fmt.Println("--- Original ---")
	printTree(root, 0)
	cp := root.copyTree()
	fmt.Println("--- Copy ---")
	printTree(cp, 0)
	fmt.Println("--- In-Order Traversal ---")
	root.morrisTreeTraversal(visit, inorder)
	fmt.Println()
	fmt.Println("--- Pre-Order Traversal ---")
	root.morrisTreeTraversal(visit, preorder)
	fmt.Println()
}
