package main

import (
	"fmt"
	"os"
	"sort"
)

func lessThan(i, j int) bool {
	return i < j
}

func isPermutation(s1, s2 string) bool {
	if len(s1) == len(s2) {
		sl1 := []rune(s1)
		sl2 := []rune(s2)

		lessThan1 := func(i, j int) bool {
			return sl1[i] < sl1[j]
		}
		lessThan2 := func(i, j int) bool {
			return sl2[i] < sl2[j]
		}
		sort.Slice(sl1, lessThan1)
		sort.Slice(sl2, lessThan2)

		for i, s := range sl1 {
			if s != sl2[i] {
				return false
			}
		}
		return true
	}
	return false
}

func isPermutation2(s1, s2 string) bool {
	if len(s1) == len(s2) {
		sm := make(map[rune]int)
		for _, r1 := range s1 {
			sm[r1]++
		}
		for _, r2 := range s2 {
			sm[r2]--
			if sm[r2] < 0 {
				return false
			}
		}
		return true
	}
	return false
}

func main() {
	var s1, s2 string
	if len(os.Args) > 2 {
		s1 = os.Args[1]
		s2 = os.Args[2]
	} else {
		s1 = "abc"
		s2 = "cba"
	}

	perm := isPermutation(s1, s2)
	fmt.Println(s1, "is permutation of", s2, ":", perm)
	perm = isPermutation2(s1, s2)
	fmt.Println(s1, "is permutation of", s2, ":", perm)
}
