package array_test

import (
	"math/rand"
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/data/array"
	"github.com/edipermadi/cs-lab-go/test"
	"github.com/stretchr/testify/require"
)

func TestArraySelectionSort(t *testing.T) {
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
				t.Run("Unique", func(t *testing.T) {
					values := rand.Perm(100)
					expected := test.Sorted(values, tc.Ascending)

					array.SelectionSort(values, tc.Ascending)
					require.Equal(t, expected, values)
				})

				t.Run("Duplicated", func(t *testing.T) {
					values := rand.Perm(50)
					values = append(values, rand.Perm(50)...)
					expected := test.Sorted(values, tc.Ascending)

					array.SelectionSort(values, tc.Ascending)
					require.Equal(t, expected, values)
				})
			})

			t.Run("String", func(t *testing.T) {
				values := []string{"ghi", "abc", "def"}
				expected := test.Sorted(values, tc.Ascending)

				array.SelectionSort(values, tc.Ascending)
				require.Equal(t, expected, values)
			})
		})
	}
}
