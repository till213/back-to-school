package distinctelem

import (
	"testing"
)

func TestDistinctElements1(t *testing.T) {

	const k = 2
	arr := []int{2, 1, 2, 1, 6}

	count := DistinctElements(arr[:], k)
	if count != 7 {
		t.Errorf("Expected 7, received %v", count)
	}
}

func TestDistinctElements2(t *testing.T) {

	const k = 1
	arr := []int{1, 2, 3, 4, 5}

	count := DistinctElements(arr[:], k)
	if count != 5 {
		t.Errorf("Expected 5, received %v", count)
	}
}
