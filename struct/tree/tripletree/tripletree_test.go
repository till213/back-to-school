package tripletree

import (
	"fmt"
	"testing"
)

func visitNode(n *Node) {
	n.SetVisited(n.Visited() + 1)
	fmt.Println("Node value: ", n.Value(), "visit:", n.Visited())
}

func TestTripleTreeTraverse(t *testing.T) {
	root := New(2)
	Add(&root, 1)
	Add(&root, 3)

	root.TTT(visitNode)
}
