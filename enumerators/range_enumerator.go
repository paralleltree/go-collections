package enumerators

type RangeEnumerator struct {
	start   int
	end     int
	step    int
	current int
	finite  bool
}

type RangeEnumeratorOption func(*RangeEnumerator)

func RangeWithStart(start int) RangeEnumeratorOption {
	return func(e *RangeEnumerator) {
		e.start = start
	}
}

func RangeWithEnd(end int) RangeEnumeratorOption {
	return func(e *RangeEnumerator) {
		e.end = end
		e.finite = true
	}
}

func RangeWithStep(step int) RangeEnumeratorOption {
	return func(e *RangeEnumerator) {
		e.step = step
	}
}

func NewRangeEnumeratorWithCount(count int) Enumerator[int] {
	return &RangeEnumerator{0, count, 1, -1, true}
}

func NewRangeEnumerator(opts ...RangeEnumeratorOption) Enumerator[int] {
	e := RangeEnumerator{0, 0, 1, -1, false}
	for _, o := range opts {
		o(&e)
	}
	e.Reset()
	return &e
}

func (e *RangeEnumerator) Current() int {
	return e.current
}

func (e *RangeEnumerator) MoveNext() bool {
	if e.finite && e.end-e.step <= e.current {
		return false
	}
	e.current += e.step
	return true
}

func (e *RangeEnumerator) Reset() {
	e.current = e.start - e.step
}
