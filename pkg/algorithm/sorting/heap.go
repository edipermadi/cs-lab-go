package sorting

import (
	"github.com/edipermadi/cs-lab-go/pkg/data_structure/tree"
	"golang.org/x/exp/constraints"
)

func buildHeap[T constraints.Ordered](values []T, ascending bool) {
	if len(values) < 2 {
		return
	}

	// ascending  : sorted means parent <= node
	// descending : sorted means parent >= node
	sorted := func(x tree.Index) bool {
		return (ascending && values[x] >= values[x.Parent()]) || (!ascending && values[x] <= values[x.Parent()])
	}

	// scan nodes left to right
	for i := tree.Index(1); int(i) < len(values); i++ {
		p := i.Parent()

		// scan nodes upward, swap if not sorted
		for j := i; !sorted(j); j = p {
			p = j.Parent()
			values[j], values[p] = values[p], values[j]
		}
	}
}

func HeapSort[T constraints.Ordered](values []T, ascending bool) {
	// when ascending, build max heap
	// when descending, build min heap
	buildHeap(values, !ascending)

	// ascending: sorted means left <= right
	// descending: sorted means left >= right
	sorted := func(i, j tree.Index) bool {
		return (ascending && values[i] <= values[j]) ||
			(!ascending && values[i] >= values[j])
	}

	// scan right to left
	for i := tree.Index(len(values)) - 1; i > 0; i-- {
		values[0], values[i] = values[i], values[0]

		// scan node top to bottom as parent
		var n tree.Index
		for p := tree.Index(0); ; p = n {
			l, r := p.LeftChild(), p.RightChild()

			// when index exceeds boundary, stop
			if int(l) >= len(values)-1 {
				break
			}

			// assume left node
			n = l

			// when in order, point to right node
			if sorted(l, r) && l < (i-1) {
				n = r
			}

			// if node not sorted against parent, swap
			if !sorted(n, p) && n < i {
				values[p], values[n] = values[n], values[p]
			}

			// dive to the next node as parent
		}
	}
}
