package set

import (
	"github.com/edipermadi/cs-lab-go/pkg/data/heap"
	"golang.org/x/exp/constraints"
)

func NewSortedSet[T constraints.Ordered](ascending bool) *SortedSet[T] {
	return &SortedSet[T]{heap: heap.NewHeap[T](ascending)}
}

type SortedSet[T constraints.Ordered] struct {
	heap *heap.Heap[T]
}

func (s *SortedSet[T]) Values() []T {
	result := make([]T, 0)
	newHeap := heap.NewHeap[T](s.heap.Ascending())
	for !s.heap.Empty() {
		entry := s.heap.Remove()
		result = append(result, entry)
		newHeap.Add(entry)
	}

	// restore values and return
	s.heap = newHeap
	return result
}

func (s *SortedSet[T]) Add(value T) {
	if !s.heap.Contains(value) {
		s.heap.Add(value)
	}
}
