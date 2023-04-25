package set

import (
	"testing"
)

func TestIntSetVersousMapSet(t *testing.T) {

	iSet := IntSet{}
	mSet := MapSet{map[int]bool{}}

	iSet.Add(1)
	mSet.Add(1)

	if mSet.Has(1) != iSet.Has(1) {
		t.Errorf("mSet.Has(1) != iSet.Has(1); want mSet.Has(1) == iSet.Has(1)")
	}

	if mSet.Len() != iSet.Len() {

		t.Errorf("mSet.Len() != iSet.Len(); want mSet.Len() == iSet.Len()")
	}

	if mSet.String() != iSet.String() {
		t.Errorf("mSet.String() != iSet.String(); want mSet.String() == iSet.String()")
	}

	mSet.Remove(1)
	iSet.Remove(1)

	if mSet.Has(1) != iSet.Has(1) {
		t.Errorf("mSet.Has(1) != iSet.Has(1); want mSet.Has(1) == iSet.Has(1)")
	}

	if mSet.Len() != iSet.Len() {
		t.Errorf("mSet.Len() != iSet.Len(); want mSet.Len() == iSet.Len()")
	}

	if mSet.String() != iSet.String() {
		t.Errorf("mSet.String() != iSet.String(); want mSet.String() == iSet.String()")
	}

}
