package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func toh(n int, x, y, z *stack.Stack) {
	if n > 0 {
		toh(n-1, x, z, y)
		e := x.Pop()
		z.Push(e)
		toh(n-1, y, x, z)
	}
}

func main() {
	x := stack.New()
	y := stack.New()
	z := stack.New()

	x.Push(5)
	x.Push(4)
	x.Push(3)
	x.Push(2)
	x.Push(1)

	toh(5, x, y, z)
	fmt.Println("Tower X")
	fmt.Printf("Len %d\n", x.Len())
	for x.Len() > 0 {
		fmt.Printf("%v ", x.Pop())
	}
	fmt.Println("")

	fmt.Println("Tower Y")
	fmt.Printf("Len %d\n", y.Len())
	for y.Len() > 0 {
		fmt.Printf("%v ", y.Pop())
	}
	fmt.Println("")

	fmt.Println("Tower Z")
	fmt.Printf("Len %d\n", z.Len())
	for z.Len() > 0 {
		fmt.Printf("%v ", z.Pop())
	}
	fmt.Println("")
}
