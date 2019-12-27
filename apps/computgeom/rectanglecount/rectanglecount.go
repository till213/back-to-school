package main

import "fmt"

type point struct {
	x, y int
}

// Stores the y-coordinates of two points having the same x-coordinate
// ("vertical line, parallel to y-axis")
type pair struct {
	y1, y2 int
}

// countRectangles counts the number of rectangles, aligned with x/y axis,
// that can be formed with the given points.
func countRectangles(points []point) int {
	var count int

	ylines := make(map[pair]int)
	for _, p1 := range points {
		for _, p2 := range points {
			if p1.y > p2.y && p1.x == p2.x {
				p := pair{p1.y, p2.y}
				// The initial count will be 0 for ...
				count += ylines[p]
				// .. as long as we have only one common y-line, for whic
				// we increase the y-line count. For each additional y-line
				// the count will then be:
				// 0 + 1 = 1, 1 + 2 = 3, 3 + 3 = 6, 6 + 4 = 10, ...
				ylines[p]++
			}
		}
	}
	return count
}

func main() {
	points := []point{{0, 0}, {0, 10}, {10, 0}, {10, 10}, {0, 20}, {10, 20}, {20, 0}, {20, 10},  {20, 20}}
	count := countRectangles(points)
	fmt.Println("Number of rectangles", count)
}
