package main

import (
	"fmt"
)

type Triple struct {
    a, b, c int
}

func createTriple(a, b, c int) Triple {
    var triple Triple
    
    if b < a { 
        a, b = b, a
    } 
    if c < b { 
        b, c = c, b      
        if b < a { 
            a, b = b, a
        } 
    } 

    triple.a = a
    triple.b = b
    triple.c = c
    return triple
}


func twoSum(nums []int, targetSum int, freq map[int]int) map[Triple]bool {
    res := make(map[Triple]bool, 0)
    for sum2, v := range freq {
        if v > 0 {
            freq[sum2]--
            sum3 := targetSum-sum2
            count := freq[sum3]
            if count > 0 {
                triple := createTriple(-targetSum, sum2, sum3)
                res[triple] = true
            }
            freq[sum2]++
        }
    }
    return res
}

func threeSum(nums []int) [][]int {
    triples := make(map[Triple]bool, 0)
    res := make([][]int, 0)
    freq := make(map[int]int, len(nums))
    for _, v := range nums {
        freq[v]++
    }
    
    for first := 0; first < len(nums)-2; first++ { 
        sum1 := nums[first]
        freq[sum1]--
        currTriples := twoSum(nums, -sum1, freq)
        for k := range currTriples {
            triples[k] = true
        }
        freq[sum1]++
    }
    
    for k := range triples {
        arr := make([]int, 3)
        arr[0] = k.a
        arr[1] = k.b
        arr[2] = k.c
        res = append(res, arr)
    }
    
    return res
}
func threeSum2(nums []int) [][]int { 

	triples := make(map[Triple]bool, 0)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ { 
		// Reduce the "three sum" problem to a 
        // "two sum problem": the trick is to
        // use a hash map (set) to remember numbers
        // that "we have seen" as we progress with j
		numSet := make(map[int]bool)
		for j := i + 1; j < len(nums); j++ {
			x := -(nums[i] + nums[j])
			if numSet[x] { 
				triple := createTriple(nums[i], nums[j], x)
                triples[triple] = true
			}  else { 
				numSet[nums[j]] = true
			} 
		} 
	} 
	for k := range triples {
        arr := []int{k.a, k.b, k.c}
        res = append(res, arr)
    }
	return res
} 

func main() {
	A := []int{-1,0,1,2,-1,-4}
	res := threeSum(A)
	fmt.Println(res)
	res = threeSum2(A)
	fmt.Println(res)
}