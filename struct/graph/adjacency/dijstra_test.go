package graph

import (
	"fmt"
	"testing"
)

func TestClassicDijkstra(t *testing.T) {
	g := DAG()
	start := g.Vertex("C")
	g.ClassicDijkstra(start)

	fmt.Println("---- Classic Dijkstra's Shortest Paths ----")
	fmt.Println("Starting vertex:", start.Name)
	fmt.Println(g.ToString())
}

func TestFrontierDijkstra(t *testing.T) {
	g := DAG()
	start := g.Vertex("C")
	g.FrontierDijkstra(start)

	fmt.Println("---- Frontier Dijkstra's Shortest Paths ----")
	fmt.Println("Starting vertex:", start.Name)
	fmt.Println(g.ToString())
}

func BenchmarkClassicDijkstra(b *testing.B) {
	g := DAG()
	start := g.Vertex("C")
	for i := 0; i < b.N; i++ {
		g.ClassicDijkstra(start)
	}
}

func BenchmarkFrontierDijkstra(b *testing.B) {
	g := DAG()
	start := g.Vertex("C")
	for i := 0; i < b.N; i++ {
		g.FrontierDijkstra(start)
	}
}
