package main

import (
	"fmt"
	"strings"
)

const errorRange = "ERROR: number must be ]0.0, 1.0["
const errorLength = "TOO LONG"

func toString(n float64) string {
	var sb strings.Builder

	if n <= 0.0 || n >= 1.0 {
		return errorRange
	}

	sb.WriteString(".")
	for n > 0.0 {
		if sb.Len() >= 32 {
			return errorLength
		}
		r := n * 2
		if r >= 1 {
			sb.WriteString("1")
			n = r - 1
		} else {
			sb.WriteString("0")
			n = r
		}
	}
	return sb.String()
}

func main() {
	n := 1.0/2.0 + 1.0/8.0 + 1.0/256.0
	fmt.Printf("Number %f in binary: %s\n", n, toString(n))
}
