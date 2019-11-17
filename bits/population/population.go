package main

import (
	"fmt"
	"math/bits"
	"time"
)

func skipZeros(x uint64) int {
	c := 0
	for x != 0 {
		c++
		x = x & (x - 1)
	}
	return c
}

func logBitSum(w uint64) int {
	// Iteration in log2(n), n=64 -> 6 iterations
	mask := [6]uint64{
		0x5555555555555555, // 01010101...01010101
		0x3333333333333333, // 00110011...00110011
		0x0F0F0F0F0F0F0F0F, // 00001111...00001111
		0x00FF00FF00FF00FF, // ...
		0x0000FFFF0000FFFF,
		0x00000000FFFFFFFF, // 00000000...11111111
	}
	var i uint64
	var even, odd uint64

	j := 1
	for i = 0; i < 6; i++ {
		even = w & mask[i]
		// Shift w right 2^i bits: 1, 2, 4, 8, 16, 32
		w >>= j
		// j = 2, 4, 8, 16, 32
		j <<= 1
		odd = w & mask[i]
		w = even + odd
	}
	return int(w)
}

func main() {
	const N = 1e8
	const all = 0xffffffffffffffff
	const half = 0x0f0f0f0f0f0f0f0f
	const sparse = 0x1001000001000001
	const none = 0x0

	start := time.Now()
	c1, c2, c3, c4 := 0, 0, 0, 0
	for i := 0; i < N; i++ {
		c1 = skipZeros(all)
		c2 = skipZeros(half)
		c3 = skipZeros(sparse)
		c4 = skipZeros(none)
	}
	end := time.Now()
	fmt.Printf("SkipZeros: c1=%d c2=%d c3=%d c4=%d elapsed: %v\n",
		c1, c2, c3, c4, end.Sub(start))

	start = time.Now()
	c1, c2, c3, c4 = 0, 0, 0, 0
	for i := 0; i < N; i++ {
		c1 = logBitSum(all)
		c2 = logBitSum(half)
		c3 = logBitSum(sparse)
		c4 = logBitSum(none)
	}
	end = time.Now()
	fmt.Printf("logBitSum: c1=%d c2=%d c3=%d c4=%d elapsed: %v\n",
		c1, c2, c3, c4, end.Sub(start))

	start = time.Now()
	c1, c2, c3, c4 = 0, 0, 0, 0
	for i := 0; i < N; i++ {
		c1 = bits.OnesCount64(all)
		c2 = bits.OnesCount64(half)
		c3 = bits.OnesCount64(sparse)
		c4 = bits.OnesCount64(none)
	}
	end = time.Now()
	fmt.Printf("OnesCount64: c1=%d c2=%d c3=%d c4=%d elapsed: %v\n",
		c1, c2, c3, c4, end.Sub(start))
}
