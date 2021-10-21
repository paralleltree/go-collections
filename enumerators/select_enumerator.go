package enumerators

type SelectEnumerator[TSource any, TResult any] struct {
	src      Enumerator[TSource]
	selector func(TSource) TResult
}

func NewSelectEnumerator[TSource any, TResult any](src Enumerator[TSource], selector func(TSource) TResult) Enumerator[TResult] {
	return &SelectEnumerator[TSource, TResult]{src, selector}
}

func (e *SelectEnumerator[TSource, TResult]) Current() TResult {
	return e.selector(e.src.Current())
}

func (e *SelectEnumerator[TSource, TResult]) MoveNext() bool {
	return e.src.MoveNext()
}

func (e *SelectEnumerator[TSource, TResult]) Reset() {
	e.src.Reset()
}
