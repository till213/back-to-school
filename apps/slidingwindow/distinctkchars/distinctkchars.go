package main

import (
	"fmt"
)

const offset = 'A'
const maxValue = 1 << 31

type charmap []int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://leetcode.com/problems/minimum-window-substring/discuss/26808/here-is-a-10-line-template-that-can-solve-most-substring-problems
func lengthOfLongestSubstringKDistinct(str string, k int) int {
	// sliding window
	var start, end int
	// count the number of distinct characters in the current sliding window
	var counter int
	maxLen := 0
	// Could be a hashmap as well, depending on performance vs space
	charmap := make(charmap, 128)
	
	// We assume ASCII characters (only)
	for end < len(str) {
		c1 := str[end] - offset
		if (charmap[c1] == 0) {
			// each time we hit a character that we have not yet
			// seen increment the "different characters"-counter
			counter++;
		}

		// count the occurence of each character
		charmap[c1]++;
		// advance to next characters
		end++;
		
		// does the current window exceed the allowed k distinct characters?
		for counter > k {
			// we have encountered more than k different characters
			// -> advance 'start' for as long as we have more than
			//    k different characters, "unsetting" ("forgetting")
			//    the corresponding characters as we go...
			c2 := str[start] - offset;
			if charmap[c2] == 1 {
				counter--;
			}
			charmap[c2]--;
			// ... until we are at k distinct characters again
			start++;
		}
		maxLen = max(maxLen, end - start);
	}
	return maxLen
}

// Length of longest substring without repeating characters.
// Examples: ABDEFGABEF -> BDEFGA and DEFGAB, with length 6 each
//           BBBB -> B, with length 1
// https://www.geeksforgeeks.org/length-of-the-longest-substring-without-repeating-characters/
func lengthOfLongestSubstringNonRepeatedChar(str string) int {
	// sliding window
	var start, end int
	// count the number of distinct characters in the current sliding window
	var counter int
	maxLen := 0
	charmap := make(charmap, 128)

    for end < len(str) {
		c1 := str[end] - offset
		if (charmap[c1] > 0) {
			// each time we hit a character twice (-> > 0) we increment the
			// "repeated character"-counter
			counter++;
		}

		charmap[c1]++;
		end++;

		for counter > 0 {
			// we have encountered a given character twice
			// -> advance 'start' for as long as we have a given
			//    character twice (the counter is > 0)
			c2 := str[start] - offset;
			if charmap[c2] > 1 {
				// we have just found the character that was found
				// twice: decrease the counter again
				counter--;
			}
			charmap[c2]--;
			start++;
		}
      	maxLen = max(maxLen, end - start);
    }
    return maxLen;
}

func initCharmap(str string) charmap {
	charmap := make(charmap, 128)
	for _, c := range str {
		charmap[c-offset]++
	}
	return charmap
}

// https://leetcode.com/problems/minimum-window-substring/
// Input: S = "ADOBECODEBANC", T = "ABC"
// Output: "BANC"
func minWindow(s, t string) string {
    // sliding window
	var start, end, minStart int
	var res string		
	minLen := maxValue
	// "Fill up" the charmap with the characters of t
	charmap := initCharmap(t)
	// The counter counts the number characters of t still remaining
	// (that we have not yet read in s)
	counter := len(t)	

	for (end < len(s)) {
		c1 := s[end] - offset
		if (charmap[c1] > 0) {
			// Each time a character hits the 0 mark we decrease the counter
			counter--;
		}

		// Other characters not in t will simply decrease the count in the character map to negative
		// Actually we would not care about them (we only care about characters in t), but for as
		// long as the other character counts never go above 0 (they shall not affect our counter)
		// we are fine
		charmap[c1]--;
		end++;

		// Did we "use up" all characters in t?
		for counter == 0 {
			// Update the minimum length
			if (minLen > end - start) {
				minLen = end - start;
				minStart = start;
			}

			// Now "fill up" the characters again
			c2 := s[start] - offset;
			charmap[c2]++;
			// Again we are actually only interested in the characters in t,
			// but they will it the "above zero" mark before all other characters
			// when "unread" ("start" has passed them)
			if charmap[c2] > 0 {
				counter++;
			}			
			start++;
		}
	}
	
	if minLen != maxValue {
		// Again, we assume ASCII only
		res = s[minStart:minStart + minLen]
	} else {
		res = ""
	}
	return res
  }

func main() {
	const k = 2

	str := "abababcdcdaaaaaaabbbb"
	maxLen := lengthOfLongestSubstringKDistinct(str, k)
	fmt.Printf("Longest sequence with %d distinct characters in %s: %d\n", k, str, maxLen)

	str = "abcddefghij"
	maxLen = lengthOfLongestSubstringNonRepeatedChar(str)
	fmt.Printf("Longest sequence without repeating characters in %s: %d\n", str, maxLen)

	str = "ADOBECDDDODEBANC"
	t :=  "ABC"
	res := minWindow(str, t)
	fmt.Printf("Minimum window of characters %s in %s: %s\n", t, str, res)
}
