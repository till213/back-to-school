package graph

import (
	"fmt"
	"testing"
)

func TestToposortDAG(t *testing.T) {
	g := DAG()

	vertices, acyclic := TopoSort(g)
	c := len(vertices)
	exp := g.VertexCount()
	if c != exp {
		t.Errorf("Number of toposorted vertices, expected %v, received %v", exp, c)
	}
	if !acyclic {
		t.Errorf("Acyclic graph expected, received %v", acyclic)
	}

	fmt.Println("---- Topological Sort (DAG) ----")
	for _, v := range vertices {
		fmt.Print("->", v.Name)
	}
	fmt.Println()
}

func TestToposortCycle(t *testing.T) {
	g := WithCycle()

	vertices, acyclic := TopoSort(g)

	if acyclic {
		t.Errorf("Cyclic graph expected, received %v", acyclic)
	}

	fmt.Println("---- Topological Sort (Cycle) ----")
	for _, v := range vertices {
		fmt.Print("->", v.Name)
	}
	fmt.Println()
}
