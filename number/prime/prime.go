package main

import (
	"math"
	"fmt"
	"os"
	"strconv"
	"errors"
)

// Prime number sieve of Eratosthenes
func sieve(n int) {
	var sieve []bool
	var p, sqrtn, i int

	sieve = make([]bool, n)
	for i := 0; i < n; i++ {
		sieve[i] = true
	}
	sqrtn = (int)(math.Floor(math.Sqrt((float64)(n))))
	for p = 2; p <= sqrtn;  {		
		for i = p * p; i <= n;  i = i + p {
			sieve[i-1] = false 
		}
		for {
			p++
			if sieve[p-1] {
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		if sieve[i] {
			fmt.Printf("%d ", i+1)
		}
	}
	fmt.Println("")

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
		sieve(n)
	} else {
		fmt.Print("Parse error: ", err)
	}
}
