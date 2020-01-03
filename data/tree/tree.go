package tree

// Node is a node in the tree.
type Node struct {
	Value       int
	Left, Right *Node
}

// CreateBinarySearchTree creates a balanced binary search
// tree, given the sorted values. Returns the root node.
func CreateBinarySearchTree(sortedValues []int) *Node {
	var n *Node
	len := len(sortedValues)
	switch len {
	case 0:
		n = nil
	case 1:
		n = new(Node)
		n.Value = sortedValues[0]
	case 2:
		n = new(Node)
		n.Value = sortedValues[0]
		n = new(Node)
		n.Value = sortedValues[1]
	default:
		n = new(Node)
		n.Value = sortedValues[len / 2]
		n.Left = CreateBinarySearchTree(sortedValues[0:len / 2])
		n.Right = CreateBinarySearchTree(sortedValues[len / 2 + 1:])
	}
	return n
}
