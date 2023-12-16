package list_test

import (
	"math/rand"
	"testing"

	"github.com/edipermadi/cs-lab-go/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSingleLinkedListFromSlice(t *testing.T) {
	values := rand.Perm(10)
	t.Logf("slice = %+v", values)

	node := test.SingleLinkedListFromSlice(values)
	require.NotNil(t, node)
	assert.Equal(t, len(values), node.Length())
	assert.Equal(t, values, test.SliceFromSingleLinkedList(node))
	assert.Equal(t, values[1:], test.SliceFromSingleLinkedList(node.Next))
	assert.Equal(t, values[len(values)-1:len(values)], test.SliceFromSingleLinkedList(node.Tail()))
}
