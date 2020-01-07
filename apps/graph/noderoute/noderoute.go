package main

import (
	"github.com/golang-collections/collections/stack"
	"github.com/till213/back-to-school/struct/graph/adjacency"
	"fmt"
	"os"
)

func dfs(g *graph.Graph, source, dest *graph.Vertex) bool {
	for _, v := range g.Vertices {
		v.Color = graph.Unvisited
	}
	stack := stack.New()
	stack.Push(source)
	source.Color = graph.Visited
	for stack.Len() > 0 {
		current := stack.Pop().(*graph.Vertex)
		for _, edge := range current.Adjacency {
			if edge.To.Color == graph.Unvisited {
				if edge.To != dest {
					stack.Push(edge.To)
					edge.To.Color = graph.Visited
				} else {
					return true
				}
			}
		}
	}
	return false
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

	exist := dfs(g, v1, v2)
	fmt.Printf("Route exists from %s to %s: %v\n", vertexName1, vertexName2, exist)
} 