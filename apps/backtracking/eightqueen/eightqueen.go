package main

import "fmt"

const boardSize = 8

var x [boardSize]int
var a [boardSize]bool
var b, c [boardSize*2 - 1]bool

func try(i int, q bool) bool {
	var j int
	for !q && j != boardSize {

		if a[j] && b[i+j] && c[i-j+boardSize-1] {
			x[i] = j
			a[j] = false
			b[i+j] = false
			c[i-j+boardSize-1] = false
			if i < 7 {
				q = try(i+1, q)
				if !q {
					a[j] = true
					b[i+j] = true
					c[i-j+7] = true
				}
			} else {
				q = true
			}
		}
		j++

	}
	return q
}

func queens() {
	for i := 0; i < boardSize; i++ {
		a[i] = true
	}
	for i := 0; i < boardSize*2-1; i++ {
		b[i] = true
		c[i] = true
	}
	try(0, false)
	for i := 0; i < boardSize; i++ {
		fmt.Println("👸 at", i, x[i])
	}
}

func main() {
	queens()
}
