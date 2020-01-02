package graph

import (
	"container/list"
	"github.com/golang-collections/collections/stack"
)

func recursiveDFS(v *Vertex, f func(v *Vertex)) {
	f(v)
	v.color = Gray
	for _, edge := range v.adjacency {
		if edge.To.color == White {
			recursiveDFS(edge.To, f)
			edge.To.color = Visited
		}
	}
	v.color = Black
}

// RecursiveDFS implements a recursive depth-first search on graph g,
// starting at vertex v, calling f(v) on each visited vertex.
// O(E + V)
func RecursiveDFS(g *Graph, v *Vertex, f func(v *Vertex)) {
	for _, v := range g.Vertices {
		v.color = White
	}
	recursiveDFS(v, f)
}

// DFS implements a non-recursive depth-first search on graph g,
// starting at vertex v, calling f(v) on each visited vertex.
// O(E + V)
func DFS(g *Graph, v *Vertex, f func(v *Vertex)) {
	for _, v := range g.Vertices {
		v.color = Unvisited
	}
	stack := stack.New()
	stack.Push(v)
	v.color = Visited
	for stack.Len() > 0 {
		current := stack.Pop().(*Vertex)
		for _, edge := range current.adjacency {
			if edge.To.color == Unvisited {
				stack.Push(edge.To)
				edge.To.color = Visited
			}
		}
		f(current)
	}
}

// BFS implements a non-recursive breadth-first search on graph g,
// starting at vertex v, calling f(v) on each visited vertex.
// O(E + V)
func BFS(g *Graph, v *Vertex, f func(v *Vertex)) {
	for _, v := range g.Vertices {
		v.color = Unvisited
	}
	list := list.New()
	list.PushFront(v)
	v.color = Visited
	for list.Len() > 0 {
		current := list.Back().Value.(*Vertex)
		list.Remove(list.Back())
		for _, edge := range current.adjacency {
			if edge.To.color == Unvisited {
				list.PushFront(edge.To)
				edge.To.color = Visited
			}
		}
		f(current)
	}
}
