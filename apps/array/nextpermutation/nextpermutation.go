package main

import (
	"fmt"
)

const notFound = -1

// Returns the first position where the number decreases,
// searching "from the right"
// Example: 1, 2, 4, 3 -> index = 1 (number 2)
func firstDecreasePosition(nums []int) int {
	for i := len(nums)-2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			return i
		}
	}
	return notFound
}

// Returns the position of next greater element of nums[i], from
// the range i+1, ..., len(nums-1). In case multiple candidate
// exists, e.g. ..., 2, 4, 3, 3, then the last candidate is
// returned, such that with the subsequent swap the 2 goes as
// "far behind as possible":
// * decrease at 2
// * swap 2 with last 3: 3, 4, 3, 2
// * reverse after 3: 3, 2, 3, 4
func nextGreaterIndex(nums []int, i int) int {
	current := nums[i]
	next := 1<<31
	nextIndex := notFound
	for j := i+1; j < len(nums); j++ {
		if nums[j] > current && nums[j] <= next {
			next = nums[j]
			nextIndex = j 
		}
	}
	return nextIndex
}

func reverse(nums []int, i int) {
	// https://github.com/golang/go/wiki/SliceTricks
	for left, right := i, len(nums)-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
}

func nextPermutation(nums []int) {	
	i := firstDecreasePosition(nums)
	if i != notFound {
		nextIndex := nextGreaterIndex(nums, i)
		// swap with next greater number
		nums[i], nums[nextIndex] = nums[nextIndex], nums[i]
		// reverse the elements (which are, left to right, in decreasing order) 
		// just after the swapped position i
		i++
	} else {
		// reverse the entire permutation, e.g. 4, 3, 2, 1 -> 1, 2, 3, 4
		i = 0
	}
	reverse(nums, i)
}

func main() {
	A := []int{2,3,1,3,3}
	nextPermutation(A)
	fmt.Println(A)
}
