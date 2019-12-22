package main

import (
	"fmt"
	"github.com/till213/back-to-school/struct/list"
)

func visit(e *list.Element) {
	fmt.Println("Element value: ", e.Val)
}

// Recursively iterates the list and visits each element
// "forward and backward"
func traverseRecursive(e *list.Element) {
	if e != nil {
		visit(e)
		traverseRecursive(e.Next)
		visit(e)
	}
}

// Iteratively iterates the list and visits each element
// "forward and backward"
func traverseIterative(p *list.Element) {
	var o, q *list.Element
	for i := 1; i <= 2; i++ {
		for p != nil {
			visit(p)
			q = p.Next
			p.Next = o
			// the fork advances
			o, p = p, q
		}
		p = o
	}
}

func createList() *list.Element {
	var e *list.Element
	head := new(list.Element)
	head.Val = -1
	e = head
	for i := 0; i < 10; i++ {
		e = e.Insert(i)
	}
	return head
}

func main() {
	head := createList()
	fmt.Println("-- Recursive visit --")
	traverseRecursive(head)
	fmt.Println("-- Iterative visit --")
	traverseIterative(head)
}
