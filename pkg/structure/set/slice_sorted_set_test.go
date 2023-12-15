package set_test

import (
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/structure/set"
	"github.com/stretchr/testify/require"
)

func TestSliceSortedSet_Add(t *testing.T) {
	t.Run("IgnoreDuplicates", func(t *testing.T) {
		t.Run("Ascending", func(t *testing.T) {
			given := []int{3, 4, 2, 1, 1, 3}
			expected := []int{1, 2, 3, 4}
			sortedSet := set.NewSliceSortedSet[int](true)
			for _, v := range given {
				sortedSet.Add(v)
			}

			require.Equal(t, expected, sortedSet.Values())
		})

		t.Run("Descending", func(t *testing.T) {
			given := []int{3, 4, 2, 1, 1, 3}
			expected := []int{4, 3, 2, 1}
			sortedSet := set.NewSliceSortedSet[int](false)
			for _, v := range given {
				sortedSet.Add(v)
			}

			require.Equal(t, expected, sortedSet.Values())
		})
	})
}
