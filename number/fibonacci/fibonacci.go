package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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
		p := 0
		q := 1
		for i := 2; i <= n; i++ {
			r = p + q
			p = q
			q = r
		}
	}
	return r
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	p, q := 0, 1
	return func() int {
		q, p = p, q+p
		return p
	}
}

func fibchannel(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		c <- x
	}
	close(c)
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
		fmt.Println("Fibonacci (closure)")
		f := fibonacci()
		for i := 0; i < n; i++ {
			fmt.Println(f())
		}

		fmt.Println("Fibonacci (channel)")
		c := make(chan int, n)
		go fibchannel(cap(c), c)
		// range c receives values from the channel repeatedly until it is closed.
		for i := range c {
			fmt.Println(i)
		}
	} else {
		fmt.Println("Parse error: ", err)
	}

}
