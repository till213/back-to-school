package main

import (
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func insertOrDelete(shorter, longer string) bool {
	for _, r1 := range shorter {
		found := false
		for _, r2 := range longer {
			if r1 == r2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func oneEdit(str1, str2 string) bool {
	var diff int
	for i := range str1 {
		if str1[i] != str2[i] {
			diff++
		}
		if diff > 1 {
			break
		}
	}
	return diff <= 1
}

func oneAway(str1, str2 string) bool {
	l1 := len(str1)
	l2 := len(str2)
	if abs(l1-l2) == 1.0 {
		if l1 < l2 {
			return insertOrDelete(str1, str2)
		}
		return insertOrDelete(str2, str1)
	} else if l1 == l2 {
		return oneEdit(str1, str2)
	}
	return false

}

func main() {
	var str1, str2 string

	if len(os.Args) > 2 {
		str1 = os.Args[1]
		str2 = os.Args[2]
	} else {
		str1 = "ab"
		str2 = "aa"
	}

	one := oneAway(str1, str2)
	fmt.Println(str1, str2, "one away", one)
}
