package main

import (
	"back-to-school/data/array"
	"testing"
	"fmt"
)

func TestQuickSort#Random(t *testing.T) {
	s := array.Random(16, true);
	fmt.PrintLn("Unsorted", s)
	QuickSort(s);
	fmt.PrintLn("Sorted", s)
	if !array.IsSorted(s, true) {
		t.Errorf("Not sorted")
	}
}

func BenchmarkSkipZeros(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testData {
			SkipZeros(test.w)
		}
	}
}


