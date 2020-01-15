package main

import (
	"fmt"
)

// Note: the longest common substring is related to the longest common subsequence (LCS).
//       However the substring must be connected, while a subsequence may have "gaps":
// Exampe: X = ABCBDAB Y = BDCABA
//         -> LCS: BCBA
//         -> Longest common substring: AB

type StringSet map[string]bool

var c [][]int

func initTable(m, n int) {
	c = make([][]int, m+1)

	for i := 0; i <= m; i++ {
		c[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		c[i][0] = 0
	}
	for j := 0; j <= n; j++ {
		c[0][j] = 0
	}
}

// The following tricks can be used to reduce the memory usage of an implementation:
// * Keep only the last and current row of the DP table to save memory  O(min(m, n)) instead of O(mn)
// * Store only non-zero values in the rows. This can be done using hash tables instead of arrays.
//   This is useful for large alphabets.
// https://en.wikipedia.org/wiki/Longest_common_substring_problem
func lcSubstring(X, Y string) StringSet {
	var ret StringSet

	m := len(X)
	n := len(Y)
	max := 0
	initTable(m, n)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// string indexing starts at 0
			if X[i-1] == Y[j-1] {
				if i == 1 || j == 1 {
					c[i][j] = 1
				} else {
					c[i][j] = c[i-1][j-1] + 1
				}

				if c[i][j] > max {
					// We have found a (new) substring which is
					// longer than the previous max length
					max = c[i][j]
					// Possible optimisation: we could only store the index i
					// of the last character. The strings can then be reconstructed
					// given their (max) length, as S[(ret[i]-max)..(ret[i])]
					str := X[i-max : i]
					// This is the single solution (for now), so
					// create a new set
					ret = make(StringSet)
					ret[str] = true
				} else if c[i][j] == max {
					// We have found another substring of equal
					// length, so add it to the existing set
					str := X[i-max : i]
					ret[str] = true
				}
			} else {
				// In contrast to the LCS problem we immediatelly set
				// the length to 0 as soon as the character at i, j
				// does not match anymore
				c[i][j] = 0
			}
		}
	}
	return ret
}

func printTable(X, Y string) {
	fmt.Println("--- C Table ---")
	fmt.Print("    ")
	for _, c := range Y {
		fmt.Printf("%c ", c)
	}
	fmt.Println()
	for i := 0; i <= len(X); i++ {
		if i > 0 {
			fmt.Printf("%c ", rune(X[i-1]))
		} else {
			fmt.Print("  ")
		}
		for j := 0; j <= len(Y); j++ {
			fmt.Printf("%d ", c[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	X := "ABCBDAB"
	Y := "BDCABA"

	ret := lcSubstring(X, Y)
	fmt.Print("Longest Common Substrings: ")
	for k := range ret {
		fmt.Printf("%s, ", k)
	}
	fmt.Println()
	printTable(X, Y)

	X = "ABCBDABABCDE"
	Y = "BDCABABCDE"

	ret = lcSubstring(X, Y)
	fmt.Print("Longest Common Substrings: ")
	for k := range ret {
		fmt.Printf("%s, ", k)
	}
	fmt.Println()
	printTable(X, Y)
}
