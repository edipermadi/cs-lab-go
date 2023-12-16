package queue

import (
	"github.com/edipermadi/cs-lab-go/pkg/structure/list"
	"golang.org/x/exp/constraints"
)

func NewLinkedListQueue[T constraints.Ordered]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{}
}

type LinkedListQueue[T constraints.Ordered] struct {
	head *list.SingleLink[T]
	tail *list.SingleLink[T]
	size int
}

func (q *LinkedListQueue[T]) Enqueue(value T) {
	entry := list.NewSingleLink(value)
	if q.size == 0 {
		q.head = entry
	} else {
		q.tail.Next = entry
	}

	q.tail = entry
	q.size++
}

func (q *LinkedListQueue[T]) Dequeue() T {
	var value T
	if q.size > 0 {
		head := q.head
		value = head.Value
		q.head = head.Next

		// unlink and decrease size
		head.Next = nil
		q.size--
	}

	return value
}

func (q *LinkedListQueue[T]) Size() int {
	return q.size
}
