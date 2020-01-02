package graph

import (
	"container/heap"
)

// Prim calculates the minimum spanning tree.
func (g *Graph) Prim(root *Vertex) {

	vertices := g.initSingleSource(root, false)

	// priority queue representing all the graph vertices to explore,
	// with the closest vertices first
	heap.Init(&vertices)

	for vertices.Len() > 0 {
		vertex := heap.Pop(&vertices).(*Vertex)
		// For all edges (of current vertex), O(E)
		for _, e := range vertex.adjacency {
			if e.To.index != notInHeap && e.Cost < e.To.distance {
				e.To.predecessor = vertex
				vertices.updatedistance(e.To, e.Cost)
			}
		}
	}
}
