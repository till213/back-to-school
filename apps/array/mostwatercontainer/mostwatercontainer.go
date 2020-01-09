
package main

import (
	"fmt"
)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func area(height []int, left, right int) int {
    h := min(height[left], height[right])
    w  := right - left
    return h * w
}

// Idea: left, right indices, increase/decrease them in a "greedy manner"
//       compare with max
func maxArea(height []int) int {   
    left := 0
	right := len(height)-1
	maxArea := area(height, left, right)
    for left < right-1  {
		// move the shorter height inwards, in the hope
		// that we (possibly) find an even taller height
        if height[left] < height[right] {
            left++
        } else {
            right--
		}
		area := area(height, left, right)
		maxArea = max(area, maxArea)
    }
    return maxArea
}

func main() {
	A := []int{1,3,2,5,25,24,5}
	area := maxArea(A)
	fmt.Println("Max area:", area)
}