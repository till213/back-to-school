package graph

import "math/rand"

// NamedVertices contains example named vertices
var NamedVertices = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M"}

// DAG returns a weighted (random costs) directed acyclic graph.
func DAG() *Graph {
	g := New()

	rand.Seed(1)
	g.AddVertices(NamedVertices)
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
	return g
}

// WithCycle returns a graph with a simple cycle
func WithCycle() *Graph {
	g := New()

	rand.Seed(1)
	g.AddVertices([]string{"A", "B", "C"})
	g.AddEdge("A", "B", rand.Intn(20))
	g.AddEdge("B", "C", rand.Intn(20))
	g.AddEdge("C", "A", rand.Intn(20))
	return g
}
