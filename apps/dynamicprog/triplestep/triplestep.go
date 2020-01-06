package main

import (
	"fmt"
	"os"
	"strconv"
)

var memo []int

// Returns the number of ways the child can jump up
// the stairs with 'steps' steps left.
// The child can jump either 1, 2 or 3 steps at once,
// but must land on the last step, that is, the remaining
// step count must never get negative
func tripleStep(steps int) int {
	switch steps {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	default:
		if memo[steps-1] == 0 {
			memo[steps-1] = tripleStep(steps-1) + tripleStep(steps-2) + tripleStep(steps-3)
		}
		return memo[steps-1]
	}
}

func main() {
	var steps int
	var err error
	if len(os.Args) > 1 {
		steps, err = strconv.Atoi(os.Args[1])
	} else {
		steps = 100
	}
	if err == nil {
		memo = make([]int, steps)
		n := tripleStep(steps)
		fmt.Printf("Number of ways to jump up %d steps: %d\n", steps, n)
	} else {
		fmt.Println("Parse error: ", err)
	}
}
