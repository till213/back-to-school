package main

import "fmt"

type state int

const (
	unknown state = iota
	good
	bad
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	memo := make([]state, len(nums))
	for i := range memo {
		memo[i] = unknown
	}
	// the target is always a "good" state
	memo[len(memo)-1] = good

	// Start at the second to last field, and proceed
	// "bottom up" (towards the start of the path)
	for i := len(nums) - 2; i >= 0; i-- {
		// We do not want to jump beyond the end of the
		// path, so the maximum distance we want to jump
		// is the minimum of the current field and the
		// entire path length
		furthestJump := min(i+nums[i], len(nums)-1)
		for j := i + 1; j <= furthestJump; j++ {
			// If we would land on a good spot J then
			// that implies that the current spot I is
			// also a good spot
			if memo[j] == good {
				memo[i] = good
				break
			}
		}
	}
	return memo[0] == good
}

func main() {
	path := [...]int{2, 3, 1, 1, 4}
	succ := canJump(path[:])
	fmt.Println("Can jump:", succ)
}
