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
	g.AddEdge("K", "J", rand.Intn(20))
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

// Directed returns a directed graph (with possible cycles)
func Directed(disconnected bool) *Graph {
	g := New()

	rand.Seed(1)
	g.AddVertices(NamedVertices)
	g.AddVertices(NamedVertices)
	g.AddEdges("A", []string{"B", "C"}, []int{rand.Intn(20), rand.Intn(20)})
	g.AddEdges("B", []string{"D"}, []int{rand.Intn(20)})
	g.AddEdges("C", []string{"E", "L"}, []int{rand.Intn(20), rand.Intn(20)})
	if !disconnected {
		g.AddEdges("D", []string{"F"}, []int{rand.Intn(20)})
		g.AddEdges("E", []string{"H"}, []int{rand.Intn(20)})
	} else {
		g.AddEdges("D", []string{}, []int{})
		g.AddEdges("E", []string{}, []int{})
	}	
	g.AddEdges("F", []string{"G", "I"}, []int{rand.Intn(20), rand.Intn(20)})
	g.AddEdges("G", []string{}, []int{})
	g.AddEdges("H", []string{"F"}, []int{rand.Intn(20)})
	g.AddEdges("I", []string{"K", "G"}, []int{rand.Intn(20), rand.Intn(20)})
	g.AddEdges("J", []string{"F", "H"}, []int{rand.Intn(20), rand.Intn(20)})
	g.AddEdges("K", []string{"J"}, []int{rand.Intn(20)})
	if !disconnected { 
		g.AddEdges("L", []string{"H"}, []int{rand.Intn(20)})
		g.AddEdges("M", []string{"B", "D", "G"}, []int{rand.Intn(20), rand.Intn(20), rand.Intn(20)})
	} else {
		g.AddEdges("L", []string{}, []int{})
		g.AddEdges("M", []string{"B", "D"}, []int{rand.Intn(20), rand.Intn(20)})
	}
	return g
}
