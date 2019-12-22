package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Walker starts the tree walking by calling the recursive function
// Walk and then closes the channel ch.
func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walker(t1, ch1)
	go Walker(t2, ch2)
	same := true
	for same {
		i1, ok1 := <-ch1
		i2, ok2 := <-ch2
		same = ok1 == ok2
		same = same && i1 == i2
		if !ok1 {
			break
		}
	}
	return same
}

func main() {
	t1 := tree.New(1)
	ch := make(chan int)
	go Walker(t1, ch)
	for i := range ch {
		fmt.Println(i)
	}
	t2 := tree.New(1)
	same := Same(t1, t2)
	fmt.Println("Same: ", same)
	t2 = tree.New(2)
	same = Same(t1, t2)
	fmt.Println("Same: ", same)
}
