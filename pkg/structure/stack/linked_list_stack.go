package stack

import (
	"github.com/edipermadi/cs-lab-go/pkg/structure/list"
	"golang.org/x/exp/constraints"
)

func NewLinkedListStack[T constraints.Ordered]() *LinkedListStack[T] {
	return &LinkedListStack[T]{}
}

type LinkedListStack[T constraints.Ordered] struct {
	head *list.SingleLink[T]
	size int
}

func (s *LinkedListStack[T]) Push(value T) {
	entry := list.NewSingleLink(value)
	entry.Next = s.head
	s.head = entry
	s.size++
}

func (s *LinkedListStack[T]) Pop() T {
	var value T
	if s.size > 0 {
		head := s.head
		value = head.Value
		s.head = head.Next

		// detach and decrease value
		head.Next = nil
		s.size--
	}

	return value
}

func (s *LinkedListStack[T]) PushPair(x T, y T) {
	s.Push(x)
	s.Push(y)
}

func (s *LinkedListStack[T]) PopPair() (T, T) {
	y := s.Pop()
	x := s.Pop()
	return x, y
}

func (s *LinkedListStack[T]) Size() int {
	return s.size
}
