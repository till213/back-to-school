package randomkeymap

import "testing"

func TestGet(t *testing.T) {
	key := 1001
	value := 42
	m := New()
	m.Insert(key, value)
	act, ok := m.Get(key)

	if !ok {
		t.Errorf("TestGet: received ok: %v, expected: %v", ok, true)
	}
	if act != value {
		t.Errorf("TestGet: received %v, expected: %v", act, value)
	}
	notExistKey := 2001
	act, ok = m.Get(notExistKey)

	if ok {
		t.Errorf("TestGet: received ok: %v, expected: %v", ok, false)
	}
	if act != nil {
		t.Errorf("TestGet: received %v, expected: %v", act, nil)
	}
}
