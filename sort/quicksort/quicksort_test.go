package main

import (
	"fmt"
	"testing"

	"github.com/till213/back-to-school/data/array"
)

const n = 1024

func TestQuickSortRandom(t *testing.T) {
	s := array.Random(n)
	fmt.Println("Unsorted", s)
	QuickSort(s)
	fmt.Println("Sorted", s)
	if !array.IsSorted(s, true) {
		t.Errorf("Not sorted")
	}
}

func TestQuickSortSortedAscending(t *testing.T) {
	s := array.Sorted(n, true)
	fmt.Println("Sorted ascending", s)
	QuickSort(s)
	fmt.Println("Sorted", s)
	if !array.IsSorted(s, true) {
		t.Errorf("Not sorted")
	}
}

func TestQuickSortSortedDescending(t *testing.T) {
	s := array.Sorted(n, false)
	fmt.Println("Sorted descending", s)
	QuickSort(s)
	fmt.Println("Sorted", s)
	if !array.IsSorted(s, true) {
		t.Errorf("Not sorted")
	}
}

func BenchmarkQuickSortRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := array.Random(n)
		QuickSort(s)
	}
}

func BenchmarkQuickSortAscending(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := array.Sorted(n, true)
		QuickSort(s)
	}
}

func BenchmarkQuickSortDescending(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := array.Sorted(n, false)
		QuickSort(s)
	}
}
