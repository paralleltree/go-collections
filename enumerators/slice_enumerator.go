package enumerators

type SliceEnumerator[T any] struct {
	s     []T
	index int
}

func NewSliceEnumerator[T any](src []T) Enumerator[T] {
	return &SliceEnumerator[T]{src, -1}
}

func (e *SliceEnumerator[T]) Current() T {
	return e.s[e.index]
}

func (e *SliceEnumerator[T]) MoveNext() bool {
	canMove := e.index < len(e.s)-1
	if canMove {
		e.index += 1
	}
	return canMove
}

func (e *SliceEnumerator[T]) Reset() {
	e.index = -1
}
