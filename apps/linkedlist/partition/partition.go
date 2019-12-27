package main

import (
	"container/list"
	"fmt"
)

func naivePartition(l *list.List, k int) {
	e := l.Front()
	for e != nil {
		if e.Value.(int) < k {
			d := e
			e = e.Next()
			l.PushFront(d.Value)
			l.Remove(d)
		} else {
			e = e.Next()
		}
	}
}

func main() {
	l := list.New()
	l.PushBack(5)
	l.PushBack(3)
	l.PushBack(5)
	l.PushBack(8)
	l.PushBack(5)
	l.PushBack(10)
	l.PushBack(2)
	l.PushBack(1)

	fmt.Println("--- Before partition ---")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	naivePartition(l, 5)
	fmt.Println("--- After partition ---")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
