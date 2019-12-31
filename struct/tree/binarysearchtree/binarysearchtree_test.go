package binarysearchtree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTreeFind(t *testing.T) {
	var tree *BinarySearchTree

	tree = New()
	tree.Add(2)
	tree.Add(1)
	tree.Add(3)

	n := tree.Find(1)
	if n == nil {
		t.Errorf("Did not find node for value 1")
	}
	n = tree.Find(2)
	if n == nil {
		t.Errorf("Did not find node for value 2")
	}
	n = tree.Find(3)
	if n == nil {
		t.Errorf("Did not find node for value 3")
	}
	n = tree.Find(4)
	if n != nil {
		t.Errorf("Did not expect a node for value 4")
	}
}

func TestBinarySearchTreeRemoveOneChild(t *testing.T) {
	var tree *BinarySearchTree

	// Setup
	tree = New()
	tree.Add(5)
	tree.Add(2)
	tree.Add(4)
	tree.Add(3)

	maxDepth := tree.MaximumDepth()
	if maxDepth != 3 {
		t.Errorf("Expected maximum depth: 3, received: %v", maxDepth)
	}
	nofNodes := tree.NofNodes()
	if nofNodes != 4 {
		t.Errorf("Expected node count: 4, received: %v", nofNodes)
	}

	// Exercise
	tree.Remove(2)

	// Verify
	maxDepth = tree.MaximumDepth()
	if maxDepth != 2 {
		t.Errorf("Expected maximum depth: 2, received: %v", maxDepth)
	}
	nofNodes = tree.NofNodes()
	if nofNodes != 3 {
		t.Errorf("Expected node count: 3, received: %v", nofNodes)
	}

	n := tree.Find(2)
	if n != nil {
		t.Errorf("Did not expect the deleted node with value 2")
	}
}

func TestBinarySearchTreeRemoveTwoChildren(t *testing.T) {
	var tree *BinarySearchTree

	// Setup
	tree = New()
	tree.Add(5)
	tree.Add(2)
	tree.Add(4)
	tree.Add(3)
	tree.Add(1)

	maxDepth := tree.MaximumDepth()
	if maxDepth != 3 {
		t.Errorf("Expected maximum depth: 3, received: %v", maxDepth)
	}
	nofNodes := tree.NofNodes()
	if nofNodes != 5 {
		t.Errorf("Expected node count: 5, received: %v", nofNodes)
	}

	// Exercise
	tree.Remove(2)

	// Verify
	maxDepth = tree.MaximumDepth()
	if maxDepth != 3 {
		t.Errorf("Expected maximum depth: 3, received: %v", maxDepth)
	}
	nofNodes = tree.NofNodes()
	if nofNodes != 4 {
		t.Errorf("Expected node count: 4, received: %v", nofNodes)
	}

	n := tree.Find(2)
	if n != nil {
		t.Errorf("Did not expect the deleted node with value 2")
	}
}

func TestBinarySearchTreeBFS(t *testing.T) {
	var tree *BinarySearchTree

	// Setup
	tree = New()
	tree.Add(5)
	tree.Add(2)
	tree.Add(4)
	tree.Add(3)
	tree.Add(1)
	tree.Add(8)
	tree.Add(9)

	// Exercise
	fmt.Println("---- BFS ----")
	tree.BFS(func(n *Node) { fmt.Println("Node value", n.Value()) })
}

func TestBinarySearchTreeDFS(t *testing.T) {
	var tree *BinarySearchTree

	// Setup
	tree = New()
	tree.Add(5)
	tree.Add(2)
	tree.Add(4)
	tree.Add(3)
	tree.Add(1)
	tree.Add(8)
	tree.Add(9)

	// Exercise
	fmt.Println("---- DFS ----")
	tree.DFS(func(n *Node) { fmt.Println("Node value", n.Value()) })
}
