package main

import (
	"fmt"
	"math/rand"

	"github.com/till213/back-to-school/data/tree"
)

func updatePathCounts(pathCounts *map[int]int, runningSum, deltaCount int) {
	newCount := (*pathCounts)[runningSum] + deltaCount
	if newCount > 0 {
		// Update
		(*pathCounts)[runningSum] = newCount
	} else {
		// Remove
		delete((*pathCounts), runningSum)
	}
}

// currentSumCount := (*sumCounts)[runningSum]
// 	currentSumCount++
// 	(*sumCounts)[runningSum] = currentSumCount

func travelPathWithSum(n *tree.Node, pathCounts *map[int]int, runningSum, targetSum int) int {
	var count int
	if n == nil {
		return 0
	}

	runningSum += n.Value
	// Get the previous count of paths summing up to sum
	sum := runningSum - targetSum
	count = (*pathCounts)[sum]
	// Did we find yet another running sum summing up to
	// target sum?
	if runningSum == targetSum {
		count++
	}

	// increment path count for current running sum
	updatePathCounts(pathCounts, runningSum, 1)
	// recursively add the subtree counts
	count += travelPathWithSum(n.Left, pathCounts, runningSum, targetSum) +
		travelPathWithSum(n.Right, pathCounts, runningSum, targetSum)
	// restore path count for current running sum, as we travel
	// back up the tree
	updatePathCounts(pathCounts, runningSum, -1)

	return count
}

func pathWithSumCount(n *tree.Node, targetSum int) int {
	pathCounts := make(map[int]int)
	return travelPathWithSum(n, &pathCounts, 0, targetSum)
}

func main() {
	const targetSum = 10
	rand.Seed(1001)
	root := tree.CreateRandomTree(-16, 16, 5, nil)
	tree.PrintTree(root, 0)

	count := pathWithSumCount(root, targetSum)
	fmt.Printf("Count of continuous paths with target sum %d: %d\n", targetSum, count)
}
