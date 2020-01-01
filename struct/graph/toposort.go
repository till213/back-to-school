package graph

import (
	"github.com/golang-collections/collections/stack"
)

// Recursive depth-first serach
func dfs(v *Vertex, sorted *stack.Stack) {
	v.Visited = true
	for _, e := range v.Adjacency {
		// Recursively visit all adjacent vertices first ("depth first")
		if !e.To.Visited {
			dfs(e.To, sorted)
		}
	}
	sorted.Push(v)
}

// TopoSort returns a topologically sorted order of
// the vertices in the graph. Nil if a cycle exists.
// PRE: the visited flag is false for all vertices in g
func TopoSort(g *Graph) []*Vertex {
	sorted := stack.New()
	for _, v := range g.Vertices {
		if !v.Visited {
			dfs(v, sorted)
		}
	}
	vertices := make([]*Vertex, 0, sorted.Len())
	l := sorted.Len()
	for i := 0; i < l; i++ {
		vertices = append(vertices, sorted.Pop().(*Vertex))
	}
	return vertices
}
