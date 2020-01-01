package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// The maximum number is 999 * 1000 * 1001 = 999999000
// There follows that the largest Fibonacci number that
// can be calculated is the 49th number, which is
// 7778742049 < 999999000
const m1 = 999
const m2 = 1000
const m3 = 1001

// PoorNumber implements the Poor Man's Large Numbers
// Case in point: we assume 16bit integers (only)
type PoorNumber [3]int16

// Init initialises all elements with value v
func (n *PoorNumber) Init(v int16) {
	for i := 0; i < 3; i++ {
		n[i] = v
	}
}

// http://www.inf.fh-flensburg.de/lang/krypto/algo/euklid.htm#section2
func extgcd(a, b int) (ar, ur, vr int) {
	u, v, s, t := 1, 0, 0, 1
	for b != 0 {
		q := a / b
		a, b = b, a-q*b
		u, s = s, u-q*s
		v, t = t, v-q*t
	}
	return a, u, v
}

// http://www.inf.fh-flensburg.de/lang/krypto/algo/chinese-remainder.htm
func chineseRemainder(nn, rr []int16) (m, x int) {
	if len(nn) == 1 {
		return int(nn[0]), int(rr[0])
	} else {
		k := len(nn) / 2
		m, a := chineseRemainder(nn[:k], rr[:k])
		n, b := chineseRemainder(nn[k:], rr[k:])
		_, u, _ := extgcd(m, n)
		x = (b-a)*u%n*m + a
		return m * n, x
	}
}

// IntVal returns the PoorNumber value as integer, by
// applying the extended Euklidean algorithm together
// with the Chinese Remainder theorem
func (n *PoorNumber) IntVal() int {
	mm := [3]int16{m1, m2, m3}
	m, x := chineseRemainder(mm[:], n[:])
	if x < 0 {
		x += m
	}
	return x
}

// Fibonacci calculates the n-th Fibonacci number,
// using Poor Man's Large Integers
func Fibonacci(n int16) PoorNumber {
	var r PoorNumber
	if n <= 1 {
		r.Init(n)
	} else {
		var p, q PoorNumber
		p.Init(0)
		q.Init(1)
		for i := 2; i <= int(n); i++ {
			for j := 0; j < 3; j++ {
				r[j] = int16(int(p[j]+q[j]) % int(m1+j))
			}
			p = q
			q = r
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
		f := Fibonacci(int16(n))
		fmt.Printf("%v-th Fibonacci number: %v (int: %v)\n", n, f, f.IntVal())
	} else {
		fmt.Println("Parse error: ", err)
	}
}
