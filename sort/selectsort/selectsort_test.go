package main

import (
	"fmt"
	"testing"
	"github.com/till213/back-to-school/data/array"
)

const n = 1024

func TestSelectionSortRandom(t *testing.T) {
	s := array.Random(n)
	fmt.Println("Unsorted", s)
	SelectionSort(s)
	fmt.Println("Sorted", s)
	if !array.IsSorted(s, true) {
		t.Errorf("Not sorted")
	}
}

func TestSelectionSortSortedAscending(t *testing.T) {
	s := array.Sorted(n, true)
	fmt.Println("Sorted ascending", s)
	SelectionSort(s)
	fmt.Println("Sorted", s)
	if !array.IsSorted(s, true) {
		t.Errorf("Not sorted")
	}
}

func TestSelectionSortSortedDescending(t *testing.T) {
	s := array.Sorted(n, false)
	fmt.Println("Sorted descending", s)
	SelectionSort(s)
	fmt.Println("Sorted", s)
	if !array.IsSorted(s, true) {
		t.Errorf("Not sorted")
	}
}

func BenchmarkSelectionSortRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := array.Random(n)
		SelectionSort(s)
	}
}

func BenchmarkSelectionSortAscending(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := array.Sorted(n, true)
		SelectionSort(s)
	}
}

func BenchmarkSelectionSortDescending(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := array.Sorted(n, false)
		SelectionSort(s)
	}
}
