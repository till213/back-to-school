package graph

import (
	"fmt"
	"testing"
)

func TestGraphAddVertices(t *testing.T) {
	g := DAG()
	c := g.VertexCount()
	exp := len(NamedVertices)
	if c != exp {
		t.Errorf("Number of vertices, expected %v, received %v", exp, c)
	}

	fmt.Println(g.ToString())
}
