package sort_test

import (
	"math/rand"
	"testing"

	"github.com/edipermadi/cs-lab-go/internal/test"
	"github.com/edipermadi/cs-lab-go/pkg/algorithm/sort"
	"github.com/edipermadi/cs-lab-go/pkg/structure/linked_list/double_linked_list"
	"github.com/stretchr/testify/assert"
)

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
				expected := test.Sorted(values, tc.Ascending)
				node := double_linked_list.FromSlice(values)
				assert.Equal(t, expected, sort.DoubleLinkedListInsertionSort(node, tc.Ascending).ToSlice())
			})

			t.Run("String", func(t *testing.T) {
				values := []string{"ghi", "abc", "def"}
				expected := test.Sorted(values, tc.Ascending)
				node := double_linked_list.FromSlice(values)
				assert.Equal(t, expected, sort.DoubleLinkedListInsertionSort(node, tc.Ascending).ToSlice())
			})
		})
	}
}
