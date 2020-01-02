package graph

import (
	"github.com/golang-collections/collections/stack"
)

// Recursive depth-first search, with circle detection.
// If no circle is detected then true is returned and
// sorted contains the topological sort order; otherwise
// false is returned and sorted contains the detected
// circle.
func dfs(v *Vertex, sorted *stack.Stack) (acyclic bool) {
	if v.color == Black {
		// Already completely dealt with
		return true
	} else if v.color == Gray {
		// Graph is not a DAG -> cycle detected
		// Push the vertex which closes the circle
		sorted.Push(v)
		return false
	}

	v.color = Gray
	acyclic = true
	for _, e := range v.adjacency {
		// Recursively visit all adjacent vertices first ("depth first")
		acyclic = dfs(e.To, sorted)
		if !acyclic {
			break
		}

	}
	v.color = Black
	sorted.Push(v)
	return acyclic
}

// TopoSort returns a topologically sorted order of
// the vertices in the graph. The boolean acyclic is true if no
// circle has been detected, otherwise sorted contains
// the vertices which form a closed circle (and acyclic is false).
func TopoSort(g *Graph) (vertices []*Vertex, acyclic bool) {
	sorted := stack.New()

	for _, v := range g.Vertices {
		v.color = White
	}
	for _, v := range g.Vertices {
		if v.color != Black {
			acyclic = dfs(v, sorted)
			if !acyclic {
				break
			}
		}
	}
	vertices = make([]*Vertex, 0, sorted.Len())
	l := sorted.Len()
	for i := 0; i < l; i++ {
		vertices = append(vertices, sorted.Pop().(*Vertex))
	}
	return vertices, acyclic
}
