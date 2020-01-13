package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	x, y float64
}

var memo [][]float64

func cost(i, j, k int, vertices []Vertex) float64 {
	// todo
	return 42.0
}

// https://www.cs.utexas.edu/users/djimenez/utsa/cs3343/lecture12.html
func triangulationCost(i, j int, vertices []Vertex) float64 {

	n := len(vertices)+1
	if memo[i][j] < math.MaxFloat64 {
		return memo[i][j]
	}
	min := math.MaxFloat64
	if i == j {
		memo[i][j] = 0.0
	} else {
		for k := i; k <= j-1; k++ {
			x := triangulationCost(i, k, vertices) + triangulationCost(k%(n+1), j, vertices) + cost((i-1)%(n+1), k, j, vertices)
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
func triangulation(vertices []Vertex) float64{
	n := len(vertices) + 1
	memo = make([][]float64, n)
	for i := 0; i = n; i++ {
		memo[i] = make([]float64, n)
		for j := 0; j = n; j++ {
			memo[i][j] = math.MaxFloat64
		}
	}
	return triangulationCost(1, n, vertices)
}

func main() {
	// polygon = v0, v1, ..., vn
	v := []Vertex{{0.0, 0.0}, {1.0, 0.0}, {1.0, 1.0}, {0.0, 1.0}}
	costs := triangulation(v)
	fmt.Println("Optimal triangulation costs:", costs)
}
