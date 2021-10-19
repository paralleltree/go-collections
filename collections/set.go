package collections

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{make(map[T]struct{})}
}

func (s *Set[T]) Add(item T) {
	s.m[item] = struct{}{}
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set[T]) Remove(item T) {
	delete(s.m, item)
}

func (a *Set[T]) Intersect(b *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item, _ := range a.m {
		if b.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

func (a *Set[T]) Union(b *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item, _ := range a.m {
		result.Add(item)
	}
	for item, _ := range b.m {
		result.Add(item)
	}
	return result
}

func (a *Set[T]) Except(b *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item, _ := range a.m {
		if !b.Contains(item) {
			result.Add(item)
		}
	}
	return result
}
