package sorting

import (
	"golang.org/x/exp/constraints"
)

func SelectionSort[T constraints.Ordered](values []T, ascending bool) {
	// scan left to right
	for i := 0; i < len(values)-1; i++ {
		// assume current index holds the smallest/largest value
		j := i

		for k := i + 1; k < len(values); k++ {
			current := values[k]
			selected := values[j]

			// when in ascending mode update j with the index that has the smallest value
			// for descending do the opposite
			if ascending && current < selected {
				j = k
			} else if !ascending && current > selected {
				j = k
			}
		}

		// if index is not the same swap them
		if i != j {
			values[i], values[j] = values[j], values[i]
		}
	}
}
