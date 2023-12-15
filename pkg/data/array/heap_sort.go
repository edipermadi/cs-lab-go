package array

import (
	"github.com/edipermadi/cs-lab-go/pkg/data/heap"
	"golang.org/x/exp/constraints"
)

func HeapSort[T constraints.Ordered](values []T, ascending bool) {
	// ascending: build max heap
	// descending: build min heap
	heap := heap.NewHeap[T](ascending)
	for _, v := range values {
		heap.Add(v)
	}

	// remove from heap one by one
	for i := 0; i < len(values); i++ {
		values[i] = heap.Remove()
	}
}
