package graph

import (
	"container/list"
	"github.com/golang-collections/collections/stack"
)

// DFS implements a non-recursive depth-first search on graph g,
// starting at vertex v, calling f(v) on each
// visited vertex.
// O(E + V)
func DFS(g *Graph, v *Vertex, f func(v *Vertex)) {
	for _, v := range g.Vertices {
		v.Visited = false
	}
	stack := stack.New()
	stack.Push(v)
	v.Visited = true
	for stack.Len() > 0 {
		current := stack.Pop().(*Vertex)
		for _, edge := range current.Adjacency {
			if !edge.To.Visited {
				stack.Push(edge.To)
				edge.To.Visited = true
			}
		}
		f(current)
	}
}

func recursiveDFS(v *Vertex, f func(v *Vertex)) {
	f(v)
	for _, edge := range v.Adjacency {
		if !edge.To.Visited {
			recursiveDFS(edge.To, f)
			edge.To.Visited = true
		}
	}
}

// RecursiveDFS implements a recursive depth-first search on graph g,
// starting at vertex v, calling f(v) on each
// visited vertex.
// O(E + V)
func RecursiveDFS(g *Graph, v *Vertex, f func(v *Vertex)) {
	for _, v := range g.Vertices {
		v.Visited = false
	}
	recursiveDFS(v, f)
}

// BFS implements a non-recursive breadth-first search on graph g,
// starting at vertex v, calling f(v) on each
// visited vertex.
// O(E + V)
func BFS(g *Graph, v *Vertex, f func(v *Vertex)) {
	for _, v := range g.Vertices {
		v.Visited = false
	}
	list := list.New()
	list.PushFront(v)
	v.Visited = true
	for list.Len() > 0 {
		current := list.Back().Value.(*Vertex)
		list.Remove(list.Back())
		for _, edge := range current.Adjacency {
			if !edge.To.Visited {
				list.PushFront(edge.To)
				edge.To.Visited = true
			}
		}
		f(current)
	}
}
