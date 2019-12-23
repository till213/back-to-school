package jumpball

import (
	"testing"
)

func TestJumpBallCanStop(t *testing.T) {
	path := []bool{true, false, true, true, true, false, true, true, false, true, true}
	canStop := CanStop(path, 0, 2)
	if !canStop {
		t.Errorf("Expected true, received %v", canStop)
	}
}

func TestJumpBallCannotStop(t *testing.T) {
	path := []bool{true, false, true, true, true, false, true, true, false, true, true}
	canStop := CanStop(path, 0, 6)
	if canStop {
		t.Errorf("Expected false, received %v", canStop)
	}
}

func TestJumpBallCanStopWithMemo(t *testing.T) {
	path := []bool{true, false, true, true, true, false, true, true, false, true, true}
	memo := make(map[state]bool)
	canStop := CanStopWithMemo(path, 0, 2, memo)
	if !canStop {
		t.Errorf("Expected true, received %v", canStop)
	}
}

func TestJumpBallCannotStopWithMemo(t *testing.T) {
	path := []bool{true, false, true, true, true, false, true, true, false, true, true}
	memo := make(map[state]bool)
	canStop := CanStopWithMemo(path, 0, 6, memo)
	if canStop {
		t.Errorf("Expected false, received %v", canStop)
	}
}
