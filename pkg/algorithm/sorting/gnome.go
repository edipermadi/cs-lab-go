package sorting

import "golang.org/x/exp/constraints"

func GnomeSort[T constraints.Ordered](values []T, ascending bool) {
	// scan left to right
	for i := 0; i < len(values); {

		// when in order, move on to the next item
		if i == 0 || (ascending && values[i-1] < values[i]) || (!ascending && values[i-1] > values[i]) {
			i++
		} else {
			values[i-1], values[i] = values[i], values[i-1]
			i--
		}
	}
}
