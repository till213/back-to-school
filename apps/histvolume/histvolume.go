package main

import (
	"fmt"
)

// NotFound indicates that no (further) histogram bar has been found
const NotFound = -1

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Returns the index of the next bar, starting from index (excluding);
// the next bar is either
// - higher than height
// - the highest of all remaining bars
// The "negative volume" (sum of all intermediate bars) is also
// returned
// Returns (NotFound, 0) if no such bar exists
func nextBarIndex(index, height int, hist []int) (nextIndex, negVol int) {
	nextIndex = NotFound
	negVol = 0
	prevHeight := 0
	prevIndex := NotFound
	for i := index + 1; i < len(hist) && nextIndex == NotFound; i++ {
		if currentHeight := hist[i]; currentHeight > 0 {
			if currentHeight >= height {
				// We have found a bar which is higher than
				// the previous one (at pos); we are done
				nextIndex = i
			} else if currentHeight > prevHeight {
				// We increase the "negative volume", optimistically
				// that we will still find a next higher bar
				negVol += currentHeight
				prevHeight = currentHeight
				prevIndex = i
			} else {
				// Lower than the previous bar; we are done
			}
		}
	}
	if nextIndex == NotFound && prevIndex != NotFound {
		nextIndex = prevIndex
		// Adjust the "negative volume"
		negVol -= prevHeight
	}
	return
}

// HistVolume returns the volume of water that can
// be filled "into the histogram" given my hist. Each
// bar has an assumed width of 1
func HistVolume(hist []int) int {
	var nextIndex int
	currIndex, currNegVol := nextBarIndex(-1, 0, hist)
	vol := 0

	for currIndex != NotFound {
		currHeight := hist[currIndex]
		fmt.Printf("Current Height: %v, current index: %v, current neg. Vol: %v\n", currHeight, currIndex, currNegVol)

		// The next bar
		// - is higher or equal than the current one (1)
		// - is the highest bar of all remaining (this bar is the highest) (2)
		// - does not exist (3)
		//
		// (1) This bar defines the maximum volume (all intermediate bars add to "negative volume")
		// (2) Next bar defines the maximum volume (all intermediate, lower bars add to "negative volume")
		// (3) We are done

		nextIndex, currNegVol = nextBarIndex(currIndex, currHeight, hist)
		if nextIndex != NotFound {
			nextHeight := hist[nextIndex]
			vol += (nextIndex-currIndex-1)*min(currHeight, nextHeight) - currNegVol
		}
		currIndex = nextIndex
	}
	return vol
}

func main() {
	hist := []int{0, 0, 4, 0, 0, 6, 0, 0, 3, 0, 5, 0, 1, 0, 0, 0}
	vol := HistVolume(hist)
	fmt.Printf("History: %v, volume: %v\n", hist, vol)
}
