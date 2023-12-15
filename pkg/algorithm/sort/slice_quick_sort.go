package sort

import (
	"github.com/edipermadi/cs-lab-go/pkg/structure/stack"
	"golang.org/x/exp/constraints"
)

func sliceQuickSortPartition[T constraints.Ordered](values []T, l int, h int, ascending bool) int {
	pivot := values[l]
	i, j := l, h

	for i < j {
		for ; i < len(values); i++ {
			if ascending && values[i] > pivot {
				// in ascending mode break when current value is greater than pivot
				break
			} else if !ascending && values[i] < pivot {
				// in ascending mode break when current value is less than pivot
				break
			}
		}

		for ; j > 0; j-- {
			if j < 0 || (ascending && values[j] <= pivot) || (!ascending && values[j] >= pivot) {
				break
			}
		}

		if i < j {
			values[i], values[j] = values[j], values[i]
		}
	}

	values[l], values[j] = values[j], values[l]
	return j
}

func SliceQuickSort[T constraints.Ordered](values []T, ascending bool) {
	var tmpStack stack.SliceStack[int]
	tmpStack.PushPair(0, len(values)-1)

	for tmpStack.Size() > 0 {
		l, h := tmpStack.PopPair()

		var j int
		if l < h {
			j = sliceQuickSortPartition(values, l, h, ascending)
			tmpStack.PushPair(l, j)
			tmpStack.PushPair(j+1, h)
		}
	}
}
