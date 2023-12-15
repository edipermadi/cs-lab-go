package sort

import "golang.org/x/exp/constraints"

func SliceGnomeSort[T constraints.Ordered](values []T, ascending bool) {
	// scan left to right
	for i := 0; i < len(values); {

		// when in order, move on to the next item
		if i == 0 {
			i++
		} else if ascending && values[i-1] <= values[i] {
			i++
		} else if !ascending && values[i-1] >= values[i] {
			i++
		} else {
			values[i-1], values[i] = values[i], values[i-1]
			i--
		}
	}
}
