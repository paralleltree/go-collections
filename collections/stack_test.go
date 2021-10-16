package collections

import (
	"testing"
)

func checkPeekValue(t *testing.T, testname string, gotValue int, gotOk bool, wantValue int, wantOk bool) {
	if wantOk != gotOk {
		t.Errorf("%s(ok): want %t, got %t", testname, wantOk, gotOk)
	}

	if wantOk {
		return
	}

	if wantValue != gotValue {
		t.Errorf("%s(v): want %v, got %v", testname, wantValue, gotValue)
	}
}

func TestStackOperations(t *testing.T) {
	s := Stack[int]{}

	s.Push(10)
	s.Push(20)

	if c := s.Count(); c != 2 {
		t.Errorf("Count(): want 2, got %d", c)
	}

	v, ok := s.Peek()
	checkPeekValue(t, "TestStackOperatons (1)", v, ok, 20, true)

	v, ok = s.Pop()
	checkPeekValue(t, "TestStackOperatons (2)", v, ok, 20, true)

	v, ok = s.Pop()
	checkPeekValue(t, "TestStackOperatons (3)", v, ok, 10, true)

	// empty stack
	v, ok = s.Pop()
	checkPeekValue(t, "TestStackOperatons (4)", v, ok, 0, false)

	v, ok = s.Peek()
	checkPeekValue(t, "TestStackOperatons (4)", v, ok, 0, false)

	if c := s.Count(); c != 0 {
		t.Errorf("Count(): want 0, got %d", c)
	}
}
