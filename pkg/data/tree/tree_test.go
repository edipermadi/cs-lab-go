package tree_test

import (
	"fmt"
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/data/tree"
	"github.com/stretchr/testify/assert"
)

func TestIndex_Level(t *testing.T) {
	type testCase struct {
		Index tree.Index
		Level int
	}

	testCases := []testCase{
		{0, 0},

		{1, 1},
		{2, 1},

		{3, 2},
		{4, 2},
		{5, 2},
		{6, 2},

		{7, 3},
		{8, 3},
		{9, 3},
		{10, 3},
		{11, 3},
		{12, 3},
		{13, 3},
		{14, 3},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.Level, tc.Index.Level())
	}
}

func TestIndex_Parent(t *testing.T) {
	type testCase struct {
		Child  tree.Index
		Parent tree.Index
	}

	testCases := []testCase{
		{0, 0},

		{1, 0},
		{2, 0},

		{3, 1},
		{4, 1},
		{5, 2},
		{6, 2},

		{7, 3},
		{8, 3},
		{9, 4},
		{10, 4},
		{11, 5},
		{12, 5},
		{13, 6},
		{14, 6},
	}

	for _, tc := range testCases {
		actual := tc.Child.Parent()
		assert.Equal(t, tc.Parent, actual, fmt.Sprintf("failed to get parent of %d, expected %d, actual %d", tc.Child, tc.Child, actual))
	}
}

func TestIndex_LeftRightChild(t *testing.T) {
	type testCase struct {
		Child  tree.Index
		Parent tree.Index
	}

	testCases := []testCase{
		{1, 0},
		{2, 0},

		{3, 1},
		{4, 1},
		{5, 2},
		{6, 2},

		{7, 3},
		{8, 3},
		{9, 4},
		{10, 4},
		{11, 5},
		{12, 5},
		{13, 6},
		{14, 6},

		{15, 7},
		{16, 7},
		{17, 8},
		{18, 8},
		{19, 9},
		{20, 9},
		{21, 10},
		{22, 10},
		{23, 11},
		{24, 11},
		{25, 12},
		{26, 12},
		{27, 13},
		{28, 13},
		{29, 14},
		{30, 14},
	}

	for _, tc := range testCases {
		if tc.Child&1 > 0 {
			actual := tc.Parent.LeftChild()
			assert.Equal(t, tc.Child, actual, fmt.Sprintf("failed to get left child of %d, expected %d, actual %d", tc.Parent, tc.Child, actual))
		} else {
			actual := tc.Parent.RightChild()
			assert.Equal(t, tc.Child, actual, fmt.Sprintf("failed to get right child of %d, expected %d, actual %d", tc.Parent, tc.Child, actual))
		}

	}
}

func TestIndex_Sibling(t *testing.T) {
	type testCase struct {
		Child   tree.Index
		Sibling tree.Index
	}

	testCases := []testCase{
		{0, 0},

		{1, 2},
		{2, 1},

		{3, 4},
		{4, 3},
		{5, 6},
		{6, 5},

		{7, 8},
		{8, 7},
		{9, 10},
		{10, 9},
		{11, 12},
		{12, 11},
		{13, 14},
		{14, 13},
	}

	for _, tc := range testCases {
		actual := tc.Child.Sibling()
		assert.Equal(t, tc.Sibling, actual, fmt.Sprintf("failed to get sibling of %d, expected %d, actual %d", tc.Child, tc.Sibling, actual))
	}
}
