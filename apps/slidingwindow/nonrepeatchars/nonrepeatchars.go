package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	// counts the number of "invalidations", where an invalidation
	// is to have repeating characters
	// invariant: count is 0 for the longest substring
	count := 0
	maxLen := 0
	start := 0
	end := 0
	for end < len(s) {
		ch := s[end]
		if charMap[ch] > 0 {
			// We have seen this character already
			count++
		}

		charMap[ch]++
		end++

		// does an invalidation exist? If so, "fix it" by
		// advancing the 'start' index
		for count > 0 {
			ch1 := s[start]
			if charMap[ch1] == 2 {
				// the count of ch1 is just about to
				// go to 1 again, so we decrease the
				// invalidations count
				count--
			} else {

			}
			charMap[ch1]--
			// now we advance 'start' for as long as
			// the invariant is violated
			start++
		}

		maxLen = max(end-start, maxLen)

	}
	return maxLen
}

func main() {
	s := "pwwkew"
	maxLen := lengthOfLongestSubstring(s)
	fmt.Printf("Length of longest substring of %s without repeated characters: %d\n", s, maxLen)
}
