package sorting_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/algorithm/sorting"
	"github.com/stretchr/testify/assert"
)

func sorted[T int | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64 | string](values []T, ordering sorting.Ordering) []T {
	result := append([]T{}, values...)
	sort.SliceStable(result, func(i, j int) bool {
		if ordering == sorting.Ascending {
			return result[i] < result[j]
		} else {
			return result[i] > result[j]
		}
	})
	return result
}

func TestSliceInsertionSort(t *testing.T) {
	t.Run("Ascending", func(t *testing.T) {
		t.Run("Int", func(t *testing.T) {
			values := rand.Perm(100)
			expected := sorted(values, sorting.Ascending)

			sorting.SliceInsertionSort(values, sorting.Ascending)
			assert.Equal(t, expected, values)
		})

		t.Run("String", func(t *testing.T) {
			values := []string{"ghi", "abc", "def"}
			expected := sorted(values, sorting.Ascending)

			sorting.SliceInsertionSort(values, sorting.Ascending)
			assert.Equal(t, expected, values)
		})
	})

	t.Run("Descending", func(t *testing.T) {
		t.Run("Int", func(t *testing.T) {
			values := rand.Perm(100)
			expected := sorted(values, sorting.Descending)

			sorting.SliceInsertionSort(values, sorting.Descending)
			assert.Equal(t, expected, values)
		})

		t.Run("String", func(t *testing.T) {
			values := []string{"ghi", "abc", "def"}
			expected := sorted(values, sorting.Descending)

			sorting.SliceInsertionSort(values, sorting.Descending)
			assert.Equal(t, expected, values)
		})
	})

}
