package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

var memo [][]float64

func dist(p1, p2 Vertex) float64 {
	return math.Sqrt(((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}

func cost(i, j, k int, vertices []Vertex) float64 {
	p1 := vertices[i]
	p2 := vertices[j]
	p3 := vertices[k]
	return dist(p1, p2) + dist(p2, p3) + dist(p3, p1)
}

// https://www.cs.utexas.edu/users/djimenez/utsa/cs3343/lecture12.html
func triangulationCost(i, j int, vertices []Vertex) float64 {
	if memo[i][j] < math.MaxFloat64 {
		return memo[i][j]
	}
	min := math.MaxFloat64
	if i == j {
		memo[i][j] = 0.0
	} else {
		for k := i; k <= j-1; k++ {
			x := triangulationCost(i, k, vertices) + triangulationCost(k+1, j, vertices) + cost(i-1, k, j, vertices)
			if x < min {
				min = x
			}
		}
	}
	memo[i][j] = min
	return min
}

// Optimal cost for triangulation of a convex polygon with vertices v0...vn
// O(n^2) space (memo-table), O(n^3) time (does "n things on n^2 array elements")
func triangulation(vertices []Vertex) float64 {
	n := len(vertices) - 1
	memo = make([][]float64, n+1)
	for i := 0; i < n; i++ {
		memo[i] = make([]float64, n+1)
		for j := 0; j < n; j++ {
			memo[i][j] = math.MaxFloat64
		}
	}
	return triangulationCost(0, n, vertices)
}

func main() {
	// polygon = v0, v1, ..., vn
	v := []Vertex{{0, 0}, {1, 0}, {2, 1}, {1, 2}, {0, 2}}
	costs := triangulation(v)
	fmt.Println("Optimal triangulation costs:", costs)
}
