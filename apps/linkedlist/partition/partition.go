package main

import (
	"container/list"
	"fmt"
)

func swapPartition(l *list.List, k int) {
	f := l.Front()
	b := l.Back()
	for f != b {
		// Forward, until we find an element >= k
		for f != b && f.Value.(int) < k {
			f = f.Next()
		}
		// Backward, until we find an element < k
		for b != f && b.Value.(int) >= k {
			b = b.Prev()
		}
		// Swap
		f.Value, b.Value = b.Value, f.Value
	}
}

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
	fmt.Println("--- After partition (Naive) ---")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	l2 := list.New()
	l2.PushBack(3)
	l2.PushBack(5)
	l2.PushBack(8)
	l2.PushBack(5)
	l2.PushBack(10)
	l2.PushBack(2)
	l2.PushBack(1)

	fmt.Println("--- Before partition ---")
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	swapPartition(l2, 5)
	fmt.Println("--- After partition (Swap) ---")
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
