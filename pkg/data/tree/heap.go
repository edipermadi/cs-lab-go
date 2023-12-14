package tree

import (
	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	Values    []T
	Ascending bool
}

func (h *Heap[T]) sorted(i, j Index) bool {
	return (h.Ascending && h.Values[i] <= h.Values[j]) || (!h.Ascending && h.Values[i] >= h.Values[j])
}

func (h *Heap[T]) Add(value T) {
	h.Values = append(h.Values, value)

	// start with the leap and scan up, make sure it's in order with parent
	for n := Index(len(h.Values) - 1); n > 0; {
		p := n.Parent()
		if !h.sorted(p, n) {
			h.Values[n], h.Values[p] = h.Values[p], h.Values[n]
		}

		n = p
	}
}

func (h *Heap[T]) Remove() T {
	v := h.Values[0]
	length := len(h.Values)
	if length == 1 {
		h.Values = nil
		return v
	}

	// move leaf to head
	h.Values[0] = h.Values[length-1]

	// truncate
	h.Values = h.Values[0 : length-1]

	// restore order, scan top to bottom
	for p := Index(0); ; {
		l, r := p.LeftChild(), p.RightChild()

		if int(l) > len(h.Values)-1 {
			break
		}

		// assume left node
		n := l

		// when right node is accessible switch with r when not sorted
		if int(r) <= len(h.Values)-1 && !h.sorted(l, r) {
			n = r
		}

		// swap if not sorted
		if !h.sorted(p, n) {
			h.Values[p], h.Values[n] = h.Values[n], h.Values[p]
		}

		// dive further to the next child
		p = n
	}

	return v
}
