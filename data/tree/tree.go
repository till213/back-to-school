package tree

import (
	"fmt"
	"math/rand"
)

// Node is a node in the tree, having a parent pointer.
type Node struct {
	Value       int
	Left, Right *Node
	Parent      *Node
}

// CreateBinarySearchTree creates a balanced binary search
// tree, given the sorted values. Returns the root node.
func CreateBinarySearchTree(sortedValues []int, parent *Node) *Node {
	var n *Node
	len := len(sortedValues)
	switch len {
	case 0:
		n = nil
	case 1:
		n = new(Node)
		n.Value = sortedValues[0]
		n.Parent = parent
	case 2:
		n = new(Node)
		n.Value = sortedValues[0]
		n.Parent = parent
		n.Right = new(Node)
		n.Right.Value = sortedValues[1]
		n.Right.Parent = parent
	default:
		n = new(Node)
		n.Value = sortedValues[len/2]
		n.Parent = parent
		n.Left = CreateBinarySearchTree(sortedValues[0:len/2], n)
		n.Right = CreateBinarySearchTree(sortedValues[len/2+1:], n)
	}
	return n
}

func createRandomTree(i, min, max, maxDepth int, parent *Node) *Node {
	n := new(Node)
	n.Value = min + rand.Intn(max-min+1)
	n.Parent = parent
	i++
	if i < maxDepth {
		if rand.Float64() < 0.9 {
			n.Left = createRandomTree(i, min, max, maxDepth, n)
		}
		if rand.Float64() < 0.9 {
			n.Right = createRandomTree(i, min, max, maxDepth, n)
		}
	}
	return n
}

// CreateRandomTree creates a random tree with a maximum Depth.
func CreateRandomTree(min, max, maxDepth int, parent *Node) *Node {
	return createRandomTree(0, min, max, maxDepth, parent)
}

// PrintTree prints the tree.
func PrintTree(node *Node, h int) {
	if node != nil {
		PrintTree(node.Right, h+1)
		for i := 0; i < h; i++ {
			fmt.Print("\t")
		}
		fmt.Printf("%4d\n", node.Value)
		PrintTree(node.Left, h+1)
	}
}
