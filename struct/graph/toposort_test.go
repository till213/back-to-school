package graph

import (
	"fmt"
	"testing"
)

func TestToposortDAG(t *testing.T) {
	g := DAG()

	vertices := TopoSort(g)
	c := len(vertices)
	exp := g.VertexCount()
	if c != exp {
		t.Errorf("Number of toposorted vertices, expected %v, received %v", exp, c)
	}

	for _, v := range vertices {
		fmt.Print("->", v.Name)
	}
	fmt.Println()
}
