package maxsum

import (
	"testing"
)

func TestNaiveMaxSum(t *testing.T) {
	const k = 4
	arr := [10]int{1, 4, 2, 10, 2, 3, 1, 0, 20}
	max := NaiveMaxSum(arr[:], k)
	if max != 24 {
		t.Errorf("Expected 24, received %v", max)
	}
}

func TestSlidingMaxSum(t *testing.T) {
	const k = 4
	arr := [10]int{1, 4, 2, 10, 2, 3, 1, 0, 20}
	max := SlidingMaxSum(arr[:], k)
	if max != 24 {
		t.Errorf("Expected 24, received %v", max)
	}
}


