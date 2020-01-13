package main

import (
	"fmt"
)

const top = '⬆'
const topLeft = '↖'
const left = '⬅'

var b [][]rune
var c [][]int

func initTables(m, n int) {
	b = make([][]rune, m+1)
	c = make([][]int, m+1)

	for i := 0; i <= m; i++ {
		b[i] = make([]rune, n+1)
		c[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		c[i][0] = 0
	}
	for j := 0; j <= n; j++ {
		c[0][j] = 0
	}
}

func lcs(X, Y string) {

	m := len(X)
	n := len(Y)
	initTables(m, n)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// string indexing starts at 0
			if X[i-1] == Y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
				b[i][j] = topLeft
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
				b[i][j] = top
			} else {
				c[i][j] = c[i][j-1]
				b[i][j] = left
			}
		}
	}
}

func printLCS(X string, i, j int) {
	if i == 0 || j == 0 {
		return
	}
	if b[i][j] == topLeft {
		printLCS(X, i-1, j-1)
		fmt.Printf("%c, ", rune(X[i-1]))
	} else if b[i][j] == top {
		printLCS(X, i-1, j)
	} else {
		printLCS(X, i, j-1)
	}
}

func printTables(X, Y string) {
	fmt.Println("--- B, C Tables ---")
	fmt.Print("      ")
	for _, c := range Y {
		fmt.Printf("%c   ", c)
	}
	fmt.Println()
	for i := 0; i <= len(X); i++ {
		if i > 0 { 
		  fmt.Printf("%c ", rune(X[i-1]))
		} else {
			fmt.Print("  ")
		}
		for j := 0; j <= len(Y); j++ {
			fmt.Printf("%d %c ", c[i][j], b[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	X := "ABCBDAB"
	Y := "BDCABA"

	lcs(X, Y)
	fmt.Print("Longest Common Subsequence (LCS): ")
	printLCS(X, len(X), len(Y))
	fmt.Println()
	printTables(X, Y)
}
