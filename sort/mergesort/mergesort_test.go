package main

import (
	"fmt"
	"github.com/till213/back-to-school/data/array"
	"testing"
)

const n = 1024

func TestMergeSortRandom(t *testing.T) {
	s := array.Random(n)
	fmt.Println("Unsorted", s)
	MergeSort(s, 0, len(s)-1)
	fmt.Println("Sorted", s)
	if !array.IsSorted(s, true) {
		t.Errorf("Not sorted")
	}
}

func TestMergeSortSortedAscending(t *testing.T) {
	s := array.Sorted(n, true)
	fmt.Println("Sorted ascending", s)
	MergeSort(s, 0, len(s)-1)
	fmt.Println("Sorted", s)
	if !array.IsSorted(s, true) {
		t.Errorf("Not sorted")
	}
}

func TestMergeSortSortedDescending(t *testing.T) {
	s := array.Sorted(n, false)
	fmt.Println("Sorted descending", s)
	MergeSort(s, 0, len(s)-1)
	fmt.Println("Sorted", s)
	if !array.IsSorted(s, true) {
		t.Errorf("Not sorted")
	}
}

func BenchmarkMergeSortRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := array.Random(n)
		MergeSort(s, 0, len(s)-1)
	}
}

func BenchmarkMergeSortAscending(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := array.Sorted(n, true)
		MergeSort(s, 0, len(s)-1)
	}
}

func BenchmarkMergeSortDescending(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := array.Sorted(n, false)
		MergeSort(s, 0, len(s)-1)
	}
}
