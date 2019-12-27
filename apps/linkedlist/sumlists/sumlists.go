package main

import (
	"container/list"
	"fmt"
)

func listToInt(l *list.List) int {
	var n int
	if l.Len() == 0 {
		return 0
	}
	m := 1
	for e := l.Front(); e != nil; e = e.Next() {
		n += e.Value.(int) * m
		m *= 10
	}
	return n
}

func intToLlist(i int) *list.List {
	r := list.New()
	for i > 0 {
		rem := i % 10
		r.PushBack(rem)
		i /= 10
	}
	return r
}

func sumLists(l1, l2 *list.List) *list.List {
 	a := listToInt(l1)
	b := listToInt(l2)
	sum := a + b
	return intToLlist(sum)
}

func main() {
	l1 := list.New()
	l1.PushBack(7)
	l1.PushBack(1)
	l1.PushBack(6)

	l2 := list.New()
	l2.PushBack(5)
	l2.PushBack(9)
	l2.PushBack(2)

	r := sumLists(l1, l2)
	
	sum := listToInt(r)
	fmt.Println("Sum:", sum)
}