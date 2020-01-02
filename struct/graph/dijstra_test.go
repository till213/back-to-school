package graph

import (
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T) {
	g := DAG()

	start := g.Vertex("C")
	g.Dijkstra(start)

	fmt.Println("---- Dijkstra's Shortest Paths ----")
	fmt.Println("Starting vertex:", start)
	fmt.Println(g.ToString())
}
