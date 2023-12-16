package queue

import (
	"golang.org/x/exp/constraints"
)

func NewSliceQueue[T constraints.Ordered]() *SliceQueue[T] {
	return &SliceQueue[T]{}
}

type SliceQueue[T constraints.Ordered] struct {
	values []T
}

func (s *SliceQueue[T]) Enqueue(value T) {
	s.values = append(s.values, value)
}

func (s *SliceQueue[T]) Dequeue() T {
	var value T

	length := len(s.values)
	if length > 0 {
		value = s.values[0]
		if length == 1 {
			s.values = nil
		} else {
			s.values = s.values[1:]
		}
	}

	return value
}

func (s *SliceQueue[T]) Size() int {
	return len(s.values)
}
