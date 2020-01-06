package main

import (
	"fmt"
)

const nofItems = 10
const maxWeight = 67
const undef = -1

// Memoization table
var value [maxWeight+1][nofItems+1]int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Returns the maximum achievable value for trying to grab the first 'item'
// number of items and put them into the knapsack with capacity. Each
// item has a weight and value, as given in the slices weights and values
// (index starts at 1).
// https://en.wikipedia.org/wiki/Knapsack_problem
func maxValue(item, capacity int, weights, values []int) int {

	if item == 0 || capacity <= 0 {
        return 0
	}
    
    if (value[capacity][item-1] == undef) {	
    	// value[i-1, j] has not been calculated, we have to call function maxValue
		value[capacity][item-1] = maxValue(item-1, capacity, weights, values) 
	}        
  
    if weights[item] > capacity {                   
    	// item i cannot fit in the bag ("put it back")
        value[capacity][item] = value[capacity][item-1]
	} else {
        if (value[capacity - weights[item]][item-1] == undef) {
			// value[i-1, j-w[i]] has not been calculated, we have to call function maxValue
			value[capacity - weights[item]][item-1] = maxValue(item-1, capacity - weights[item], weights, values)
		}

		// store the (intermediate) maximum value
		// max(value[i-1,j], value[i-1, j-w[i]] + v[i])
		previousValue := value[capacity][item-1]
		newValue := value[capacity - weights[item]][item-1] + values[item]
        value[capacity][item] = max(previousValue, newValue)
	}
    return value[capacity][item]
}

func initMemo() {
	for weight := 0; weight <= maxWeight; weight++ {
		for item := 0; item <= nofItems; item++ {
			value[weight][item] = undef
		}
	}
}

// Solves the 0/1  knapsack problem which allows to grab only the entire
// item ("take it or leave it"). A greedy approach won't work here, as
// in the worst case it would select only half of the possible (optimum)
// value.
// w: weight of each item (starting from index 1)
// v: value of each item (starting from index 1)
func grabAndRun(weights, values []int) {
	initMemo()
	maxValue := maxValue(nofItems, maxWeight, weights, values)
	fmt.Println("Maximum value", maxValue)
}

func main() {
	// Indexing starts at 1
	weights := []int{undef, 23, 26, 20, 18, 32, 27, 29, 26, 30, 27}
	values := []int{undef, 505, 352, 458, 220, 354, 414, 498, 545, 473, 543}
	grabAndRun(weights, values)
}
