package main

import (
	"container/list"
	"github.com/till213/back-to-school/struct/graph"
	"fmt"
	"os"
)

func concurrentBfsWalk(v *graph.Vertex, color graph.Color, ch chan bool) {
	list := list.New()
	list.PushFront(v)
	v.Color = color
	for list.Len() > 0 {
		current := list.Back().Value.(*graph.Vertex)
		
		list.Remove(list.Back())
		for _, edge := range current.Adjacency {
			if edge.To.Color == graph.Unvisited {
				list.PushFront(edge.To)
				edge.To.Color = color
			} else if edge.To.Color != color {
				// vertex has already been visited by
				// other BFS walk -> route exists
				ch <- true
			}
		}
	}
	ch <- false
}

func concurrentBfs(v *graph.Vertex, color graph.Color, ch chan bool) {
	concurrentBfsWalk(v, color, ch)
	close(ch)
}

// Returns whether there are common reachable vertices between v1 and v2
// in the directed graph g.
func hasCommonReachableVertex(g *graph.Graph, v1, v2 *graph.Vertex) bool {
	var exists1, exists2 bool
	for _, v := range g.Vertices {
		v.Color = graph.Unvisited
	}

	ch1 := make(chan bool)
	ch2 := make(chan bool)
	go concurrentBfs(v1, graph.Blue, ch1)
	go concurrentBfs(v2, graph.Red, ch2)

	exists1 = true
	exists2 = true
	for {
		select {
		case exists1 = <-ch1:
			if exists1 {
				return true
			} else if !exists2 {
				return false
			}
		case exists2 = <-ch2:
			if exists2 {
				return true
			} else if !exists1 {
				return false
			}
		default:
			// No decision taken yet
		}
	}
	return exists1 && exists2
}

func main() {
	var vertexName1, vertexName2 string
	if len(os.Args) > 2 {
		vertexName1 = os.Args[1]
		vertexName2 = os.Args[2]
	} else {
		vertexName1 = "A"
		vertexName2 = "K"
	}

	g := graph.Directed(true)
	v1 := g.Vertex(vertexName1)
	v2 := g.Vertex(vertexName2)

	exist := hasCommonReachableVertex(g, v1, v2)
	fmt.Printf("Common reachable vertex exists for %s and %s: %v\n", vertexName1, vertexName2, exist)
} 