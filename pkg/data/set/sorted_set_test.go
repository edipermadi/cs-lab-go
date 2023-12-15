package set_test

import (
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/data/set"
	"github.com/stretchr/testify/require"
)

func TestSortedSet_Add(t *testing.T) {
	t.Run("IgnoreDuplicates", func(t *testing.T) {
		t.Run("Ascending", func(t *testing.T) {
			given := []int{3, 4, 2, 1, 1, 3}
			expected := []int{1, 2, 3, 4}
			sortedSet := set.NewSortedSet[int](true)
			for _, v := range given {
				sortedSet.Add(v)
			}

			require.Equal(t, expected, sortedSet.Values())
		})

		t.Run("Descending", func(t *testing.T) {
			given := []int{3, 4, 2, 1, 1, 3}
			expected := []int{4, 3, 2, 1}
			set := set.NewSortedSet[int](false)
			for _, v := range given {
				set.Add(v)
			}

			require.Equal(t, expected, set.Values())
		})
	})
}
