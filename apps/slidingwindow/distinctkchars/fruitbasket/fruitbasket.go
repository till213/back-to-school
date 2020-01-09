package main

import "fmt"

const numBaskets = 2

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func totalFruit(tree []int) int {
	// count the number of distinct fruits in the current sliding window
	var counter int
	maxFruit := 0
	// Count the number of each fruit type (key)
	fruitMap := make(map[int]int)

	// sliding window
	start := 0
	end := 0
	for end < len(tree) {
		fruit := tree[end]
		if fruitMap[fruit] == 0 {
			// each time we get a fruit that we have not yet
			// picked increment the "different fruits"-counter
			counter++
		}

		// count the occurence of each fruit
		fruitMap[fruit]++
		// advance to next "tree"
		end++

		// does the current window exceed the allowed numBaskets distinct fruits?
		for counter > numBaskets {
			// we have encountered more than numBaskets different fruits
			// -> advance 'start' for as long as we have more than
			//    numBaskets different fruits, "unsetting" ("forgetting")
			//    the corresponding fruits as we go...
			fruit2 := tree[start]
			if fruitMap[fruit2] == 1 {
				counter--
			}
			fruitMap[fruit2]--
			// ... until we are at numBaskets distinct fruits again
			start++
		}
		maxFruit = max(maxFruit, end-start)
	}
	return maxFruit
}

func main() {
	trees := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	maxFruits := totalFruit(trees)
	fmt.Printf("Maximum number of fruits (with %d baskets): %d\n", numBaskets, maxFruits)
}
