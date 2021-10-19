package collections

import (
	"testing"
)

func checkIntSetContains(t *testing.T, testname string, s *Set[int], item int, want bool) {
	if got := s.Contains(item); want != got {
		t.Errorf("%s: Contains(%v): want %t, got %t", testname, item, want, got)
	}
}

func TestSetOperations(t *testing.T) {
	s := NewSet[int]()

	checkIntSetContains(t, "TestSetOperations (1)", s, 10, false)

	s.Add(10)

	checkIntSetContains(t, "TestSetOperations (2)", s, 10, true)

	s.Remove(10)

	checkIntSetContains(t, "TestSetOperations (3)", s, 10, false)
}

func TestSet_Intersect(t *testing.T) {
	a := NewSet[int]()
	a.Add(10)
	a.Add(20)
	b := NewSet[int]()
	b.Add(10)

	result := a.Intersect(b)
	checkIntSetContains(t, "Intersect", result, 10, true)
	checkIntSetContains(t, "Intersect", result, 20, false)
}

func TestSet_Union(t *testing.T) {
	a := NewSet[int]()
	a.Add(10)
	a.Add(20)
	b := NewSet[int]()
	b.Add(10)

	result := a.Union(b)
	checkIntSetContains(t, "Union", result, 10, true)
	checkIntSetContains(t, "Union", result, 20, true)
}

func TestSet_Except(t *testing.T) {
	a := NewSet[int]()
	a.Add(10)
	a.Add(20)
	b := NewSet[int]()
	b.Add(10)

	result := a.Except(b)
	checkIntSetContains(t, "Except", result, 10, false)
	checkIntSetContains(t, "Except", result, 20, true)
}
