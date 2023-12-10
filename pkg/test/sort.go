package test

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Sorted[T constraints.Ordered](values []T, ascending bool) []T {
	result := append([]T{}, values...)
	sort.SliceStable(result, func(i, j int) bool {
		if ascending {
			return result[i] < result[j]
		} else {
			return result[i] > result[j]
		}
	})
	return result
}
