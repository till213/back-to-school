package main

import (
	"fmt"
)

const n = 5

type matrix [5][5]bool

// Boolean matrix multiplication c = a * b
func bmm(a, b matrix, c *matrix) {
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

// ORs the given matrix a with I (identity matrix)
func orIdentity(a *matrix) {
	for i := 0; i < n; i++ {
		a[i][i] = true
	}
}

// Defines an example Adjacency matrix of a simple, directed graph E
func exampleAdjacency(A *matrix) {
	A[0][1] = true
	A[1][2] = true
	A[2][3] = true
	A[2][4] = true
	A[3][1] = true
	A[4][3] = true
}

func printMatrix(A *matrix) {
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

func main() {
	var A matrix

	exampleAdjacency(&A)
	fmt.Println("Adjacency A")
	printMatrix(&A)

	orIdentity(&A)
	fmt.Println("A'")
	printMatrix(&A)

	// A'^2
	bmm(A, A, &A)
	fmt.Println("A'^2")
	printMatrix(&A)

	// A'^4
	bmm(A, A, &A)

	// C = A'^(n-1)
	fmt.Println("C")
	printMatrix(&A)
}
