package enumerators

import (
	"testing"
)

func TestNewRangeEnumeratorWithCount(t *testing.T) {
	count := 5
	e := NewRangeEnumeratorWithCount(count)
	for i := 0; i < count; i++ {
		if m := e.MoveNext(); !m {
			t.Errorf("MoveNext(): want %t, got %t", true, m)
		}

		if v := e.Current(); v != i {
			t.Errorf("Current(): want %d, got %d", i, v)
		}
	}

	if m := e.MoveNext(); m {
		t.Errorf("MoveNext(): want %t, got %t", false, m)
	}
}

func TestNewRangeEnumeratorWithStart(t *testing.T) {
	start := 10
	e := NewRangeEnumerator(RangeWithStart(start))

	for i := start; i < start+10; i++ {
		if m := e.MoveNext(); !m {
			t.Errorf("MoveNext(): want %t, got %t", true, m)
		}

		if v := e.Current(); v != i {
			t.Errorf("Current(): want %d, got %d", i, v)
		}
	}

	// infinite sequence
	if m := e.MoveNext(); !m {
		t.Errorf("MoveNext(): want %t, got %t", true, m)
	}
}

func TestNewRangeEnumeratorWithStartAndEnd(t *testing.T) {
	start := 10
	end := 20
	// step := 1
	e := NewRangeEnumerator(RangeWithStart(start), RangeWithEnd(end))

	for i := start; i < end; i++ {
		if m := e.MoveNext(); !m {
			t.Errorf("MoveNext(): want %t, got %t", true, m)
		}

		if v := e.Current(); v != i {
			t.Errorf("Current(): want %d, got %d", i, v)
		}
	}

	if m := e.MoveNext(); m {
		t.Errorf("MoveNext(): want %t, got %t", false, m)
	}
}

func TestNewRangeEnumeratorWithStartAndEndAndStep(t *testing.T) {
	start := 10
	end := 20
	step := 3
	e := NewRangeEnumerator(RangeWithStart(start), RangeWithEnd(end), RangeWithStep(step))

	for i := start; i < end; i += step {
		if m := e.MoveNext(); !m {
			t.Errorf("MoveNext(): want %t, got %t", true, m)
		}

		if v := e.Current(); v != i {
			t.Errorf("Current(): want %d, got %d", i, v)
		}
	}

	if m := e.MoveNext(); m {
		t.Errorf("MoveNext(): want %t, got %t", false, m)
	}
}
