package orderedhashset

import "testing"

func TestContains(t *testing.T) {
	value := 42
	s := New()
	s.Add(value)
	act := s.Contains(value)

	if !act {
		t.Errorf("TestContains(%d) = %v; want %v", value, act, true)
	}
	notExistValue := 1001
	act = s.Contains(notExistValue)

	if act {
		t.Errorf("TestContains(%d) = %v; want %v", notExistValue, act, false)
	}
}

func TestKeys(t *testing.T) {
	s := New()

	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		s.Add(v)
	}

	keys := s.Keys()
	for i, k := range keys {
		if k != values[i] {
			t.Errorf("TestKeys actual: %v; expected %v", k, values[i])
		}
	}
}

func TestReinsertKey(t *testing.T) {
	s := New()

	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		s.Add(v)
	}

	s.Add(1)

	keys := s.Keys()
	for i, k := range keys {
		if k != values[i] {
			t.Errorf("TestKeys actual: %v; expected %v", k, values[i])
		}
	}
}

func TestReinsertKeyWithRemoval(t *testing.T) {
	s := New()

	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		s.Add(v)
	}

	s.Remove(1)
	s.Add(1)

	keys := s.Keys()
	expected := []int{2, 3, 4, 5, 1}
	for i, k := range keys {
		if k != expected[i] {
			t.Errorf("TestKeys actual: %v; expected %v", k, expected[i])
		}
	}
}

func TestIterator(t *testing.T) {
	s := New()

	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		s.Add(v)
	}

	i := 0
	for iter := s.Iterator(); iter != nil; iter = iter.Next() {
		if v := iter.Value().(int); v != values[i] {
			t.Errorf("TestIterator actual: %v; expected %v", v, values[i])
		}
		i++
	}
}
