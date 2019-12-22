package main

import (
	"fmt"
	"github.com/till213/back-to-school/data/array"
	"testing"
)

const n = 1024

func TestInsertionSortRandom(t *testing.T) {
	s := array.Random(n)
	fmt.Println("Unsorted", s)
	InsertionSort(s)
	fmt.Println("Sorted", s)
	if !array.IsSorted(s, true) {
		t.Errorf("Not sorted")
	}
}

func TestInsertionSortSortedAscending(t *testing.T) {
	s := array.Sorted(n, true)
	fmt.Println("Sorted ascending", s)
	InsertionSort(s)
	fmt.Println("Sorted", s)
	if !array.IsSorted(s, true) {
		t.Errorf("Not sorted")
	}
}

func TestInsertionSortSortedDescending(t *testing.T) {
	s := array.Sorted(n, false)
	fmt.Println("Sorted descending", s)
	InsertionSort(s)
	fmt.Println("Sorted", s)
	if !array.IsSorted(s, true) {
		t.Errorf("Not sorted")
	}
}

func BenchmarkInsertionSortRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := array.Random(n)
		InsertionSort(s)
	}
}

func BenchmarkInsertionSortAscending(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := array.Sorted(n, true)
		InsertionSort(s)
	}
}

func BenchmarkInsertionSortDescending(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := array.Sorted(n, false)
		InsertionSort(s)
	}
}
