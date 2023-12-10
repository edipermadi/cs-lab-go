package sorting_test

import (
	"github.com/edipermadi/cs-lab-go/pkg/data_structure/linked_list/double_linked_list"
	"math/rand"
	"sort"
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/algorithm/sorting"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func sorted[T constraints.Ordered](values []T, ordering sorting.Ordering) []T {
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

func TestDoubleLinkedListInsertionSort(t *testing.T) {
	t.Run("Ascending", func(t *testing.T) {
		t.Run("Int", func(t *testing.T) {
			values := rand.Perm(100)
			expected := sorted(values, sorting.Ascending)
			node := double_linked_list.FromSlice(values)
			assert.Equal(t, expected, sorting.DoubleLinkedListInsertionSort(node, sorting.Ascending).ToSlice())
		})

		t.Run("String", func(t *testing.T) {
			values := []string{"ghi", "abc", "def"}
			expected := sorted(values, sorting.Ascending)
			node := double_linked_list.FromSlice(values)
			assert.Equal(t, expected, sorting.DoubleLinkedListInsertionSort(node, sorting.Ascending).ToSlice())
		})
	})

	t.Run("Descending", func(t *testing.T) {
		t.Run("Int", func(t *testing.T) {
			values := rand.Perm(100)
			expected := sorted(values, sorting.Descending)
			node := double_linked_list.FromSlice(values)
			assert.Equal(t, expected, sorting.DoubleLinkedListInsertionSort(node, sorting.Descending).ToSlice())
		})

		t.Run("String", func(t *testing.T) {
			values := []string{"ghi", "abc", "def"}
			expected := sorted(values, sorting.Descending)
			node := double_linked_list.FromSlice(values)
			assert.Equal(t, expected, sorting.DoubleLinkedListInsertionSort(node, sorting.Descending).ToSlice())
		})
	})
}
