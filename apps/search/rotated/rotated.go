package main

import (
	"os" 
	"strconv" 
	"fmt"
)

const notFound = -1

func binarySearch(nums []int, l, r int, target int) int {
	m := (l + r) / 2;
	if target == nums[m] {
		return m
	}
	if l == r {
		return notFound
	}
	if nums[l] <= nums[m] {
		// array [l...m] is sorted
		if nums[l] <= target && target <= nums[m] {
			// target must lie in left array
			return binarySearch(nums, l, m-1, target)
		}
		// target must lie in right array
		return binarySearch(nums, m+1, r, target)
	} 
	// array [m...r] must be sorted
	if nums[m] <= target && target <= nums[r] {
		// target must lie in right array
		return binarySearch(nums, m+1, r, target)
	}
	// target must lie in left array
	return binarySearch(nums, l, m-1, target)
}

func iterativeBinarySearch(nums []int, l, r int, target int) int {
	var m int
	for l <= r {
		m = (l + r) / 2;
		if nums[m] == target {
			return m
		}
		if nums[l] <= nums[m] {
			// array [l...m] is sorted
			if nums[l] <= target && target <= nums[m] {
				// target must lie in left array
				r = m - 1
			} else {
				// target must lie in right array
				l = m + 1
			}
		} else {
			// array [m...r] must be sorted
			if nums[m] <= target && target <= nums[r] {
				// target must lie in right array
				l = m + 1
			} else {
				// target must lie in left array
				r = m - 1
			}
		}
	}
	return notFound
}

func recursiveSearch(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func iterativeSearch(nums []int, target int) int {
	return iterativeBinarySearch(nums, 0, len(nums)-1, target)
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
    var n int
	var err error
	if len(os.Args) > 1 {
		n, err = strconv.Atoi(os.Args[1])
	} else {
		n = 1
	}
	if err == nil {
		index := recursiveSearch(nums, n)
		fmt.Printf("Recursive: Found %d at index: %d\n", n, index)
		index = iterativeSearch(nums, n)
		fmt.Printf("Iterative:Found %d at index: %d\n", n, index)
	fmt.Println("Index", index)
	} else {
		fmt.Println("Parse error: ", err)
	}
}
