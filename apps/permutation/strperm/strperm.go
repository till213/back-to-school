package main

import (
	"fmt"
	"os"
)

// Permute returns all permutations of the given string s. There are
// n! permutation, where n = len(s).
func Permute(s string) []string {
	l := len(s)
	var perms []string

	if l == 0 {
		perms = make([]string, 0)
	} else if l == 1 {
		perms = make([]string, 1)
		perms[0] = s
	} else {
		l1 := l - 1
		subString := s[:l1]
		c := s[l1:]
		u := Permute(subString)
		for _, s1 := range u {
			for i := 0; i <= l1; i++ {
				// Insert c between every position of string s1:
				// c**, *c*, **c
				p := s1[:i] + c + s1[i:]
				perms = append(perms, p)
			}
		}
	}
	return perms
}

func main() {

	var input string
	args := os.Args[1:]
	if len(args) > 0 {
		input = args[0]
	} else {
		input = "abc"
	}
	s := Permute(input)

	fmt.Println(s)

}
