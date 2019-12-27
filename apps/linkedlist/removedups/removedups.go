package main

import (
	"container/list"
	"fmt"
)

func removeDuplicates(l *list.List) {
	elems := make(map[int]bool)
	e := l.Front()
	for e != nil  {
		if elems[e.Value.(int)] {
			d := e
			e = e.Next()
			l.Remove(d)
		} else {
			elems[e.Value.(int)] = true
			e = e.Next()
		}
	}
}

func removeDuplicatesWithoutAuxStruct(l *list.List) {
	e := l.Front()
	for e != nil {
		f := e.Next()
		for f != nil {
			if e.Value == f.Value {
				d := f
				f = f.Next()
				l.Remove(d)
			} else {
				f = f.Next()
			}
		}
		e = e.Next()
	}
}

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(3)
	l.PushBack(2)
	l.PushBack(5)
	l.PushBack(1)
	l.PushBack(2)

	fmt.Println("--- With duplicates ---")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	removeDuplicates(l)
	fmt.Println("--- Without duplicates ---")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	l2 := list.New()
	l2.PushBack(1)
	l2.PushBack(3)
	l2.PushBack(2)
	l2.PushBack(5)
	l2.PushBack(1)
	l2.PushBack(2)

	fmt.Println("--- With duplicates ---")
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	removeDuplicatesWithoutAuxStruct(l2)
	fmt.Println("--- Without duplicates ---")
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
