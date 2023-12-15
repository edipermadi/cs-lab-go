package sort

import "golang.org/x/exp/constraints"

func SliceExchangeSort[T constraints.Ordered](values []T, ascending bool) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			swap := false
			if ascending {
				swap = values[i] > values[j]
			} else {
				swap = values[i] < values[j]
			}

			if swap {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}
