package sorting

import (
	"golang.org/x/exp/constraints"
)

func InsertionSort[T constraints.Ordered](values []T, ascending bool) {
	for i := 0; i < len(values); i++ {
		var left T
		current := values[i]

		// scan values left to right
		for j := 0; j <= i; j++ {
			right := values[j]

			process := false
			if ascending {
				// find index where left < current < right
				process = (j == 0 || left < current) && current < right
			} else {
				// find index where left > current > right
				process = (j == 0 || left > current) && current > right
			}

			// when condition met, shift and insert value
			if process {
				// shift right unordered values
				for k := i - 1; k >= j; k-- {
					values[k+1] = values[k]
				}

				// insert value
				values[j] = current
			}

			// update next left
			left = right
		}
	}
}
