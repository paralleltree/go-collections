package collections

type Stack[T any] struct {
	buffer []T
}

func (s *Stack[T]) Push(item T) {
	s.buffer = append(s.buffer, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	v, ok := s.Peek()
	if ok {
		s.buffer = s.buffer[0 : len(s.buffer)-1]
	}
	return v, ok
}

func (s *Stack[T]) Peek() (result T, ok bool) {
	if len(s.buffer) == 0 {
		ok = false
		return
	}
	result = s.buffer[len(s.buffer)-1]
	ok = true
	return
}

func (s *Stack[T]) Count() int {
	return len(s.buffer)
}
