package main

import (
	"fmt"
)

var memo [][]int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int { 
    return min(min(a, b), c); 
} 

func initTable(m, n int) {
	memo = make([][]int, m+1)

	for i := 0; i <= m; i++ {
		memo[i] = make([]int, n+1)
	}
	// This initialisation is not strictly required,
	// as in Go ints default to 0 anyway
	for i := 1; i <= m; i++ {
		memo[i][0] = 0
	}
	for j := 0; j <= n; j++ {
		memo[0][j] = 0
	}
}

func printTable(X, Y string) {
	fmt.Println("--- Memo Table ---")
	fmt.Print("      ")
	for _, c := range Y {
		fmt.Printf("%c  ", c)
	}
	fmt.Println()
	for i := 0; i <= len(X); i++ {
		if i > 0 {
			fmt.Printf("%c ", rune(X[i-1]))
		} else {
			fmt.Print("  ")
		}
		for j := 0; j <= len(Y); j++ {
			fmt.Printf("%2d ", memo[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// Finds the edit distance given the operations Insert, Remove and Replace, all
// of which having equal costs.
func editDistance(str1, str2 string) int { 
	// Create a table to store results of subproblems ¨
	m := len(str1)
	n := len(str2)
    initTable(m, n)
  
    // Fill memo[][] in bottom up manner 
    for i := 0; i <= m; i++ { 
        for j := 0; j <= n; j++ { 
            if (i == 0) {
				// If first string is empty, only option is to 
                // insert all characters of second string 
				// Min. operations = j 
				memo[i][j] = j; 
			} else if (j == 0) {
				// If second string is empty, only option is to 
            	// remove all characters of second string 
				// Min. operations = i 
				memo[i][j] = i; 
			} else if str1[i - 1] == str2[j - 1] {
				// If last characters are same, ignore last char 
				// and recur for remaining string 
                memo[i][j] = memo[i - 1][j - 1]; 
			} else {
				// If the last character is different, consider all 
            	// possibilities and find the minimum 
                memo[i][j] = 1 + min3(memo[i][j - 1],      // Insert 
                                      memo[i - 1][j],      // Remove 
									  memo[i - 1][j - 1]); // Replace 
			}
        } 
    } 
    return memo[m][n]
} 
  
func main() {
	str1 := "Mamma mia"
	str2 := "Hotel Mamma"
	editDistance := editDistance(str1, str2)
	fmt.Printf(`Edit distance for "%s" and "%s": %d`, str1, str2, editDistance)
	fmt.Println()
	printTable(str1, str2)
}