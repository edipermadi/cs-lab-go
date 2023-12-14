package array

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](values []T, ascending bool) {
	// keep swapping until no more
	for {
		swapping := false
		// scan left to right by pair and swap when not in order
		for i := 0; i < len(values)-1; i++ {
			swap := false
			if ascending {
				swap = values[i] > values[i+1]
			} else {
				swap = values[i] < values[i+1]
			}

			if swap {
				values[i], values[i+1] = values[i+1], values[i]
				swapping = true
			}
		}

		if !swapping {
			break
		}
	}
}
