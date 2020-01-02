package graph

import (
	"fmt"
	"testing"
)

func TestPrim(t *testing.T) {
	g := DAG()
	root := g.Vertex("C")
	g.Prim(root)

	fmt.Println("---- Minimum Spanning Tree (Prim) ----")
	fmt.Println("Root vertex:", root.name)
	fmt.Println(g.ToString())
}
