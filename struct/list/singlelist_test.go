package list

import (
	"testing"
)

func TestSingleList_Insert(t *testing.T) {
	const v1, v2 = 42, 1001
	head := new(Element)
	head.val = v1

	e1 := head.Insert(v1)
	if head.next != e1 {
		t.Errorf("Head next (%v) is not e1 (%v)", head.next, e1)
	}

	e2 := e1.Insert(v2)
	if e1.next != e2 {
		t.Errorf("e1 next (%v) is not e2 (%v)", e1.next, e2)
	}

}

func TestSingleList_Remove(t *testing.T) {
	const v1, v2, v3 = 42, 1001, 27
	head := new(Element)
	head.val = v1
	e1 := head.Insert(v2)
	e1.Insert(v3)

	e1.RemoveNext()
	if e1.next != nil {
		t.Errorf("e1 next (%v) is not nil", e1.next)
	}

	head.Remove()
	if head.val != v2 {
		t.Errorf("head's value %v is not %v", head.val, v2)
	}
}
