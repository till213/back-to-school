package graph

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGraphAddVertices(t *testing.T) {
	g := New()
	namedVertices := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M"}
	g.AddVertices(namedVertices)
	g.AddEdge("A", "D", rand.Intn(20))
	g.AddEdge("B", "D", rand.Intn(20))
	g.AddEdge("C", "A", rand.Intn(20))
	g.AddEdge("C", "B", rand.Intn(20))
	g.AddEdges("D", []string{"H", "G"}, []int{rand.Intn(20), rand.Intn(20)})
	g.AddEdges("E", []string{"A", "D", "F"}, []int{rand.Intn(20), rand.Intn(20), rand.Intn(20)})
	g.AddEdges("F", []string{"K", "J"}, []int{rand.Intn(20), rand.Intn(20)})
	g.AddEdge("G", "I", rand.Intn(20))
	g.AddEdges("H", []string{"I", "J"}, []int{rand.Intn(20), rand.Intn(20)})
	g.AddEdge("I", "J", rand.Intn(20))
	g.AddEdges("J", []string{"L", "M"}, []int{rand.Intn(20), rand.Intn(20)})

	c := g.VertexCount()
	exp := len(namedVertices)
	if c != exp {
		t.Errorf("Number of vertices, expected %v, received %v", exp, c)
	}

	fmt.Println(g.ToString())
}
