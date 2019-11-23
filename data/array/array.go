package array

import "math/rand"

// Sorted returns an already sorted slice of size n,
// either with ascending or descending order.
func Sorted(n int, ascending bool) []int {
	a := make([]int, n)

	for i := range a {
		if ascending {
			a[i] = i
		} else {
			a[i] = len(a) - i - 1
		}
	}
	return a;
}

// Random returns a slice of size n with random elements [0, n).
func Random(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = rand.Intn(n)
	}
	return a;
}

// IsSorted returns true if slice a is sorted, either
// in ascending or descending order
func IsSorted(a []int, ascending bool) bool {
	var sorted bool
	if len(a) > 1 {
		sorted = true
		for i, n := 1, len(a); sorted && i < n; i++ {
			if ascending {
				sorted = a[i-1] <= a[i]
			} else {
				sorted = a[i-1] >= a[i]
			}
		}
	} else {
		sorted = true
	}
	return sorted
}
