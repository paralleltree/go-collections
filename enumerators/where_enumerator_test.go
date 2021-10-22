package enumerators

import (
	"testing"
)

func TestNewWhereEnumerator(t *testing.T) {
	src := []int{0, 1, 2, 3}
	predicate := func(v int) bool { return v%2 == 0 }
	e := NewSliceEnumerator(src)
	e = NewWhereEnumerator(e, predicate)

	wants := []int{0, 2}

	CheckIntEnumerator(t, wants, e)
	e.Reset()
	CheckIntEnumerator(t, wants, e)
}
