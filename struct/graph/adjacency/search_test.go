package graph

import (
	"fmt"
	"testing"
)

func visit(v *Vertex) {
	fmt.Println(v.Name)
}

func TestDFS(t *testing.T) {
	g := DAG()

	fmt.Println("--- Iterative Depth First Search ---")
	DFS(g, g.Vertices["C"], visit)
}

func TestRecursiveDFS(t *testing.T) {
	g := DAG()

	fmt.Println("--- Recursive Depth First Search ---")
	RecursiveDFS(g, g.Vertices["C"], visit)
}

func TestBFS(t *testing.T) {
	g := DAG()

	fmt.Println("--- Iterative Breadth First Search ---")
	BFS(g, g.Vertices["C"], visit)
}
