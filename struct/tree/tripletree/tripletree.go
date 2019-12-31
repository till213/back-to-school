package tripletree

// Node is a triple tree (traversal) element: left and
// right point to the node itself, in case there are no
// left and/or right children
type Node struct {
	value       int
	visited     int
	left, right *Node
}

// New creates a new root node, each left/right pointer
// pointing to itself
func New(value int) *Node {
	n := new(Node)
	n.value = value
	n.visited = 0
	n.left = n
	n.right = n
	return n
}

// IsLeaf returns true if n does not have any children
// (i.e. its left and right pointers point to n itself)
func (n *Node) IsLeaf() bool {
	return n.left == n && n.right == n
}

// Add adds the value to this triple tree (in a binary search
// order)
func Add(n **Node, value int) {
	if *n == nil {
		*n = New(value)
	} else if value <= (*n).value {
		if (*n).left == *n {
			(*n).left = nil
		}
		Add(&(*n).left, value)
	} else {
		if (*n).right == *n {
			(*n).right = nil
		}
		Add(&(*n).right, value)
	}
}

// Value returns the value of the node
func (n *Node) Value() int {
	return n.value
}

// Visited returns the number of times this node
// has been visited
func (n *Node) Visited() int {
	return n.visited
}

// SetVisited sets the number of times this node
// has been visited
func (n *Node) SetVisited(visited int) {
	n.visited = visited
}

// TTT implements a triple tree traversal,
// calling f(n) for each node. Each node is visited
// three times (hence the name), with auxiliary memory O(1)
// (no stack or any other auxiliary space)
func (n *Node) TTT(f func(*Node)) {
	// o lags always "one behind" o
	var o *Node
	p := n
	for p != nil {
		// called three times
		f(p)
		// aux q is used to "advance the o, p fork"
		q := p.left
		// rotate left pointer
		p.left = p.right
		// rotate right pointer
		p.right = o
		o, p = p, q
	}
}
