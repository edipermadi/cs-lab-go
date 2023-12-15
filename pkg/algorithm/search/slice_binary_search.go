package search

import "golang.org/x/exp/constraints"

func SliceBinarySearch[T constraints.Ordered](values []T, wanted T) int {
	left, right := 0, len(values)-1
	for left <= right {
		gap := right - left
		center := left + (gap / 2)
		found := values[center]
		switch {
		// found value is larger than wanted value, next time scan from left to center -1
		case found > wanted:
			right = center - 1
			// found value is smaller than wanted value, next time search from center + 1 to right
		case found < wanted:
			left = center + 1
		default:
			return center
		}
	}

	return -1
}
