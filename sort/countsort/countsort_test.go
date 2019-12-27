package main

import (
	"fmt"
	"github.com/till213/back-to-school/data/array"
	"testing"
)

const n = 1024

func TestInsertionSortRandom(t *testing.T) {
	arr := array.Random(n)
	sorted := make([]int, len(arr))
	fmt.Println("Unsorted", arr)
	CountSort(arr, sorted, n)
	fmt.Println("Sorted", sorted)
	if !array.IsSorted(sorted, true) {
		t.Errorf("Not sorted")
	}
}

func TestInsertionSortSortedAscending(t *testing.T) {
	arr := array.Sorted(n, true)
	sorted := make([]int, len(arr))
	fmt.Println("Sorted ascending", arr)
	CountSort(arr, sorted, n)
	fmt.Println("Sorted", sorted)
	if !array.IsSorted(sorted, true) {
		t.Errorf("Not sorted")
	}
}

func TestInsertionSortSortedDescending(t *testing.T) {
	arr := array.Sorted(n, false)
	sorted := make([]int, len(arr))
	fmt.Println("Sorted descending", arr)
	CountSort(arr, sorted, n)
	fmt.Println("Sorted", sorted)
	if !array.IsSorted(sorted, true) {
		t.Errorf("Not sorted")
	}
}

func BenchmarkInsertionSortRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := array.Random(n)
		sorted := make([]int, len(arr))
		CountSort(arr, sorted, n)
	}
}

func BenchmarkInsertionSortAscending(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := array.Sorted(n, true)
		sorted := make([]int, len(arr))
		CountSort(arr, sorted, n)
	}
}

func BenchmarkInsertionSortDescending(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := array.Sorted(n, false)
		sorted := make([]int, len(arr))
		CountSort(arr, sorted, n)
	}
}
