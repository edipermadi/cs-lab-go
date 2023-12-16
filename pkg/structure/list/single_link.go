package list

import "golang.org/x/exp/constraints"

func NewSingleLink[T constraints.Ordered](value T) *SingleLink[T] {
	return &SingleLink[T]{Value: value}
}

type SingleLink[T constraints.Ordered] struct {
	Value T
	Next  *SingleLink[T]
}

func (s *SingleLink[T]) Tail() *SingleLink[T] {
	for node := s; ; {
		if node.Next == nil {
			return node
		}

		node = node.Next
	}
}

func (s *SingleLink[T]) Length() int {
	count := 0
	for node := s; ; {
		count++
		if node.Next == nil {
			break
		}

		node = node.Next
	}

	return count
}

func (s *SingleLink[T]) Append(x *SingleLink[T]) *SingleLink[T] {
	s.Next = x
	return x
}

func (s *SingleLink[T]) Prepend(x *SingleLink[T]) *SingleLink[T] {
	x.Next = s
	return x
}

func (s *SingleLink[T]) Splice(prev *SingleLink[T]) {
	next := s.Next

	// disconnect current from previous
	if prev != nil {
		prev.Next = next
	}

	s.Next = nil
}
