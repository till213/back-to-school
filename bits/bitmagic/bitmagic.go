package main

import "fmt"

type bits uint8

const (
	f0 bits = 1 << iota
	f1
	f2
	maxnegbit = -1 << 31
	maxbit    = 1 << 31
)

func atMostOneBit(bits int) bool {
	return bits&(bits-1) == 0
}

func set(b, flag bits) bits    { return b | flag }
func clear(b, flag bits) bits  { return b &^ flag }
func toggle(b, flag bits) bits { return b ^ flag }
func has(b, flag bits) bool    { return b&flag != 0 }

func main() {
	fmt.Printf("Max neg Bit: %b\nMax pos bit:  %b \n", maxnegbit, maxbit)
	fmt.Printf("%b at most one bit: %v\n", 32, atMostOneBit(32))
	fmt.Printf("%b at most one bit: %v\n", 32+16, atMostOneBit(32+16))

	var b bits
	b = set(b, f0)
	b = toggle(b, f2)
	for i, flag := range []bits{f0, f1, f2} {
		fmt.Println(i, has(b, flag))
	}
}
