package heap

import (
	"github.com/edipermadi/cs-lab-go/pkg/data/stack"
	"github.com/edipermadi/cs-lab-go/pkg/data/tree"
	"golang.org/x/exp/constraints"
)

func NewHeap[T constraints.Ordered](ascending bool) *Heap[T] {
	return &Heap[T]{ascending: ascending}
}

type Heap[T constraints.Ordered] struct {
	values    []T
	ascending bool
}

func (h *Heap[T]) Ascending() bool {
	return h.ascending
}

func (h *Heap[T]) Values() []T {
	cloned := make([]T, len(h.values))
	copy(cloned, h.values)
	return cloned
}

func (h *Heap[T]) sorted(i, j tree.Index) bool {
	return (h.ascending && h.values[i] <= h.values[j]) || (!h.ascending && h.values[i] >= h.values[j])
}

func (h *Heap[T]) Add(value T) {
	h.values = append(h.values, value)

	// start with the leap and scan up, make sure it's in order with parent
	for n := tree.Index(len(h.values) - 1); n > 0; {
		p := n.Parent()
		if !h.sorted(p, n) {
			h.values[n], h.values[p] = h.values[p], h.values[n]
		}

		n = p
	}
}

func (h *Heap[T]) Remove() T {
	v := h.values[0]
	length := len(h.values)
	if length == 1 {
		h.values = nil
		return v
	}

	// move leaf to head
	h.values[0] = h.values[length-1]

	// truncate
	h.values = h.values[0 : length-1]

	// restore order, scan top to bottom
	for p := tree.Index(0); ; {
		l, r := p.LeftChild(), p.RightChild()

		if int(l) > len(h.values)-1 {
			break
		}

		// assume left node
		n := l

		// when right node is accessible switch with r when not sorted
		if int(r) <= len(h.values)-1 && !h.sorted(l, r) {
			n = r
		}

		// swap if not sorted
		if !h.sorted(p, n) {
			h.values[p], h.values[n] = h.values[n], h.values[p]
		}

		// dive further to the next child
		p = n
	}

	return v
}

func (h *Heap[T]) Empty() bool {
	return len(h.values) == 0
}

func (h *Heap[T]) Peek() T {
	var result T
	if len(h.values) > 0 {
		result = h.values[0]
	}

	return result
}

func (h *Heap[T]) Contains(wanted T) bool {
	if len(h.values) == 0 {
		return false
	}

	var stk stack.Stack[tree.Index]
	limit := len(h.values)
	stk.Push(0)
	for stk.Size() > 0 {
		p := stk.Pop()
		found := h.values[p]
		if found == wanted {
			return true
		}

		left, right := p.LeftChild(), p.RightChild()
		if int(left) < limit {
			stk.Push(left)
		}

		if int(right) < limit {
			stk.Push(right)
		}
	}

	return false
}
