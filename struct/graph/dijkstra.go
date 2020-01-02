package graph

import (
	"container/heap"
)

const maxDistance = 1<<31 - 1
const notInHeap = -1

// PathVertices represent (shortest) paths in the graph
type PathVertices []*Vertex

// initSingleSource returns all the vertices as a vertex slice.
// The vertex distances are all initialised to maxDistance, except
// for the start vertex, for which it is set to 0.
// The index - which is used to keep track of the position in the
// heap respectively the returned vertex slice - is updated
// according to uddateIndex (true), respective set to notInHeap (-1)
// when set to false.
func (g *Graph) initSingleSource(start *Vertex, updateIndex bool) PathVertices {
	pathVertices := make(PathVertices, len(g.Vertices))
	i := 0
	for _, v := range g.Vertices {
		if v != start {
			v.Distance = maxDistance
		} else {
			v.Distance = 0
		}
		v.Predecessor = nil
		if updateIndex {
			v.index = i
		} else {
			v.index = notInHeap
		}
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

// updateDistance (aka "decrease key") updates the distance for the given
// vertex, and re-establishes the underlying heap structure (in O(log n)).
func (v *PathVertices) updateDistance(vertex *Vertex, distance int) {
	vertex.Distance = distance
	heap.Fix(v, vertex.index)
}

// ClassicDijkstra calculates all the shortest paths, starting from
// start vertex to any other connected vertex. Costs must all
// be non-negative.
//
// The "classic" implementation is derived from an "Introduction
// to Algorithms" and builds the heap up front, with all
// graph vertices. That is, no "is element of" check is hence
// necessary (but that check can be easily implemented in O(1) time).
//
// Assuming that the implementation of Golang heap is a binary heap
// the runtime is hence O((V + E)log V).
// With a Fibonacci heap which supports decreaseKey in amortised O(1)
// the total runtime would only be O(V log V + E).
func (g *Graph) ClassicDijkstra(start *Vertex) {

	vertices := g.initSingleSource(start, true)

	// priority queue representing all the graph vertices to explore,
	// with the closest vertices first
	heap.Init(&vertices)

	// For all vertices O(V)
	for vertices.Len() > 0 {
		vertex := heap.Pop(&vertices).(*Vertex)
		// For all edges (of current vertex), O(E)
		for _, e := range vertex.Adjacency {
			distance := vertex.Distance + e.Cost
			if distance < e.To.Distance {
				// Relax vertex ("decrease key")
				vertices.updateDistance(e.To, distance)
				e.To.Predecessor = vertex
			}
		}
	}
}

// FrontierDijkstra calculates the shortest paths like the
// above ClassicDijkstra, except that the heap - the "frontier" of
// newly discovered vertices - is extended as we go. This
// requires an efficient "find" ("member of") implementation, but
// which is easily implemented in O(1) by setting the index flag from
// its initial notInHeap value ("not member of") to a valid index [0, V[
// ("member of").
func (g *Graph) FrontierDijkstra(start *Vertex) {

	vertices := g.initSingleSource(start, false)

	frontier := make(PathVertices, 1)
	start.index = 0
	frontier[0] = start

	// priority queue representing the frontier of vertices
	// to exolore, with the closest vertices first
	heap.Init(&frontier)

	// Each graph vertex is added exactly once to the frontier,
	// so O(V)
	for frontier.Len() > 0 {
		vertex := heap.Pop(&frontier).(*Vertex)
		// For all edges (of current vertex), O(E)
		for _, e := range vertex.Adjacency {
			distance := vertex.Distance + e.Cost
			if distance < e.To.Distance {
				e.To.Predecessor = vertex
				if e.To.index != notInHeap {
					// We have already visited this vertex,
					// so simply relax this vertex ("decrease key")
					vertices.updateDistance(e.To, distance)
				} else {
					// New vertex discovered in "uncharted territory",
					// so add it to our "frontier"
					e.To.Distance = distance
					frontier.Push(e.To)
				}
			}
		}
	}
}
