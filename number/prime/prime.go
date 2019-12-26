package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/yourbasic/bit"
)

// Prime number sieve of Eratosthenes
func sieve(n int) {
	var sieve []bool
	var p, sqrtn, i int

	sieve = make([]bool, n)
	for i := 1; i < n; i++ {
		sieve[i] = true
	}
	sqrtn = int(math.Floor(math.Sqrt(float64(n))))
	for p = 2; p <= sqrtn; {
		for i = p * p; i <= n; i = i + p {
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

// Prime number sieve of Eratosthenes
func bitsSieve(n int) {

	sieve := bit.New().AddRange(2, n)
	sqrtN := int(math.Sqrt(float64(n)))
	for p := 2; p <= sqrtN; p = sieve.Next(p) {
		for k := p * p; k < n; k += p {
			sieve.Delete(k)
		}
	}
	fmt.Println(sieve)
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
		bitsSieve(n)
	} else {
		fmt.Println("Parse error: ", err)
	}
}
