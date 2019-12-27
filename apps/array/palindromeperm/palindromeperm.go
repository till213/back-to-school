package main

import (
	"fmt"
	"os"
)

const blank = 32

func palindromePerm(str string) bool {
	var nofBlanks int
	var nofOdds int
	cm := make(map[rune]int)

	for _, v := range str {
		// Blanks are not part of a palindrome
		if v != blank {
			cm[v]++
		} else {
			nofBlanks++
		}
	}
	// In a palindrome at most one rune appears an odd
	// number of times, and only if the length of the
	// string itself is odd; all other runes must appear an
	// even number of times
	// abcabacba: a = 4, b = 3, c = 2 (even, odd, even - length: 9 = odd)
	// abba: a = b = 2 (even, even - length: 4 = even)
	// abaa: a = 3, b = 1 - however length 4 is even -> NOT a palindrome
	evenPalindrome := (len(str)-nofBlanks)%2 == 0

	for _, v := range cm {
		if v%2 == 1 {
			nofOdds++
		}
		if evenPalindrome {
			if nofOdds > 0 {
				return false
			}
		} else {
			if nofOdds > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	var s string
	if len(os.Args) > 1 {
		s = os.Args[1]
	} else {
		s = "tact coa"
	}

	palindrome := palindromePerm(s)
	fmt.Println(s, "is palindrome permutation:", palindrome)
}
