package array

import (
	"github.com/edipermadi/cs-lab-go/pkg/data/stack"
	"golang.org/x/exp/constraints"
)

func partition[T constraints.Ordered](values []T, l int, h int, ascending bool) int {
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

func QuickSort[T constraints.Ordered](values []T, ascending bool) {
	var stack stack.Stack[int]
	stack.PushPair(0, len(values)-1)

	for stack.Size() > 0 {
		l, h := stack.PopPair()

		var j int
		if l < h {
			j = partition(values, l, h, ascending)
			stack.PushPair(l, j)
			stack.PushPair(j+1, h)
		}
	}
}
