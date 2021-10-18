package enumerators

type Enumerator[T any] interface {
	Current() T
	MoveNext() bool
	Reset()
}
