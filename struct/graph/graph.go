package graph

import "strings"

import "fmt"

// UnweightedCost represents the cost of an unweighted edge
const UnweightedCost = 0

// Vertex represents a vertex in a graph
type Vertex struct {
	Name      string
	Visited   bool
	Adjacency []*Edge
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
// returns it.
func (g *Graph) AddVertex(name string) *Vertex {
	v := new(Vertex)
	v.Name = name
	v.Adjacency = make([]*Edge, 0, 8)
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

// ToString returns a textual represenation of this graph.
func (g *Graph) ToString() string {
	var sb strings.Builder

	for k, v := range g.Vertices {
		fmt.Fprintf(&sb, "Vertex %s, visited %v\n", k, v.Visited)
		for _, v := range v.Adjacency {
			fmt.Fprintf(&sb, "\t-> %s cost: %d\n", v.To.Name, v.Cost)
		}
	}
	return sb.String()
}

// VertexCount returns the number of vertices in this graph
func (g *Graph) VertexCount() int {
	return len(g.Vertices)
}
