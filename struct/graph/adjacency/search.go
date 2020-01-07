package graph

import (
	"container/list"
	"github.com/golang-collections/collections/stack"
)

func recursiveDFS(v *Vertex, f func(v *Vertex)) {
	f(v)
	v.Color = Gray
	for _, edge := range v.Adjacency {
		if edge.To.Color == White {
			recursiveDFS(edge.To, f)
			edge.To.Color = Visited
		}
	}
	v.Color = Black
}

// RecursiveDFS implements a recursive depth-first search on graph g,
// starting at vertex v, calling f(v) on each visited vertex.
// O(E + V)
func RecursiveDFS(g *Graph, v *Vertex, f func(v *Vertex)) {
	for _, v := range g.Vertices {
		v.Color = White
	}
	recursiveDFS(v, f)
}

// DFS implements a non-recursive depth-first search on graph g,
// starting at vertex v, calling f(v) on each visited vertex.
// O(E + V)
func DFS(g *Graph, v *Vertex, f func(v *Vertex)) {
	for _, v := range g.Vertices {
		v.Color = Unvisited
	}
	stack := stack.New()
	stack.Push(v)
	v.Color = Visited
	for stack.Len() > 0 {
		current := stack.Pop().(*Vertex)
		for _, edge := range current.Adjacency {
			if edge.To.Color == Unvisited {
				stack.Push(edge.To)
				edge.To.Color = Visited
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
		v.Color = Unvisited
	}
	list := list.New()
	list.PushFront(v)
	v.Color = Visited
	for list.Len() > 0 {
		current := list.Back().Value.(*Vertex)
		list.Remove(list.Back())
		for _, edge := range current.Adjacency {
			if edge.To.Color == Unvisited {
				list.PushFront(edge.To)
				edge.To.Color = Visited
			}
		}
		f(current)
	}
}
