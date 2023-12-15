package heap_test

import (
	"math/rand"
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/data/heap"
	"github.com/stretchr/testify/require"
)

func TestHeap_Add(t *testing.T) {
	type testCase struct {
		Title     string
		Ascending bool
		Amount    int
		Expected  int
	}

	testCases := []testCase{
		{"Min", true, 10, 0},
		{"Max", false, 10, 9},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			h := heap.NewHeap[int](tc.Ascending)
			for _, v := range rand.Perm(tc.Amount) {
				h.Add(v)
			}

			t.Logf("%s heap %+v\n", tc.Title, h.Values())
			require.Equal(t, tc.Expected, h.Peek())
		})
	}
}

func TestHeap_Remove(t *testing.T) {
	type testCase struct {
		Title     string
		Ascending bool
		Amount    int
	}

	testCases := []testCase{
		{"Min", true, 10},
		{"Max", false, 10},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			h := heap.NewHeap[int](tc.Ascending)
			for _, v := range rand.Perm(tc.Amount) {
				h.Add(v)
			}

			popped := make([]int, 0)
			for i := 0; i < tc.Amount; i++ {
				popped = append(popped, h.Remove())
			}

			t.Logf("popped = %+v", popped)
		})
	}
}
