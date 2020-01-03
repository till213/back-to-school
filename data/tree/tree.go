package tree

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
		n.Value = sortedValues[len / 2]
		n.Parent = parent
		n.Left = CreateBinarySearchTree(sortedValues[0:len / 2], n)
		n.Right = CreateBinarySearchTree(sortedValues[len / 2 + 1:], n)
	}
	return n
}
