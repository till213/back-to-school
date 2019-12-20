package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

// n-th Fibonacci number (recursive)
func rfib(n int) int {
  if n <= 1 {
	  return n
  } else {
	  return rfib(n-1) + rfib(n-2)
  }
}

// n-th Fibonacci number (iterative)
func ifib(n int) int {
	var r int
	if n <= 1 {
		return n
	} else {
		p := 0;
		q := 1;
		for i := 2; i <= n; i++ {
			r = p + q;
			p = q;
			q = r;
		}
	}
	return r
  }

func main() {
	var n int
	var err error
	if len(os.Args) > 1 {
		n, err = strconv.Atoi(os.Args[1])
	} else {
		err = errors.New("No arg")
	}
	if err == nil {
		f1 := rfib(n)
		fmt.Printf("%v-th Fibonacci number (recursive): %v\n", n, f1)
		f2 := ifib(n)
		fmt.Printf("%v-th Fibonacci number (iterative): %v\n", n, f2)
	} else {
		fmt.Print("Parse error: ", err)
	}
	
}
