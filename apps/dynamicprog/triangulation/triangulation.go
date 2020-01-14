package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

type Cost struct {
	// the triangulation cost
	cost float64
	// the third vertex k which implements the minimum cost
	minK int
}

// Returns the Euklidean distance between vertices p1 and p2.
func dist(v1, v2 Vertex) float64 {
	return math.Sqrt(((v1.x-v2.x)*(v1.x-v2.x) + (v1.y-v2.y)*(v1.y-v2.y)))
}

// The sum of the Euklidean distance between vertices i, j, and k
// is the cost function
func cost(i, j, k int, vertices []Vertex) float64 {
	p1 := vertices[i]
	p2 := vertices[j]
	p3 := vertices[k]
	return dist(p1, p2) + dist(p2, p3) + dist(p3, p1)
}

func printMemo(memo [][]Cost) {
	for i := 0; i < len(memo); i++ {
		for j := 0; j < len(memo[i]); j++ {
			fmt.Printf("%6.3f (%2d) ", memo[i][j].cost, memo[i][j].minK)
		}
		fmt.Println()
	}
}

// Optimal cost for triangulation of a convex polygon with vertices v0...vn
// O(n^2) space (memo-table), O(n^3) time (does "n things on n^2 array elements")
// Implementation according to Introduction to Algorithms.
// Also refer to: https://www.cs.utexas.edu/users/djimenez/utsa/cs3343/lecture12.html
//
// This implementation creates a somewhat smaller memo table (dimension #vertices-1),
// but from which it is not so straight forward to reconstruct the triangulation:
// the table indices do not correspond 1:1 to the indices into the vertices slice                     
func iterativeTriangulation1(vertices []Vertex) [][]Cost {
	n := len(vertices) - 1

	if n < 2 {
		// there must be at least 3 points to form a triangle
		return [][]Cost{{Cost{0.0, -1}}}
	}

	// n x n memo table, diagonal all 0s (implicitly set)
	memo := make([][]Cost, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]Cost, n)
		// initialise diagonal with 0 cost and "no triangulation"
		// -> invalid index -1
		memo[i][i] = Cost{0.0, -1}
	}

	for l := 2; l <= n; l++ {
		for i := 1; i <= n - l + 1; i++ {
			j := i + l - 1
			// index in memo starts at 0
			memo[i-1][j-1].cost = math.MaxFloat64
			memo[i-1][j-1].minK = -1
			for k := i; k <= j - 1; k++ {
				q := memo[i-1][k-1].cost + memo[k][j-1].cost + cost(i-1, k, j, vertices)
				if q < memo[i-1][j-1].cost {
					memo[i-1][j-1].cost = q
					memo[i-1][j-1].minK = k
				}
			}
		}
	}
	return memo
}

// https://www.geeksforgeeks.org/minimum-cost-polygon-triangulation/
// http://www.cs.utoronto.ca/~heap/Courses/270F02/A4/chains/node2.html 
func iterativeTriangulation2(vertices []Vertex) [][]Cost {

	n := len(vertices)

	// There must be at least 3 points to form a triangle 
	if n < 3 {
		return [][]Cost{{Cost{0.0, 0}}}
	}

	// table to store results of subproblems.  table[i][j] stores cost of 
	// triangulation of points from i to j.  The entry table[0][n-1] stores 
	// the final result. 
	memo := make([][]Cost, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]Cost, n)
	}

	// Fill table using the recursive cost formula. Note that the table 
	// is filled in diagonal fashion i.e., from diagonal elements to 
	// table[0][n-1] which is the result.
	
	// the gap defines the distance between the "fixed points" i and j
	for gap := 0; gap < n; gap++ {
		i := 0
		for j := gap; j < n; j++ {
			if j < i + 2 {
				// points and lines are not triangles and have hence
				// a convenient cost of 0.0
				memo[i][j] = Cost{0.0, -1}
			} else { 
				memo[i][j].cost = math.MaxFloat64;
				// between this "gap" we iterate a third point k 
				for k := i + 1; k < j; k++ {
					// and we choose k such that it minimises the cost,
					// which is the sum of triangulating the "left polygon",
					// the "right polygon" and the cost of the current
					// triangle i, k, j (a given triangle divides a convex polygon
					// into three parts: "left", "right" and the current triangle)
					val := memo[i][k].cost + memo[k][j].cost + cost(i, k, j, vertices)
					if (memo[i][j].cost > val) {
						memo[i][j].cost = val; 
						memo[i][j].minK = k; 
					} 
				} 
			} 
			i++
		}
	} 
	return memo;
}

func printTriangulation(i, j int, memo [][]Cost) {
	// Check if there is still a triangle left
	if memo[i][j].minK != -1 {
 		fmt.Printf("<%d, %d, %d>, ", i, memo[i][j].minK, j)
		printTriangulation(i, memo[i][j].minK, memo)
		printTriangulation(memo[i][j].minK, j, memo)
	}
}

func main() {
	// polygon = v0, v1, ..., vn
	v := []Vertex{{0, 0}, {1, 0}, {2, 1}, {1, 2}, {0, 2}}
	memo := iterativeTriangulation2(v)
	printMemo(memo)
	n := len(v)
	fmt.Println("Optimal triangulation costs:", memo[0][n-1].cost)
	printTriangulation(0, n-1, memo)
}
