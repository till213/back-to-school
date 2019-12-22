package list

import (
	"testing"
)

func TestSingleList_Insert(t *testing.T) {
	const v1, v2 = 42, 1001
	head := new(Element)
	head.Val = v1

	e1 := head.Insert(v1)
	if head.Next != e1 {
		t.Errorf("Head Next (%v) is not e1 (%v)", head.Next, e1)
	}

	e2 := e1.Insert(v2)
	if e1.Next != e2 {
		t.Errorf("e1 Next (%v) is not e2 (%v)", e1.Next, e2)
	}

}

func TestSingleList_Remove(t *testing.T) {
	const v1, v2, v3 = 42, 1001, 27
	head := new(Element)
	head.Val = v1
	e1 := head.Insert(v2)
	e1.Insert(v3)

	e1.RemoveNext()
	if e1.Next != nil {
		t.Errorf("e1 Next (%v) is not nil", e1.Next)
	}

	head.Remove()
	if head.Val != v2 {
		t.Errorf("head's Value %v is not %v", head.Val, v2)
	}
}
