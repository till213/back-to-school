package main

import (
	"fmt"
	"os"
	"strconv"
)

var h [][]int

// Possible 8 knight moves
const nofKnightMoves = 8

var dx = []int{2, 1, -1, -2, -2, -1, 1, 2}
var dy = []int{1, 2, 2, 1, -1, -2, -2, -1}

func vertialBorder(boardSize int) {
	fmt.Print("+")
	for i := 0; i < boardSize; i++ {
		fmt.Print("--+")
	}
	fmt.Println("")
}

func printBoard(boardSize int) {
	vertialBorder(boardSize)
	for j := 0; j < boardSize; j++ {
		fmt.Print("|")
		for i := 0; i < boardSize; i++ {
			fmt.Printf("%2d|", h[j][i])
		}
		fmt.Println(" ", j)
		vertialBorder(boardSize)
	}
	fmt.Print(" ")
	for j := 0; j < boardSize; j++ {
		fmt.Printf("%2d ", j)
	}
	fmt.Println("")
}

// Finds the first (any) solution to the Eight Queen problem
func try(i, x, y int, boardSize int) bool {
	k := 0
	done := false
	sizeSquare := boardSize * boardSize
	for !done && k < nofKnightMoves {
		done = false
		u := x + dx[k]
		v := y + dy[k]
		if 0 <= u && u < boardSize && 0 <= v && v < boardSize && h[v][u] == 0 {
			h[v][u] = i
			if i < sizeSquare {
				done = try(i+1, u, v, boardSize)
				if !done {
					h[v][u] = 0
				}
			} else {
				done = true
			}
		}
		k++
	}
	return done
}

func initBoard(boardSize int) {
	h = make([][]int, boardSize)
	for j := 0; j < boardSize; j++ {
		h[j] = make([]int, boardSize)
	}
}

// Starts a Knight's Tour from i, j, returns true
// on success (all boardSizexboardSize fields visited)
func knightsTour(i, j int, boardSize int) bool {
	initBoard(boardSize)
	h[j][i] = 1
	return try(2, i, j, boardSize)
}

// Usage: app i j n
func main() {
	var boardSize, i, j int
	var err error
	if len(os.Args) > 2 {
		i, err = strconv.Atoi(os.Args[1])
		if err == nil {
			j, err = strconv.Atoi(os.Args[2])
		}
	} else {
		i, j = 0, 0
	}
	if len(os.Args) > 3 {
		boardSize, err = strconv.Atoi(os.Args[3])
	} else {
		boardSize = 12
	}
	if err == nil {
		success := knightsTour(i, j, boardSize)
		if success {
			printBoard(boardSize)
		} else {
			fmt.Println("No solution found for starting point:", i, j, "on board size:", boardSize)
		}
	} else {
		fmt.Println("Parse error: ", err)
	}
}
