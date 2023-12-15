package sort

import (
	"github.com/edipermadi/cs-lab-go/pkg/algorithm/shuffle"
	"golang.org/x/exp/constraints"
)

func isSliceSorted[T constraints.Ordered](values []T, ascending bool) bool {
	for i := 0; i < len(values)-1; i++ {
		if ascending && values[i] > values[i+1] {
			return false
		}

		if !ascending && values[i] < values[i+1] {
			return false
		}
	}

	return true
}

func SliceBogoSort[T constraints.Ordered](values []T, ascending bool) {
	// keep permutating until its sort

	for !isSliceSorted(values, ascending) {
		shuffle.SliceFisherYatesShuffle(values)
	}
}
