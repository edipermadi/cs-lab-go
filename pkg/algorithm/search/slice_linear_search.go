package search

import "golang.org/x/exp/constraints"

func SliceLinearSearch[T constraints.Ordered](values []T, wanted T) int {
	for i := 0; i < len(values); i++ {
		if values[i] == wanted {
			return i
		}
	}

	return -1
}
