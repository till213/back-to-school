package main

import (
    "fmt"
    "github.com/emirpasic/gods/maps/treemap"
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

func oddEvenJumpsWithTreeMap(A []int) int {
    N := len(A)
    if (N <= 1) {
        return N;
    }
    odd := make([]bool, N)
    even := make([]bool, N)
    odd[N-1] = true
    even[N-1] = true

    // TreeMap<Integer, Integer> vals = new TreeMap();
    // The map vals will map values v = A[i] to indices i.
    vals := treemap.NewWithIntComparator()
    vals.Put(A[N-1], N-1)
    for i := N-2; i >= 0; i-- {
        v := A[i]
        j, found := vals.Get(v)
        if (found) {
            odd[i] = even[j.(int)]
            even[i] = odd[j.(int)]
        } else {
            lowerKey, lowerValue := vals.Floor(v)
            higherKey, higherValue := vals.Ceiling(v)

            // In an "even" jump we want to jump from i to j, where 
            // * (1) A[i] >= A[j] and
            // * (2) j is the smallest such index
            //
            // Both is answered by the Tree Map
            // * The next lower key ("Floor") is the largest such value A[j]
            // * As a given key can only exist once in the (Tree) Map we also 
            //   get the "smallest such index j", because we iterate backwards
            //   over A and keep updating (inserting, "Put") the existing keys
            //
            // So does such a lower value (which we use as key in the treemap) exist?
            if lowerKey != nil {
                // If so, we can only do the "even" jump from i when it is possible
                // to do an "odd" jump from j towards the goal (end of A)
                even[i] = odd[lowerValue.(int)]
            }
            // In analogy for "odd" jumps, with:
            // * (1) A[i] <= A[j] and
            // * (2) j is the smallest such index
            if higherKey != nil {
                odd[i] = even[higherValue.(int)]
            }
        }
        vals.Put(v, i)
    }

    ans := 0;
    //for (boolean b: odd)
    for _, v := range odd {
        if (v) {
            ans++
        }
    }
    return ans
}

func main() {
	A := []int{5,1,3,4,2}
	count := oddEvenJumps(A)
    fmt.Println("Jumps:", count)
    count = oddEvenJumpsWithTreeMap(A)
	fmt.Println("Jumps (with treemap):", count)
}