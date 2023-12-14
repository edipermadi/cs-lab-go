package array_test

import (
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/data/array"
	"github.com/stretchr/testify/assert"
)

func TestStack_Push(t *testing.T) {
	var stack array.Stack[int]

	for i := 0; i < 10; i++ {
		stack.Push(i)
	}

	for i := 9; i >= 0; i-- {
		assert.Equal(t, i, stack.Pop())
	}
}

func TestStack_PushPair(t *testing.T) {
	var stack array.Stack[int]

	for i := 0; i < 10; i += 2 {
		x, y := i, i+1
		stack.PushPair(x, y)
	}

	for i := 9; i >= 0; i -= 2 {
		x, y := stack.PopPair()
		assert.Equal(t, i-1, x)
		assert.Equal(t, i, y)
	}
}
