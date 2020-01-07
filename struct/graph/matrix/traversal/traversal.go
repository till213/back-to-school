package main

import (
	"fmt"
	"container/list"
)

// Matrix represents the matrix graph
type Matrix [][]int

// Position defines the current position in the graph (matrix)
type Position struct {
	row, col int
}

// Direction indicates the movement (delta between positions)
type Direction struct {
	// We can move +/- 1 in each direction (up/down via rows, left/right via columns)
	row, col int
}

// MatrixGraph is a matrix-based graph, where one can travel up, down, right, left
// (considering the boundaries)
type MatrixGraph struct {
	matrix     Matrix
	rows, cols int
	p          Position
}

var directions = [...]Direction{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// Create creates a matrix graph with dimension [rows, cols].
func Create(rows, cols int) *MatrixGraph {
	mg := new(MatrixGraph)
	mg.matrix = make(Matrix, rows)
	for i := range mg.matrix {
		mg.matrix[i] = make([]int, cols)
	}

	mg.rows = rows
	mg.cols = cols

	i := 0
	for row := range mg.matrix {
		for col := range mg.matrix[row] {
			i++
			mg.matrix[row][col] = i
		}
	}
	return mg
}

func (mg *MatrixGraph) traverseDfs(to Position, visited map[Position]bool, f func(pos Position, mg *MatrixGraph)) {

	if visited[to] {
		return
	}
	visited[to] = true
	// For all neighbours
	for _, d := range directions {
		nextRow, nextCol := to.row+d.row, to.col+d.col
		if 0 <= nextRow && nextRow < mg.rows && 0 <= nextCol && nextCol < mg.cols {
			mg.traverseDfs(Position{nextRow, nextCol}, visited, f)
		}
	}
	f(to, mg)
}

func (mg *MatrixGraph) dfs(f func(pos Position, mg *MatrixGraph)) {
	if mg == nil {
		return
	}
	visited := make(map[Position]bool)
	for row := range mg.matrix {
		for col := range mg.matrix[row] {
			mg.traverseDfs(Position{row, col}, visited, f)
		}
	}
}

func (mg *MatrixGraph) traverseBfs(to Position, visited map[Position]bool, f func(pos Position, mg *MatrixGraph)) {
	queue := list.New()
	queue.PushBack(to)
	for queue.Len() > 0 {
		e := queue.Front()
		pos := e.Value.(Position)
		queue.Remove(e)
		if !visited[pos] {
			visited[pos] = true
			for _, d := range directions {
				nextRow, nextCol := to.row+d.row, to.col+d.col
				if 0 <= nextRow && nextRow < mg.rows && 0 <= nextCol && nextCol < mg.cols {
					queue.PushBack(Position{nextRow, nextCol})
				}
			}
			f(pos, mg)
		}
	}
}

func (mg *MatrixGraph) bfs(f func(pos Position, mg *MatrixGraph)) {
	if mg == nil {
		return
	}
	visited := make(map[Position]bool)
	for row := range mg.matrix {
		for col := range mg.matrix[row] {
			mg.traverseBfs(Position{row, col}, visited, f)
		}
	}
}

func visit(p Position, mg *MatrixGraph) {
	fmt.Printf("%d, ", mg.matrix[p.row][p.col])
}

func main() {
	mg := Create(4, 8)
	fmt.Println("DFS")
	mg.dfs(visit)
	fmt.Println()
	fmt.Println("BFS")
	mg.bfs(visit)
}
