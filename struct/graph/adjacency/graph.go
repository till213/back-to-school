package graph

import "strings"

import "fmt"

// UnweightedCost represents the cost of an unweighted edge
const UnweightedCost = 0

// Color defines the state ("colour") of a given vertex, e.g. black, grey, white
// or simply visited, unvisited, depending on the use case.
type Color int

// The vertex colours.
const (
	Black Color = iota
	Gray
	White
	Red
	Blue
	Visited
	Unvisited
)

// Vertex represents a vertex in a graph
type Vertex struct {
	Name        string
	Color       Color
	Adjacency   []*Edge
	distance    int
	predecessor *Vertex
	// Index for heap structure: the index is needed by update
	// and is maintained by the heap.Interface methods.
	index int
}

// Edge represents a directed connection between a
// vertex to another To vertex, associated
// with Cost.
type Edge struct {
	To   *Vertex
	Cost int
}

// Graph represents a directed graph with
// weighted (cost) edges
type Graph struct {
	Vertices map[string]*Vertex
}

// New creates and initialises a new empty graph.
func New() *Graph {
	g := new(Graph)
	g.Vertices = make(map[string]*Vertex)
	return g
}

// AddVertex creates and adds a named vertex, and
// returns it. Color is set to unvisited.
func (g *Graph) AddVertex(name string) *Vertex {
	v := new(Vertex)
	v.Name = name
	v.Adjacency = make([]*Edge, 0, 8)
	v.Color = Unvisited
	g.Vertices[name] = v
	return v
}

// AddVertices creates and adds named vertices.
func (g *Graph) AddVertices(names []string) {
	for _, n := range names {
		g.AddVertex(n)
	}
}

// Vertex returns the named vertex. If not exist
// then nil is returned
func (g *Graph) Vertex(name string) *Vertex {
	return g.Vertices[name]
}

// AddEdge adds an edge with cost to this vertex,
// connecting with the to vertex.
func (v *Vertex) AddEdge(to *Vertex, cost int) {
	e := new(Edge)
	e.Cost = cost
	e.To = to
	v.Adjacency = append(v.Adjacency, e)
}

// AddUnweightedEdge adds an unweighted edge (with cost=0) to this vertex,
// connecting with the to vertex.
func (v *Vertex) AddUnweightedEdge(to *Vertex) {
	v.AddEdge(to, UnweightedCost)
}

// AddEdge adds an edge with cost, connecting named
// vertices fromName and toName. No edge is added if either
// vertex does not exist.
func (g *Graph) AddEdge(fromName, toName string, cost int) {
	from, ok := g.Vertices[fromName]
	if ok {
		to, ok := g.Vertices[toName]
		if ok {
			from.AddEdge(to, cost)
		}
	}
}

// AddEdges adds multiple edge with costs, connecting named
// vertices fromName and toName. No edge is added if either
// vertex does not exist.
func (g *Graph) AddEdges(fromName string, toNames []string, costs []int) {
	for i, toName := range toNames {
		cost := costs[i]
		g.AddEdge(fromName, toName, cost)
	}
}

// AddUnweightedEdge adds an unweighted edge (with cost=0),
// connecting named vertices fromName and toName. No edge is added if
// either vertex does not exist.
func (g *Graph) AddUnweightedEdge(fromName, toName string) {
	g.AddEdge(fromName, toName, UnweightedCost)
}

// ToString adds a textual represenation of this vertex to sb.
func (v *Vertex) ToString(sb *strings.Builder) {
	fmt.Fprintf(sb, "-- Vertex %s --\n", v.Name)
	fmt.Fprintf(sb, "\tColor %v\n", v.Color)
	fmt.Fprintf(sb, "\tdistance %v\n", v.distance)
	fmt.Fprintf(sb, "\tpredecessors:\n\t")
	for p := v.predecessor; p != nil; p = p.predecessor {
		fmt.Fprintf(sb, "-> %s (%d) ", p.Name, p.distance)
	}
	fmt.Fprintf(sb, "\nAdjacency\n")
	for _, v := range v.Adjacency {
		fmt.Fprintf(sb, "\t-> %s cost: %d\n", v.To.Name, v.Cost)
	}
}

// ToString returns a textual represenation of this graph.
func (g *Graph) ToString() string {
	var sb strings.Builder
	for _, v := range g.Vertices {
		v.ToString(&sb)
	}
	return sb.String()
}

// VertexCount returns the number of vertices in this graph
func (g *Graph) VertexCount() int {
	return len(g.Vertices)
}
