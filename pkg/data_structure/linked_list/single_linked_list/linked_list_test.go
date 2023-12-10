package single_linked_list_test

import (
	"math/rand"
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/data_structure/linked_list/single_linked_list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromSlice(t *testing.T) {
	values := rand.Perm(10)
	t.Logf("slice = %+v", values)

	node := single_linked_list.FromSlice(values)
	require.NotNil(t, node)
	assert.Equal(t, len(values), node.Length())
	assert.Equal(t, values, node.ToSlice())
	assert.Equal(t, values[1:], node.Next.ToSlice())
	assert.Equal(t, values[len(values)-1:len(values)], node.Tail().ToSlice())
}
