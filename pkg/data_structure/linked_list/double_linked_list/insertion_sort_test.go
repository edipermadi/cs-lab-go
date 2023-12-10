package double_linked_list_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/data_structure/linked_list/double_linked_list"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func sorted[T constraints.Ordered](values []T, ascending bool) []T {
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

func TestDoubleLinkedListInsertionSort(t *testing.T) {
	type testCase struct {
		Title     string
		Ascending bool
	}

	testCases := []testCase{
		{Title: "Ascending", Ascending: true},
		{Title: "Descending", Ascending: false},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			t.Run("Int", func(t *testing.T) {
				values := rand.Perm(100)
				expected := sorted(values, tc.Ascending)
				node := double_linked_list.FromSlice(values)
				assert.Equal(t, expected, double_linked_list.InsertionSort(node, tc.Ascending).ToSlice())
			})

			t.Run("String", func(t *testing.T) {
				values := []string{"ghi", "abc", "def"}
				expected := sorted(values, tc.Ascending)
				node := double_linked_list.FromSlice(values)
				assert.Equal(t, expected, double_linked_list.InsertionSort(node, tc.Ascending).ToSlice())
			})
		})
	}
}
