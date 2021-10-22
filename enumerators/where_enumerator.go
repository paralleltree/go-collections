package enumerators

type WhereEnumerator[T any] struct {
	src       Enumerator[T]
	predicate func(T) bool
}

func NewWhereEnumerator[T any](src Enumerator[T], predicate func(T) bool) Enumerator[T] {
	return &WhereEnumerator[T]{src, predicate}
}

func (e *WhereEnumerator[T]) Current() T {
	return e.src.Current()
}

func (e *WhereEnumerator[T]) MoveNext() bool {
	for e.src.MoveNext() {
		if e.predicate(e.src.Current()) {
			return true
		}
	}
	return false
}

func (e *WhereEnumerator[T]) Reset() {
	e.src.Reset()
}
