package main

import "fmt"

type elem struct {
	val  int
	next *elem
}

func remove(e *elem) {
	if e.next != nil {
		e.val = e.next.val
		e.next = e.next.next
	}
}

func main() {
	var r *elem
	head := new(elem)
	head.val = 0
	for i := 10; i > 0; i-- {
		e := new(elem)
		e.val = i
		e.next = head.next
		head.next = e
		if i == 5 {
			r = e
		}
	}
	fmt.Println("--- Original list ---")
	for e := head; e != nil; e = e.next {
		fmt.Println("Value: ", e.val)
	}
	remove(r)
	fmt.Println("--- Removed element ---")
	for e := head; e != nil; e = e.next {
		fmt.Println("Value: ", e.val)
	}
}
