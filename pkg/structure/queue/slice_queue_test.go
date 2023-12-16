package queue_test

import (
	"math/rand"
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/structure/queue"
	"github.com/stretchr/testify/require"
)

func TestSliceQueue_Enqueue(t *testing.T) {
	instance := queue.NewSliceQueue[int]()
	expected := rand.Perm(100)
	for _, v := range expected {
		instance.Enqueue(v)
	}

	require.Equal(t, len(expected), instance.Size())
}

func TestSliceQueue_Dequeue(t *testing.T) {
	instance := queue.NewSliceQueue[int]()
	expected := rand.Perm(100)
	for _, v := range expected {
		instance.Enqueue(v)
	}

	actual := make([]int, 0)
	for instance.Size() > 0 {
		actual = append(actual, instance.Dequeue())
	}

	require.Equal(t, expected, actual)
}
