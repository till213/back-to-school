package main

import (
	"fmt"
	"sort"
)

// 09:xx, 19:xx, 2*3*:xx
var maxNumbers [][]int

const zero = '0'

func timeToNumbers(time string) []int {
	numbers := make([]int, 4)
	numbers[0] = int(time[0] - zero)
	numbers[1] = int(time[1] - zero)
	numbers[2] = int(time[3] - zero)
	numbers[3] = int(time[4] - zero)
	return numbers
}

func nextClosestTime(time string) string {

	maxNumbers = [][]int{
		{2, 9, 5, 9},
		{2, 9, 5, 9},
		{2, 3, 5, 9},
	}

	// hour1:hour2:min1:min2
	numbers := timeToNumbers(time)
	sortedNumbers := make([]int, 4)
	copy(sortedNumbers, numbers)
	sort.Ints(sortedNumbers)
	done := false
	// Try to increase any number with the smallest possible
	// next number, while remaining within the max limits for
	// each number
	hour1 := numbers[0]
	for i := 3; i >= 0 && !done; i-- {
		currentNumber := numbers[i]
		if currentNumber < maxNumbers[hour1][i] {
			for j := 0; j <= 3 && !done; j++ {
				nextNumber := sortedNumbers[j]
				if currentNumber < nextNumber && nextNumber <= maxNumbers[hour1][i] {
					numbers[i] = nextNumber
					done = true
				}
			}
		}
	}
	if !done {
		// no number could be increased without exceeding
		// any limit, so set *every* number to the lowest
		// possible value (which must be either 0, 1 or 2
		// from hour1)
		for i := 0; i <=3; i++ {
			numbers[i] = sortedNumbers[0]
		}
	}

	str := fmt.Sprintf("%d%d:%d%d", numbers[0], numbers[1], numbers[2], numbers[3])
	return str
}

func main() {
	//time := "19:34"
	time := "18:42"
	next := nextClosestTime(time)
	fmt.Println("Current:", time, "Next Time:", next)

	time = "20:59"
	next = nextClosestTime(time)
	fmt.Println("Current:", time, "Next Time:", next)

	time = "23:59"
	next = nextClosestTime(time)
	fmt.Println("Current:", time, "Next Time:", next)
}
