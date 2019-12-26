package main

import (
	"errors"
	"fmt"
	"os"
)

func uniqueChars(s string) bool {
	runes := make(map[rune]bool)
	unique := true
	for _, r := range s {
		if runes[r] {
			// Key exists already
			unique = false
		} else {
			runes[r] = true
		}
	}
	return unique
}

func uniqueCharsNoAdditionalDataStruct(s string) bool {
	unique := true
	for i, r1 := range s[:len(s)-1] {
		for _, r2 := range s[i+1:] {
			if r1 == r2 {
				unique = false
				break
			}
		}
		if !unique {
			break
		}
	}
	return unique
}

func main() {
	var s string
	var err error
	if len(os.Args) > 1 {
		s = os.Args[1]
	} else {
		err = errors.New("No arg")
	}
	if err == nil {
		unique := uniqueChars(s)
		fmt.Println(s, "has unique chars:", unique)
		unique = uniqueCharsNoAdditionalDataStruct(s)
		fmt.Println(s, "has unique chars:", unique)
	} else {
		fmt.Println("Parse error: ", err)
	}
}
