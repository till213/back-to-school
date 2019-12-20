package main

import (
	"fmt"
	"math"
)

const n = 5

// Boolean matrix
type bmatrix [n][n]bool

// Float matrix
type fmatrix [n][n]float64

// Boolean matrix multiplication c = a * b
func bmm(a, b bmatrix, c *bmatrix) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c[i][j] = false
			// Optimised variant
			k := 1
			for !c[i][j] && k < n {
				c[i][j] = a[i][k] && b[k][j]
				k++
			}
		}
	}
}

func warshall(a *bmatrix) {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// This assignment mixes values of the old and new matrix
				a[i][j] = a[i][j] || a[i][k] && a[k][j]
			}
		}
	}
}

// Calculates the minimum distance (Floyd's algorithm)
func floyd(a *fmatrix) {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// This assignment mixes values of the old and new matrix
				a[i][j] = math.Min(a[i][j], a[i][k]+a[k][j])
			}
		}
	}
}

// ORs the given matrix a with I (identity matrix)
func orIdentity(a *bmatrix) {
	for i := 0; i < n; i++ {
		a[i][i] = true
	}
}

// Defines an example Adjacency matrix of a simple, directed graph E
func exampleAdjacency(A *bmatrix) {
	A[0][1] = true
	A[1][2] = true
	A[2][3] = true
	A[2][4] = true
	A[3][1] = true
	A[4][3] = true
}

// Defines an example Distance matrix of a simple, directed graph E
func exampleDistance(A *fmatrix) {

	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			A[i][j] = math.MaxFloat64
		}
	}

	A[0][1] = 2.0
	A[1][2] = 3.0
	A[2][3] = 4.0
	A[2][4] = 0.5
	A[3][1] = 5.0
	A[4][3] = 2.5
}

// Prints the Boolean matrix A
func printBMatrix(A *bmatrix) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] {
				fmt.Printf("1 ")
			} else {
				fmt.Printf("0 ")
			}
		}
		fmt.Println("")
	}
}

// Prints the Float matrix A
func printFMatrix(A *fmatrix) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			f := A[i][j];
			if f < math.MaxFloat64 {
				fmt.Printf("%6.2f ", A[i][j])
			} else {
				fmt.Printf("______ ")
			}
		}
		fmt.Println("")
	}
}

func main() {
	var A bmatrix

	exampleAdjacency(&A)
	fmt.Println("Adjacency A")
	printBMatrix(&A)

	orIdentity(&A)
	fmt.Println("A'")
	printBMatrix(&A)

	// A'^2
	bmm(A, A, &A)
	fmt.Println("A'^2")
	printBMatrix(&A)

	// A'^4
	bmm(A, A, &A)

	// C = A'^(n-1)
	fmt.Println("C")
	printBMatrix(&A)

	var B bmatrix

	exampleAdjacency(&B)
	fmt.Println("Adjacency B")
	printBMatrix(&B)

	warshall(&B)
	fmt.Println("Warshall B")
	printBMatrix(&B)

	var C fmatrix
	exampleDistance(&C)
	fmt.Println("Distance C")
	printFMatrix(&C)

	floyd(&C)
	fmt.Println("Floyd's C")
	printFMatrix(&C)
}
