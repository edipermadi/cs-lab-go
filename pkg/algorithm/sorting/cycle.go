package sorting

import (
	"golang.org/x/exp/constraints"
)

func CycleSort[T constraints.Ordered](values []T, ascending bool) {
	// scan left to right
	for {
		swapping := false
		for i := 0; i < len(values)-1; {
			left := values[i]

			// ascending mode: for each value pointed by i, count smaller numbers
			// descending mode: for each value pointed by i, count larger numbers
			cnt := 0
			for j := i + 1; j < len(values); j++ {
				right := values[j]
				if ascending && left > right {
					cnt++
				} else if !ascending && left < right {
					cnt++
				}
			}

			x := cnt + i
			if x == i || values[i] == values[x] {
				// x == i: invalid swap, move on to the next index
				// values[i] == values[x]: duplicated value. stop swapping and move on to the next index
				i++
			} else {
				// ascending mode: index of a number is equal to count of smaller numbers
				// descending mode: index of a number is equal to count of larger numbers
				values[i], values[x] = values[x], values[i]
				swapping = true
			}
		}

		// if no more swapping, consider done
		if !swapping {
			break
		}
	}
}
