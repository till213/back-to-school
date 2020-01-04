package main

import (
	"container/list"
	"fmt"

	"github.com/till213/back-to-school/data/tree"
)

func cloneList(source *list.List) *list.List {
	result := list.New()
	for e := source.Front(); e != nil; e = e.Next() {
		result.PushBack(e.Value)
	}
	return result
}

func addAll(list1, toAdd *list.List) {
	for e := toAdd.Front(); e != nil; e = e.Next() {
		list1.PushBack(e.Value)
	}
}

func weave(first, second, prefix *list.List, results *[]*list.List) {

	if first.Len() == 0 || second.Len() == 0 {
		result := cloneList(prefix)
		addAll(result, first)
		addAll(result, second)
		*results = append(*results, result)
		return
	}

	// Weave with first list first
	// Temporarily remove first element and append it to the prefix list
	headFirst := first.Front()
	first.Remove(headFirst)
	prefix.PushBack(headFirst.Value)
	weave(first, second, prefix, results)
	// Restore lists
	prefix.Remove(prefix.Back())
	first.PushFront(headFirst.Value)

	// Then with second list
	headSecond := second.Front()
	second.Remove(headSecond)
	prefix.PushBack(headSecond.Value)
	weave(first, second, prefix, results)
	// Restore lists
	prefix.Remove(prefix.Back())
	second.PushFront(headSecond.Value)
}

func bstSequences(n *tree.Node) []*list.List {

	results := make([]*list.List, 0)
	var leftSq, rightSq []*list.List

	if n == nil {
		results = append(results, list.New())
		return results
	}

	leftSq = bstSequences(n.Left)
	rightSq = bstSequences(n.Right)

	prefix := list.New()
	prefix.PushBack(n.Value)
	for _, v1 := range leftSq {
		for _, v2 := range rightSq {
			weaved := make([]*list.List, 0)
			weave(v1, v2, prefix, &weaved)
			results = append(results, weaved...)
		}
	}
	return results
}

func main() {
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := tree.CreateBinarySearchTree(sorted, nil)

	sq := bstSequences((root))
	for i, seq := range sq {
		fmt.Println("Sequence: ", i+1)
		for e := seq.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v, ", e.Value.(int))
		}
		fmt.Println()
	}
}
