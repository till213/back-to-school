package graph

import (
	"container/heap"
)

const maxDistance = 1<<31 - 1

// PathVertices represent (shortest) paths in the graph
type PathVertices []*Vertex

// initSingleSource returns all the vertices as a path vertex slice.
// The vertex distances are all initialised to maxDistance, except
// for the start vertex, for which it is set to 0
func (g *Graph) initSingleSource(start *Vertex) PathVertices {
	pathVertices := make(PathVertices, len(g.Vertices))
	i := 0
	for _, v := range g.Vertices {
		if v != start {
			v.Distance = maxDistance
		} else {
			v.Distance = 0
		}
		v.Predecessor = nil
		v.index = i
		pathVertices[i] = v
		i++
	}
	return pathVertices
}

// Sort interface

// Len returns the count of all vertices.
func (v PathVertices) Len() int {
	return len(v)
}

// Less returns true if element with index i is less
// than element with index j, by comparing the
// vertex distances.
func (v PathVertices) Less(i, j int) bool {
	return v[i].Distance < v[j].Distance
}

// Swap swaps element i and j.
func (v PathVertices) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
	v[i].index = i
	v[j].index = j
}

// Heap interface

// Push appends x, which must be of type pointer to a vertex (*Vertex).
func (v *PathVertices) Push(x interface{}) {
	n := len(*v)
	item := x.(*Vertex)
	item.index = n
	*v = append(*v, item)
}

// Pop removes and returns the last vertex
func (v *PathVertices) Pop() interface{} {
	old := *v
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*v = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (v *PathVertices) updateDistance(vertex *Vertex, distance int) {
	vertex.Distance = distance
	heap.Fix(v, vertex.index)
}

// Dijkstra calculates all the shortest paths, starting from
// start vertex to any other connected vertex. Costs must all
// be non-negative.
func (g *Graph) Dijkstra(start *Vertex) {

	vertices := g.initSingleSource(start)

	// priority queue representing the vertices to exolore,
	// with the closest vertices first
	heap.Init(&vertices)

	// For all vertices O(V)
	for vertices.Len() > 0 {
		vertex := heap.Pop(&vertices).(*Vertex)
		// For all edges (of current vertex)
		for _, e := range vertex.Adjacency {
			distance := vertex.Distance + e.Cost
			if distance < e.To.Distance {
				// Relax vertex
				vertices.updateDistance(e.To, distance)
				e.To.Predecessor = vertex
			}
		}
	}
}
