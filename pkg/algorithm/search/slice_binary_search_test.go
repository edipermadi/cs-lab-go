package search_test

import (
	"math/rand"
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/algorithm/search"
	"github.com/edipermadi/cs-lab-go/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceBinarySearch(t *testing.T) {
	t.Run("Sorted", func(t *testing.T) {
		sorted := test.Sorted[int](rand.Perm(100), true)
		wanted := rand.Intn(99)
		index := search.SliceBinarySearch(sorted, wanted)
		require.True(t, index >= 0)
		assert.Equal(t, wanted, sorted[index])
	})

	t.Run("Random", func(t *testing.T) {
		values := rand.Perm(100)
		wanted := rand.Intn(99)
		index := search.SliceBinarySearch(values, wanted)
		require.True(t, index < 0)
	})
}
