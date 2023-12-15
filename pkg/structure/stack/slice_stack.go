package stack

import "golang.org/x/exp/constraints"

type SliceStack[T constraints.Ordered] struct {
	values []T
	tos    int
}

func (s *SliceStack[T]) Push(v T) {
	s.values = append(s.values, v)
	if len(s.values) == 1 {
		s.tos = 0
	} else {
		s.tos++
	}
}

func (s *SliceStack[T]) PushPair(x T, y T) {
	s.Push(x)
	s.Push(y)
}

func (s *SliceStack[T]) PopPair() (T, T) {
	y := s.Pop()
	x := s.Pop()
	return x, y
}

func (s *SliceStack[T]) Pop() T {
	v := s.values[s.tos]
	if len(s.values) == 1 {
		s.values = nil
		s.tos = 0
	} else {
		s.values = s.values[0 : len(s.values)-1]
		s.tos--
	}

	return v
}

func (s *SliceStack[T]) Size() int {
	return len(s.values)
}
