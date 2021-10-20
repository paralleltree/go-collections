package enumerators

import (
	"testing"
)

func TestSliceEnumerator(t *testing.T) {
	s := []int{10, 20, 30}
	e := NewSliceEnumerator(s)

	CheckIntEnumerator(t, s, e)
	e.Reset()
	CheckIntEnumerator(t, s, e)
}
