package main

import (
	"fmt"
	"os"
	"strconv"
)

var x []int
var a []bool
var b, c []bool
var numSolution int

func vertialBorder(boardSize int) {
	fmt.Print("+")
	for i := 0; i < boardSize; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

// Prints the square (boardSizexboardSize) chessboard
// in beautiful ASCII + Emoji art.
func printBoard(numSolution, boardSize int) {
	fmt.Println("--- Solution:", numSolution, "---")
	vertialBorder(boardSize)
	for j := 0; j < boardSize; j++ {
		fmt.Print("|")
		for i := 0; i < boardSize; i++ {
			if x[i] == j {
				fmt.Print("👸")
			} else {
				black := i%2 == 0
				if j%2 == 0 {
					black = !black
				}
				if black {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println("|", j)
	}
	vertialBorder(boardSize)
	fmt.Print(" ")
	for j := 0; j < boardSize; j++ {
		fmt.Print(j)
	}
	fmt.Println("")
}

// Finds the first (any) solution to the Eight Queen problem
func try(i int, success bool, boardSize int) bool {
	var j int
	for !success && j != boardSize {

		if a[j] && b[i+j] && c[i-j+boardSize-1] {
			x[i] = j
			a[j] = false
			b[i+j] = false
			c[i-j+boardSize-1] = false
			if i < boardSize-1 {
				success = try(i+1, success, boardSize)
				if !success {
					a[j] = true
					b[i+j] = true
					c[i-j+boardSize-1] = true
				}
			} else {
				success = true
			}
		}
		j++

	}
	return success
}

func tryAll(i, boardSize int) {
	for j := 0; j < boardSize; j++ {
		if a[j] && b[i+j] && c[i-j+boardSize-1] {
			x[i] = j
			a[j] = false
			b[i+j] = false
			c[i-j+boardSize-1] = false
			if i < boardSize-1 {
				tryAll(i+1, boardSize)
			} else {
				numSolution++
				printBoard(numSolution, boardSize)
			}
			a[j] = true
			b[i+j] = true
			c[i-j+boardSize-1] = true
		}
	}
}

func initBoard(boardSize int) {
	for i := 0; i < boardSize; i++ {
		a[i] = true
	}
	for i := 0; i < boardSize*2-1; i++ {
		b[i] = true
		c[i] = true
	}
	numSolution = 0
}

func queens(boardSize int) {

	x = make([]int, boardSize)
	a = make([]bool, boardSize)
	b = make([]bool, boardSize*2-1)
	c = make([]bool, boardSize*2-1)
	initBoard(boardSize)
	// Single solution
	try(0, false, boardSize)

	for i := 0; i < boardSize; i++ {
		fmt.Println("👸 at", i, x[i])
	}
	printBoard(1, boardSize)

	// All solutions
	initBoard(boardSize)
	tryAll(0, boardSize)
}

func main() {
	var n int
	var err error
	if len(os.Args) > 1 {
		n, err = strconv.Atoi(os.Args[1])
	} else {
		n = 8
	}
	if err == nil {
		queens(n)
	} else {
		fmt.Println("Parse error: ", err)
	}
}
