package shuffle_test

import (
	"math/rand"
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/algorithm/shuffle"
	"github.com/stretchr/testify/require"
)

func TestSliceFisherYatesShuffle(t *testing.T) {
	expected := rand.Perm(10)
	given := append(make([]int, 0), expected...)
	shuffle.SliceFisherYatesShuffle(given)
	require.Equal(t, len(expected), len(given))
	require.NotEqual(t, expected, given)
}
