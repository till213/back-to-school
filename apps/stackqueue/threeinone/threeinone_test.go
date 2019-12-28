package main

import (
	"testing"
)

func TestMiddleStack(t *testing.T) {
	var s Stack

	s.init(7)
	if e := s.empty(middle); !e {
		t.Errorf("Excpected empty, got %v", e)
	}
	s.push(middle, 1)
	s.push(middle, 2)
	s.push(middle, 3)
	if e := s.empty(middle); e {
		t.Errorf("Excpected not empty, got %v", e)
	}
	if v := s.pop(middle); v != 3 {
		t.Errorf("Excpected 3, got %v", v)
	}
	if v := s.pop(middle); v != 2 {
		t.Errorf("Excpected 2, got %v", v)
	}
	if v := s.pop(middle); v != 1 {
		t.Errorf("Excpected 1, got %v", v)
	}
}
