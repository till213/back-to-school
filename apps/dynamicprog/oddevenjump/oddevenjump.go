package main

import (
	"fmt"
)

const Max = 1<<31
const Min = -1<<31
const Impossible = -1

type State struct {
	odd bool
	pos int
}

func minIndex(A []int, start int) int {
    min := Max
	minIndex := Impossible
    for j := start+1; j < len(A); j++ {
        if A[j] >= A[start] && A[j] < min {
            min = A[j]
            minIndex = j
        }
    }
    return minIndex
}

func maxIndex(A []int, start int) int {
    max := Min
    maxIndex := Impossible
    for j := start+1; j < len(A); j++ {
        if A[j] <= A[start] && A[j] > max {
            max = A[j]
            maxIndex = j
        }
    }
    return maxIndex
}

func nextPossibleOddJump(A []int, start int) int {
    minIndex := minIndex(A, start)
    if minIndex != Impossible {
        return minIndex
    }
    return Impossible
}

func nextPossibleEvenJump(A []int, start int) int {
    maxIndex := maxIndex(A, start)
    if maxIndex != Impossible {
        return maxIndex
    }
    return Impossible
}

func nextPossibleJump(A []int, start int, odd bool) int {
    if odd {
        return nextPossibleOddJump(A, start)
    }
    return nextPossibleEvenJump(A, start)
}

func canJump(A []int, start int, odd bool, memo map[State]bool) bool {
	var jumpPossible, ok bool
    if start == len(A) - 1 {
        return true
	}
	state := State{odd, start}
	jumpPossible, ok = memo[state]
	if ok {
		return jumpPossible
	} else {
		next := nextPossibleJump(A, start, odd)
		if next != Impossible {
			jumpPossible = canJump(A, next, !odd, memo)
		} else {
			jumpPossible = false
		}
		memo[state] = jumpPossible
		return jumpPossible
	} 
}

func oddEvenJumps(A []int) int {
	count := 0
    memo := make(map[State]bool, 0)
    for i := 0; i < len(A); i++ {
        if canJump(A, i, true, memo) {
            count++
        } 
    }
    return count
}

func main() {
	A := []int{5,1,3,4,2}
	count := oddEvenJumps(A)
	fmt.Println("Jumps", count)
}