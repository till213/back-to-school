package main

import (
	"container/list"
	"fmt"
	"errors"
	"strconv"
	"os"
)

func kToLast(l *list.List, k int) *list.List {
	var i int
	var e *list.Element

	len := l.Len()
	if k > len {
		return nil
	}

	for e = l.Front(); i < len-k && e != nil; e = e.Next() {
		i++
	}

	r := list.New()
	for e != nil {
		r.PushBack(e.Value)
		e = e.Next()
	}
	return r
}

func main() {

	var k int
	var err error
	if len(os.Args) > 1 {
		k, err = strconv.Atoi(os.Args[1])
	} else {
		err = errors.New("No arg")
	}
	if err == nil {
		l := list.New()
		l.PushBack(1)
		l.PushBack(3)
		l.PushBack(2)
		l.PushBack(5)
		l.PushBack(1)
		l.PushBack(2)

		fmt.Println("--- Original ---")
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
		r := kToLast(l, k)
		if r != nil {
			fmt.Printf("--- %v last values ---\n", k)
			for e := r.Front(); e != nil; e = e.Next() {
				fmt.Println(e.Value)
			}
		}
	} else {
		fmt.Println("Parse error: ", err)
	}
}