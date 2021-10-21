package enumerators

import (
	"testing"
)

func TestNewSelectEnumerator(t *testing.T) {
	src := []int{0, 1, 2, 3}
	selector := func(v int) int { return v * 2 }
	e := NewSliceEnumerator(src)
	e = NewSelectEnumerator(e, selector)

	wants := []int{0, 2, 4, 6}

	CheckIntEnumerator(t, wants, e)
	e.Reset()
	CheckIntEnumerator(t, wants, e)
}
