package sort

import (
	"github.com/edipermadi/cs-lab-go/pkg/structure/heap"
	"golang.org/x/exp/constraints"
)

func SliceHeapSort[T constraints.Ordered](values []T, ascending bool) {
	// ascending: build max heap
	// descending: build min heap
	tmpHeap := heap.NewSliceHeap[T](ascending)
	for _, v := range values {
		tmpHeap.Add(v)
	}

	// remove from heap one by one
	for i := 0; i < len(values); i++ {
		values[i] = tmpHeap.Remove()
	}
}
