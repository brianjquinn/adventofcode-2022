package day5

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(t T) {
	s.items = append(s.items, t)
}

func (s *Stack[T]) Pop() T {
	n := len(s.items)
	t := s.items[n-1]
	var zero T
	s.items[n-1] = zero
	s.items = s.items[:n-1]
	return t
}
