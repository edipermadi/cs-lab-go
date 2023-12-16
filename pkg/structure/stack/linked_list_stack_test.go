package stack_test

import (
	"testing"

	"github.com/edipermadi/cs-lab-go/pkg/structure/stack"
	"github.com/stretchr/testify/assert"
)

func TestLinkedListStack_Push(t *testing.T) {
	instance := stack.NewLinkedListStack[int]()

	for i := 0; i < 10; i++ {
		instance.Push(i)
	}

	for i := 9; i >= 0; i-- {
		assert.Equal(t, i, instance.Pop())
	}
}

func TestLinkedListStack_PushPair(t *testing.T) {
	instance := stack.NewLinkedListStack[int]()

	for i := 0; i < 10; i += 2 {
		x, y := i, i+1
		instance.PushPair(x, y)
	}

	for i := 9; i >= 0; i -= 2 {
		x, y := instance.PopPair()
		assert.Equal(t, i-1, x)
		assert.Equal(t, i, y)
	}
}
