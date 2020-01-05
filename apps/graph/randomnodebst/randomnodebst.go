package main

import (
	"fmt"
	"math/rand"
)

// Node is an element of the BST.
type Node struct {
	value       int
	left, right *Node
	// the node count of all children, including
	// this node
	count int
}

func (n *Node) random(r int) *Node {
	var rn *Node
	var leftCount int
	if n.left != nil {
		leftCount = n.left.count
	} else {
		leftCount = 0
	}
	if r < leftCount {
		rn = n.left.random(r)
	} else if r == leftCount {
		rn = n
	} else {
		// the original random vale minus the count
		// of the left tree and minus this node (+1)
		rr := r - (leftCount + 1)
		rn = n.right.random(rr)
	}
	return rn
}

// Random returns a random note (including this node) of the BST,
// with each node having equal probability of being selected.
func (n *Node) Random() *Node {
	// [0..n[
	r := rand.Intn(n.count)
	return n.random(r)
}

// Add adds a new value to the BST.
func Add(n **Node, value int) {
	if *n == nil {
		*n = new(Node)
		(*n).value = value
		(*n).count = 1
		return
	}

	(*n).count++
	if value <= (*n).value {
		Add(&(*n).left, value)
	} else {
		Add(&(*n).right, value)
	}
}

func main() {

	var root *Node
	Add(&root, 2)
	Add(&root, 1)
	Add(&root, 3)

	for i := 0; i < 10; i++ {
		fmt.Println("Random node:", root.Random().value)
	}

}
