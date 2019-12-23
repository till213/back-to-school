package jumpball

import (
	"fmt"
)

// CanStop returns whether the jumping ball, starting at
// the first field of path, can eventually stop, starting
// off with speed. At each step the jumping ball
// can increase (+1), keep (0) or decrease (-1) its speed.
//
// The jumping ball must never land on a field (element)
// with value false, which indicates a "spike" and the
// jumping ball would explode.
func CanStop(path []bool, pos, speed int) bool {
	if pos < len(path) {
		fmt.Printf("Current position path: %v position: %v speed: %v\n", path[pos], pos, speed)
	} else {
		fmt.Printf("OVERSHOT PATH, position: %v speed: %v\n", pos, speed)
	}
	if speed == 0 {
		return true
	} else if pos >= len(path) || pos+speed-1 >= len(path) || !path[pos] {
		return false
	} else {
		return CanStop(path, pos+speed+1, speed+1) ||
			   CanStop(path, pos+speed, speed) ||
			   CanStop(path, pos+speed-1, speed-1)
	}
}

type state struct {
	pos, speed int
}

// CanStopWithMemo is an optimised variant of the above
func CanStopWithMemo(path []bool, pos, speed int, memo map[state]bool) bool {

	v, ok := memo[state{pos, speed}]
	if ok {
		return v
	}

	if pos < len(path) {
		fmt.Printf("Current position path: %v position: %v speed: %v\n", path[pos], pos, speed)
	} else {
		fmt.Printf("OVERSHOT PATH, position: %v speed: %v\n", pos, speed)
	}
	if speed == 0 {
		memo[state{pos, speed}] = true
		return true
	} else if pos >= len(path) || pos+speed-1 >= len(path) || !path[pos] {
		memo[state{pos, speed}] = false
		return false
	} else {

		for _, adjust := range []int{speed+1, speed, speed-1} {
			canStop := CanStopWithMemo(path, pos+adjust, adjust, memo)
			if canStop {
				memo[state{pos, speed}] = true
				return true
			}
		}
		memo[state{pos, speed}] = false
		return false
	
	}
}
