package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// percentage
const blockProbability = 20
const width = 64
const height = 32
const block = '🏔'
const empty = ' '
const futile = '🔴'

var grid [][]rune

func initGrid(width, height int) {
	grid = make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
	}

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if rand.Intn(100) <= blockProbability {
				grid[h][w] = block
			} else {
				grid[h][w] = empty
			}
		}
	}
}

func findPath(x, y int) bool {
	if x == width || y == height || grid[y][x] == block {
		return false
	}
	if grid[y][x] == futile {
		return false
	}
	if x == width-1 && y == height-1 {
		// Reached the target
		grid[y][x] = '🍾'
		return true
	}
	success := false
	move := 0
	for !success && move < 3 {
		switch move {
		case 0:
			grid[y][x] = '→' //
			success = findPath(x+1, y)
		case 1:
			grid[y][x] = '⇲'
			success = findPath(x+1, y+1)
		case 2:
			grid[y][x] = '↓'
			success = findPath(x, y+1)
		default:
			// all moves tried, we are blocked -> backtrack
			success = false
			grid[y][x] = empty
		}
		move++
	}
	if !success {
		// Memoizdation: we mark this field as "no way"
		grid[y][x] = futile
	}
	return success
}

func printGrid() {
	var sb strings.Builder
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			sb.WriteRune(grid[h][w])
		}
		sb.WriteRune('\n')
	}
	fmt.Println(sb.String())
}

func main() {
	initGrid(width, height)
	found := findPath(0, 0)
	printGrid()
	fmt.Println("Path found:", found)
}
