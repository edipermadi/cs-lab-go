package set

import (
	"github.com/edipermadi/cs-lab-go/pkg/structure/heap"
	"golang.org/x/exp/constraints"
)

func NewSliceSortedSet[T constraints.Ordered](ascending bool) *SliceSortedSet[T] {
	return &SliceSortedSet[T]{heap: heap.NewSliceHeap[T](ascending)}
}

type SliceSortedSet[T constraints.Ordered] struct {
	heap *heap.SliceHeap[T]
}

func (s *SliceSortedSet[T]) Values() []T {
	result := make([]T, 0)
	newHeap := heap.NewSliceHeap[T](s.heap.Ascending())
	for !s.heap.IsEmpty() {
		entry := s.heap.Remove()
		result = append(result, entry)
		newHeap.Add(entry)
	}

	// restore values and return
	s.heap = newHeap
	return result
}

func (s *SliceSortedSet[T]) Add(value T) {
	if !s.heap.Contains(value) {
		s.heap.Add(value)
	}
}
