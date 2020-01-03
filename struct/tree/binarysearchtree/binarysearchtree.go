package binarysearchtree

// Node is a binary search tree element
type Node struct {
	value       int
	left, right *Node
}

// BinarySearchTree implements an (unbalanced)
// binary search tree:
// value of left child <= parent < right child
// (which allows for multiple elements with the
// same value)
type BinarySearchTree struct {
	root *Node
}

// New creates and initialises a binary search tree.
func New() *BinarySearchTree {
	return new(BinarySearchTree)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func add(n **Node, value int) {
	if *n == nil {
		*n = new(Node)
		(*n).value = value
	} else if value <= (*n).value {
		add(&(*n).left, value)
	} else {
		add(&(*n).right, value)
	}
}

// Value returns the node value
func (n *Node) Value() int {
	return n.value
}

func (n *Node) find(value int) *Node {
	if n == nil {
		return nil
	}
	if n.value == value {
		return n
	} else if value < n.value {
		return n.left.find(value)
	}
	return n.right.find(value)
}

// Copy recursively copies this node n and all its children,
// returning the copy of n.
func (n *Node) copy() *Node {
	c := new(Node)
	c.value = n.value
	if n.left != nil {
		c.left = n.left.copy()
	}
	if n.right != nil {
		c.right = n.right.copy()
	}
	return c
}

func (n *Node) nofChildren() int {
	var count int
	if n == nil {
		return 0
	}
	if n.left != nil {
		if n.right != nil {
			count = 2
		} else {
			count = 1
		}
	} else {
		if n.right != nil {
			count = 1
		} else {
			count = 0
		}
	}
	return count
}

func (n *Node) remove(parent **Node, value int) {
	if n.value == value {
		switch n.nofChildren() {
		case 0:
			if n != *parent {
				if n.value <= (*parent).value {
					(*parent).left = nil
				} else {
					(*parent).right = nil
				}
			} else {
				*parent = nil
			}
		case 1:
			if n.left != nil {
				if n.value <= (*parent).value {
					(*parent).left = n.left
				} else {
					(*parent).right = n.left
				}
			} else {
				if n.value <= (*parent).value {
					(*parent).left = n.right
				} else {
					(*parent).right = n.right
				}
			}
		default:
			// Two children, replace with largest element
			// in left subtree, pointed at by t, with parent p
			p := n
			t := n.left
			for t.right != nil {
				p = t
				t = t.right
			}
			// Replace n's value with the found maximum value
			// of left subtree
			n.value = t.value
			// t - being the largest element in the left subtree -
			// has at most one child, which is removed as case 0
			// or 1 above
			switch t.nofChildren() {
			case 0:
				if t.value <= p.value {
					p.left = nil
				} else {
					p.right = nil
				}
			default:
				// At most one child, which must be a left child
				// of t
				p.right = t.left
			}
		}

	} else if value < n.value {
		n.left.remove(&n, value)
	} else {
		n.right.remove(&n, value)
	}
}

func (n *Node) maximumDepth() int {
	if n == nil {
		return -1
	}
	return 1 + max(n.left.maximumDepth(), n.right.maximumDepth())
}

func (n *Node) nofNodes() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.nofNodes() + n.right.nofNodes()
}

func bfs(queue []*Node, f func(*Node)) {
	if len(queue) == 0 {
		return
	}
	n := queue[0]
	f(n)
	if n.left != nil {
		queue = append(queue, n.left)
	}
	if n.right != nil {
		queue = append(queue, n.right)
	}
	bfs(queue[1:], f)
}

// dfs iterates in-order: left, f(node), right
func (n *Node) dfs(f func(*Node)) {
	if n.left != nil {
		n.left.dfs(f)
	}
	f(n)
	if n.right != nil {
		n.right.dfs(f)
	}
}

// Add adds the value
func (b *BinarySearchTree) Add(value int) {
	add(&b.root, value)
}

// Find returns the Node containing value. Nil if not found.
func (b *BinarySearchTree) Find(value int) *Node {
	return b.root.find(value)
}

// Remove removes the value from the tree (if it exists).
func (b *BinarySearchTree) Remove(value int) {
	b.root.remove(&b.root, value)
}

// Copy creates and returns a deep copy of this binary
// search tree.
func (b *BinarySearchTree) Copy() *BinarySearchTree {
	if b.root != nil {
		c := New()
		c.root = b.root.copy()
		return c
	}
	return nil
}

// MaximumDepth returns the maximum depth of the binary search tree,
// where root (1 element) has depth 0
func (b *BinarySearchTree) MaximumDepth() int {
	return b.root.maximumDepth()
}

// NofNodes returns the number of nodes in the binary search tree
func (b *BinarySearchTree) NofNodes() int {
	return b.root.nofNodes()
}

// BFS iterates over the tree in a Breadth-First Search manner,
// calling f(n) for each node
func (b *BinarySearchTree) BFS(f func(*Node)) {
	if b.root != nil {
		queue := []*Node{b.root}
		bfs(queue, f)
	}
}

// DFS iterates over the tree in a Depth-First Search manner,
// calling f() for each node
func (b *BinarySearchTree) DFS(f func(*Node)) {
	if b.root != nil {
		b.root.dfs(f)
	}
}
