package main

import (
	"fmt"
)

// Naive: O(n²)
func twoSum(nums []int, target int) []int {
	found := false
	var x, y int
	for i := 0; !found && i < len(nums)-1; i++ {
		for j := i + 1; !found && j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				found = true
				x, y = i, j

			}
		}
	}
	return []int{x, y}
}

func twoSum2(nums []int, target int) []int {
    // "You may not use the same element twice"
    // -> we can use a hashmap to store the previous occurence of any given number
    numMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
        x := target - nums[i]
        v, ok := numMap[x]
        if ok {
            return []int{v, i}
		} 
		numMap[nums[i]] = i           
	}
    // "You may assume that each input would have exactly one solution"
    // -> should never come here :)
    return nil
}

func main() {
	const target = 6
	A := []int{3, 2, 4}
	res := twoSum(A, target)
	fmt.Println(res)
	res = twoSum2(A, target)
	fmt.Println(res)
}
